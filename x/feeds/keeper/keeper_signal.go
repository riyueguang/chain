package keeper

import (
	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v3/x/feeds/types"
)

// GetVote returns all signals voted by a voter.
func (k Keeper) GetVote(ctx sdk.Context, voter sdk.AccAddress) []types.Signal {
	bz := ctx.KVStore(k.storeKey).Get(types.VoteStoreKey(voter))
	if bz == nil {
		return nil
	}

	var v types.Vote
	k.cdc.MustUnmarshal(bz, &v)

	return v.Signals
}

// DeleteVote deletes all signals voted by a voter.
func (k Keeper) DeleteVote(ctx sdk.Context, voter sdk.AccAddress) {
	ctx.KVStore(k.storeKey).Delete(types.VoteStoreKey(voter))
}

// SetVote sets signals voted by a voter.
func (k Keeper) SetVote(ctx sdk.Context, vote types.Vote) {
	ctx.KVStore(k.storeKey).
		Set(types.VoteStoreKey(sdk.MustAccAddressFromBech32(vote.Voter)), k.cdc.MustMarshal(&vote))
}

// GetVoteIterator returns an iterator of the vote store.
func (k Keeper) GetVoteIterator(ctx sdk.Context) dbm.Iterator {
	return storetypes.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.VoteStoreKeyPrefix)
}

// GetVotes returns all votes.
func (k Keeper) GetVotes(ctx sdk.Context) (votes []types.Vote) {
	iterator := k.GetVoteIterator(ctx)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var v types.Vote
		k.cdc.MustUnmarshal(iterator.Value(), &v)
		votes = append(votes, v)
	}

	return
}

// SetVotes sets multiple votes.
func (k Keeper) SetVotes(ctx sdk.Context, votes []types.Vote) {
	for _, v := range votes {
		k.SetVote(ctx, v)
	}
}

// SetSignalTotalPower sets signal-total-power to the store.
func (k Keeper) SetSignalTotalPower(ctx sdk.Context, signal types.Signal) {
	prevSignalTotalPower, err := k.GetSignalTotalPower(ctx, signal.ID)
	if err == nil {
		k.deleteSignalTotalPowerByPowerIndex(ctx, prevSignalTotalPower)
	}

	if signal.Power == 0 {
		k.deleteSignalTotalPower(ctx, signal.ID)
		emitEventDeleteSignalTotalPower(ctx, signal)
	} else {
		ctx.KVStore(k.storeKey).
			Set(types.SignalTotalPowerStoreKey(signal.ID), k.cdc.MustMarshal(&signal))
		k.setSignalTotalPowerByPowerIndex(ctx, signal)
		emitEventUpdateSignalTotalPower(ctx, signal)
	}
}

// GetSignalTotalPower gets a signal-total-power from specified signal id.
func (k Keeper) GetSignalTotalPower(ctx sdk.Context, signalID string) (types.Signal, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.SignalTotalPowerStoreKey(signalID))
	if bz == nil {
		return types.Signal{}, types.ErrSignalTotalPowerNotFound.Wrapf(
			"failed to get signal-total-power for signal id: %s",
			signalID,
		)
	}

	var s types.Signal
	k.cdc.MustUnmarshal(bz, &s)

	return s, nil
}

// deleteSignalTotalPower deletes a signal-total-power by signal id.
func (k Keeper) deleteSignalTotalPower(ctx sdk.Context, signalID string) {
	ctx.KVStore(k.storeKey).Delete(types.SignalTotalPowerStoreKey(signalID))
}

// SetSignalTotalPowers sets multiple signal-total-powers.
func (k Keeper) SetSignalTotalPowers(ctx sdk.Context, signalTotalPowersList []types.Signal) {
	for _, stp := range signalTotalPowersList {
		k.SetSignalTotalPower(ctx, stp)
	}
}

func (k Keeper) setSignalTotalPowerByPowerIndex(ctx sdk.Context, signalTotalPower types.Signal) {
	ctx.KVStore(k.storeKey).
		Set(types.SignalTotalPowerByPowerIndexKey(signalTotalPower.ID, signalTotalPower.Power), []byte(signalTotalPower.ID))
}

func (k Keeper) deleteSignalTotalPowerByPowerIndex(ctx sdk.Context, signalTotalPower types.Signal) {
	ctx.KVStore(k.storeKey).
		Delete(types.SignalTotalPowerByPowerIndexKey(signalTotalPower.ID, signalTotalPower.Power))
}

// GetSignalTotalPowersByPower gets the current signal-total-power sorted by power-rank.
func (k Keeper) GetSignalTotalPowersByPower(ctx sdk.Context, limit uint64) []types.Signal {
	signalTotalPowers := make([]types.Signal, limit)

	iterator := k.SignalTotalPowersByPowerStoreIterator(ctx)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(limit); iterator.Next() {
		bz := iterator.Value()
		signalID := string(bz)
		signalTotalPower, err := k.GetSignalTotalPower(ctx, signalID)
		if err != nil || signalTotalPower.Power == 0 {
			// this should not happen
			continue
		}

		signalTotalPowers[i] = signalTotalPower
		i++
	}

	return signalTotalPowers[:i] // trim
}

// SignalTotalPowersByPowerStoreIterator returns an iterator for signal-total-powers by power index store.
func (k Keeper) SignalTotalPowersByPowerStoreIterator(ctx sdk.Context) dbm.Iterator {
	return storetypes.KVStoreReversePrefixIterator(
		ctx.KVStore(k.storeKey),
		types.SignalTotalPowerByPowerIndexKeyPrefix,
	)
}

// LockVoterPower locks the voter's power equal to the sum of the signal powers.
// It returns an error if the voter does not have enough power to lock.
func (k Keeper) LockVoterPower(
	ctx sdk.Context,
	voter sdk.AccAddress,
	signals []types.Signal,
) error {
	sumPower := types.SumPower(signals)
	if err := k.restakeKeeper.SetLockedPower(ctx, voter, types.ModuleName, math.NewInt(sumPower)); err != nil {
		return err
	}

	return nil
}

// UpdateVoteAndReturnPowerDiff delete previous signals and add new signals.
// It also calculates feed power differences from voter's previous signals and new signals.
func (k Keeper) UpdateVoteAndReturnPowerDiff(
	ctx sdk.Context,
	voter sdk.AccAddress,
	signals []types.Signal,
) map[string]int64 {
	signalIDToPowerDiff := make(map[string]int64)

	prevSignals := k.GetVote(ctx, voter)
	k.DeleteVote(ctx, voter)

	for _, prevSignal := range prevSignals {
		signalIDToPowerDiff[prevSignal.ID] -= prevSignal.Power
	}

	k.SetVote(ctx, types.NewVote(voter.String(), signals))

	for _, signal := range signals {
		signalIDToPowerDiff[signal.ID] += signal.Power
	}

	return signalIDToPowerDiff
}
