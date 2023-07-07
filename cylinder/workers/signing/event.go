package signing

import (
	"errors"

	"github.com/bandprotocol/chain/v2/pkg/event"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Event represents the data structure for request-sign events.
type Event struct {
	GroupID       tss.GroupID
	SigningID     tss.SigningID
	MemberIDs     []tss.MemberID
	GroupPubNonce tss.Point
	Data          []byte
	BindingFactor tss.Scalar
	PubDE         types.DE
}

// ParseEvent parses the request-sign event from the given message log.
// It extracts the signing information from the log and returns the parsed Event or an error if parsing fails.
func ParseEvent(log sdk.ABCIMessageLog, address string) (*Event, error) {
	gid, err := event.GetEventValueUint64(log, types.EventTypeRequestSign, types.AttributeKeyGroupID)
	if err != nil {
		return nil, err
	}

	sid, err := event.GetEventValueUint64(log, types.EventTypeRequestSign, types.AttributeKeySigningID)
	if err != nil {
		return nil, err
	}

	groupPubNonce, err := event.GetEventValueBytes(log, types.EventTypeRequestSign, types.AttributeKeyGroupPubNonce)
	if err != nil {
		return nil, err
	}

	bindingFactors, err := event.GetEventValuesBytes(
		log,
		types.EventTypeRequestSign,
		types.AttributeKeyBindingFactor,
	)
	if err != nil {
		return nil, err
	}

	data, err := event.GetEventValueBytes(log, types.EventTypeRequestSign, types.AttributeKeyMessage)
	if err != nil {
		return nil, err
	}

	midInts, err := event.GetEventValuesUint64(log, types.EventTypeRequestSign, types.AttributeKeyMemberID)
	if err != nil {
		return nil, err
	}

	pubDs, err := event.GetEventValuesBytes(log, types.EventTypeRequestSign, types.AttributeKeyPubD)
	if err != nil {
		return nil, err
	}

	pubEs, err := event.GetEventValuesBytes(log, types.EventTypeRequestSign, types.AttributeKeyPubE)
	if err != nil {
		return nil, err
	}

	var pubD, pubE tss.Point
	var mids []tss.MemberID
	var bindingFactor tss.Scalar
	found := false

	members := event.GetEventValues(log, types.EventTypeRequestSign, types.AttributeKeyMember)
	for i, member := range members {
		mids = append(mids, tss.MemberID(midInts[i]))
		if member == address {
			found = true
			bindingFactor = bindingFactors[i]
			pubD = pubDs[i]
			pubE = pubEs[i]
		}
	}

	if !found {
		return nil, errors.New("failed to find member in the event")
	}

	return &Event{
		GroupID:       tss.GroupID(gid),
		SigningID:     tss.SigningID(sid),
		MemberIDs:     mids,
		GroupPubNonce: groupPubNonce,
		Data:          data,
		BindingFactor: bindingFactor,
		PubDE: types.DE{
			PubD: pubD,
			PubE: pubE,
		},
	}, nil
}
