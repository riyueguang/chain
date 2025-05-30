package proof

import (
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/libs/protoio"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttypes "github.com/cometbft/cometbft/types"
)

// TMSignatureEthereum is an Ethereum version of TMSignature for solidity ABI-encoding.
type TMSignatureEthereum struct {
	R                common.Hash
	S                common.Hash
	V                uint8
	EncodedTimestamp []byte
}

func (signature *TMSignature) encodeToEthFormat() TMSignatureEthereum {
	return TMSignatureEthereum{
		R:                common.BytesToHash(signature.R),
		S:                common.BytesToHash(signature.S),
		V:                uint8(signature.V),
		EncodedTimestamp: signature.EncodedTimestamp,
	}
}

func recoverETHAddress(msg, sig, signer []byte) ([]byte, uint8, error) {
	for i := uint8(0); i < 2; i++ {
		pubuc, err := crypto.SigToPub(tmhash.Sum(msg), append(sig, i))
		if err != nil {
			return nil, 0, err
		}
		pub := crypto.CompressPubkey(pubuc)
		var tmp [33]byte

		copy(tmp[:], pub)
		if string(signer) == string(secp256k1.PubKey(tmp[:]).Address()) {
			return crypto.PubkeyToAddress(*pubuc).Bytes(), 27 + i, nil
		}
	}
	return nil, 0, fmt.Errorf("no match address found")
}

func GetPrefix(t cmtproto.SignedMsgType, height int64, round int64) ([]byte, error) {
	prefix, err := protoio.MarshalDelimited(
		&cmtproto.CanonicalVote{
			Type:   t,
			Height: height,
			Round:  round,
		},
	)
	if err != nil {
		return nil, err
	}
	length := int(prefix[0])
	// prefix should be X + default timestamp that equals to `2a0b088092b8c398feffffff01`, so we trim last 13 bytes
	return prefix[1 : length-12], nil
}

// GetSignaturesAndPrefix returns a list of TMSignature from Tendermint signed header.
func GetSignaturesAndPrefix(info *cmttypes.SignedHeader) ([]TMSignature, CommonEncodedVotePart, error) {
	addrs := []string{}
	mapAddrs := map[string]TMSignature{}

	prefix, err := GetPrefix(
		cmtproto.SignedMsgType(byte(cmtproto.PrecommitType)),
		info.Commit.Height,
		int64(info.Commit.Round),
	)
	if err != nil {
		return nil, CommonEncodedVotePart{}, err
	}
	// Append with 4 fixed bytes
	// 34 is a key for the CanonicalBlockID ( 34 == (4 << 3) | 2 )
	// 72 is the length of the CanonicalBlockID message ( 72 == (1+1 + 32) + (1+1 + (1 + 1) + (1+1 + 32)) )
	// 10 is a key for the blockHash ( 10 == (1 << 3) | 2 )
	// 32 is the length of the blockHash
	prefix = append(prefix, []byte{34, 72, 10, 32}...)

	suffix, err := protoio.MarshalDelimited(
		&cmtproto.CanonicalPartSetHeader{
			Total: info.Commit.BlockID.PartSetHeader.Total,
			Hash:  info.Commit.BlockID.PartSetHeader.Hash,
		},
	)
	if err != nil {
		return nil, CommonEncodedVotePart{}, err
	}
	// Append with 1 fixed byte
	// 18 is a key for the CanonicalPartSetHeader ( 18 == (2 << 3) | 2 )
	suffix = append([]byte{18}, suffix...)

	commonVote := CommonEncodedVotePart{SignedDataPrefix: prefix, SignedDataSuffix: suffix}

	commonPart := append(commonVote.SignedDataPrefix, info.Commit.BlockID.Hash...)
	commonPart = append(commonPart, commonVote.SignedDataSuffix...)

	chainIDBytes := []byte(info.ChainID)
	encodedChainIDConstant := append([]byte{50, uint8(len(chainIDBytes))}, chainIDBytes...)

	for _, vote := range info.Commit.Signatures {
		if vote.BlockIDFlag != cmttypes.BlockIDFlagCommit {
			continue
		}

		encodedTimestamp := encodeTime(vote.Timestamp)

		msg := append(commonPart, []byte{42, uint8(len(encodedTimestamp))}...)
		msg = append(msg, encodedTimestamp...)
		msg = append(msg, encodedChainIDConstant...)
		msg = append([]byte{uint8(len(msg))}, msg...)

		addr, v, err := recoverETHAddress(msg, vote.Signature, vote.ValidatorAddress)
		if err != nil {
			return nil, CommonEncodedVotePart{}, err
		}
		addrs = append(addrs, string(addr))
		mapAddrs[string(addr)] = TMSignature{
			vote.Signature[:32],
			vote.Signature[32:],
			uint32(v),
			encodedTimestamp,
		}
	}
	if len(addrs) == 0 {
		return nil, CommonEncodedVotePart{}, fmt.Errorf("no valid precommit")
	}

	signatures := make([]TMSignature, len(addrs))
	sort.Strings(addrs)
	for i, addr := range addrs {
		signatures[i] = mapAddrs[addr]
	}

	return signatures, commonVote, nil
}
