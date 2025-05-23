package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v3/pkg/tss"
)

// NewMember creates a new Member instance.
func NewMember(
	id tss.MemberID,
	groupID tss.GroupID,
	addr sdk.AccAddress,
	pubKey tss.Point,
	isMalicious bool,
	isActive bool,
) Member {
	return Member{
		ID:          id,
		GroupID:     groupID,
		Address:     addr.String(),
		PubKey:      pubKey,
		IsMalicious: isMalicious,
		IsActive:    isActive,
	}
}

// Validate performs basic validation of group information.
func (m Member) Validate() error {
	if m.ID == 0 {
		return ErrInvalidMember.Wrap("member id cannot be 0")
	}

	if m.GroupID == 0 {
		return ErrInvalidMember.Wrap("group id cannot be 0")
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return ErrInvalidMember.Wrapf("invalid member address: %v", err)
	}

	if err := m.PubKey.Validate(); err != nil {
		return ErrInvalidMember.Wrapf("invalid member public key: %v", err)
	}

	return nil
}

// Members represents a slice of Member values.
type Members []Member

// GetIDs returns an array of MemberIDs from a collection of members
func (ms Members) GetIDs() []tss.MemberID {
	var mids []tss.MemberID
	for _, member := range ms {
		mids = append(mids, member.ID)
	}

	return mids
}

// HaveMalicious checks if any member in the collection is marked as malicious
func (ms Members) HaveMalicious() bool {
	for _, m := range ms {
		if m.IsMalicious {
			return true
		}
	}

	return false
}
