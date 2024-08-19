package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/restake/types"
)

// SetLockedPower sets the new locked power of the address to the key
// This function will override the previous locked power.
func (k Keeper) SetLockedPower(ctx sdk.Context, stakerAddr sdk.AccAddress, keyName string, power sdkmath.Int) error {
	if !power.IsUint64() {
		return types.ErrInvalidPower
	}

	// check if delegation is not less than power
	delegation := k.stakingKeeper.GetDelegatorBonded(ctx, stakerAddr)
	if delegation.LT(power) {
		return types.ErrDelegationNotEnough
	}

	key, err := k.GetOrCreateKey(ctx, keyName)
	if err != nil {
		return err
	}

	if !key.IsActive {
		return types.ErrKeyNotActive
	}

	// check if there is a lock before
	lock, err := k.GetLock(ctx, stakerAddr, keyName)
	if err != nil {
		lock = types.Lock{
			StakerAddress:  stakerAddr.String(),
			Key:            keyName,
			Power:          sdkmath.NewInt(0),
			PosRewardDebts: sdk.NewDecCoins(),
			NegRewardDebts: sdk.NewDecCoins(),
		}
	}

	diffPower := power.Sub(lock.Power)

	key.TotalPower = key.TotalPower.Add(diffPower)
	k.SetKey(ctx, key)

	additionalDebts := key.RewardPerPowers.MulDecTruncate(sdkmath.LegacyNewDecFromInt(diffPower.Abs()))
	if diffPower.IsPositive() {
		lock.PosRewardDebts = lock.PosRewardDebts.Add(additionalDebts...)
	} else {
		lock.NegRewardDebts = lock.NegRewardDebts.Add(additionalDebts...)
	}
	lock.Power = power
	k.SetLock(ctx, lock)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLockPower,
			sdk.NewAttribute(types.AttributeKeyStaker, stakerAddr.String()),
			sdk.NewAttribute(types.AttributeKeyKey, keyName),
			sdk.NewAttribute(types.AttributeKeyPower, power.String()),
		),
	)

	return nil
}

// GetLockedPower returns locked power of the address to the key.
func (k Keeper) GetLockedPower(ctx sdk.Context, stakerAddr sdk.AccAddress, keyName string) (sdkmath.Int, error) {
	key, err := k.GetKey(ctx, keyName)
	if err != nil {
		return sdkmath.Int{}, types.ErrKeyNotFound
	}

	if !key.IsActive {
		return sdkmath.Int{}, types.ErrKeyNotActive
	}

	lock, err := k.GetLock(ctx, stakerAddr, keyName)
	if err != nil {
		return sdkmath.Int{}, types.ErrLockNotFound
	}

	return lock.Power, nil
}

// getAccumulatedRewards gets the accumulatedRewards of a lock if they lock since beginning.
func (k Keeper) getAccumulatedRewards(ctx sdk.Context, lock types.Lock) sdk.DecCoins {
	key := k.MustGetKey(ctx, lock.Key)

	return key.RewardPerPowers.MulDecTruncate(sdkmath.LegacyNewDecFromInt(lock.Power))
}

// getReward gets the reward of a lock by using accumulated rewards and reward debts.
func (k Keeper) getReward(ctx sdk.Context, lock types.Lock) types.Reward {
	totalRewards := k.getAccumulatedRewards(ctx, lock)

	return types.Reward{
		Key:     lock.Key,
		Rewards: totalRewards.Add(lock.NegRewardDebts...).Sub(lock.PosRewardDebts),
	}
}

// -------------------------------
// store part
// -------------------------------

// GetLocksIterator gets iterator of lock store.
func (k Keeper) GetLocksIterator(ctx sdk.Context) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.LockStoreKeyPrefix)
}

// GetLocksByAddressIterator gets iterator of locks of the speicfic address.
func (k Keeper) GetLocksByAddressIterator(ctx sdk.Context, addr sdk.AccAddress) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.LocksByAddressStoreKey(addr))
}

// GetLocksByAddress gets all locks of the address.
func (k Keeper) GetLocksByAddress(ctx sdk.Context, addr sdk.AccAddress) (locks []types.Lock) {
	iterator := k.GetLocksByAddressIterator(ctx, addr)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var lock types.Lock
		k.cdc.MustUnmarshal(iterator.Value(), &lock)
		locks = append(locks, lock)
	}

	return locks
}

// GetLocks gets all locks in the store.
func (k Keeper) GetLocks(ctx sdk.Context) (locks []types.Lock) {
	iterator := k.GetLocksIterator(ctx)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var lock types.Lock
		k.cdc.MustUnmarshal(iterator.Value(), &lock)
		locks = append(locks, lock)
	}

	return locks
}

// HasLock checks if lock exists in the store.
func (k Keeper) HasLock(ctx sdk.Context, addr sdk.AccAddress, keyName string) bool {
	return ctx.KVStore(k.storeKey).Has(types.LockStoreKey(addr, keyName))
}

// GetLock gets a lock from store by address and key name.
func (k Keeper) GetLock(ctx sdk.Context, addr sdk.AccAddress, keyName string) (types.Lock, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.LockStoreKey(addr, keyName))
	if bz == nil {
		return types.Lock{}, types.ErrLockNotFound.Wrapf(
			"failed to get lock of %s with key name: %s",
			addr.String(),
			keyName,
		)
	}

	var lock types.Lock
	k.cdc.MustUnmarshal(bz, &lock)

	return lock, nil
}

// SetLock sets a lock to the store.
func (k Keeper) SetLock(ctx sdk.Context, lock types.Lock) {
	addr := sdk.MustAccAddressFromBech32(lock.StakerAddress)
	k.DeleteLock(ctx, addr, lock.Key)

	ctx.KVStore(k.storeKey).Set(types.LockStoreKey(addr, lock.Key), k.cdc.MustMarshal(&lock))
	k.setLockByPower(ctx, lock)
}

// DeleteLock deletes a lock from the store.
func (k Keeper) DeleteLock(ctx sdk.Context, addr sdk.AccAddress, keyName string) {
	lock, err := k.GetLock(ctx, addr, keyName)
	if err != nil {
		return
	}
	ctx.KVStore(k.storeKey).Delete(types.LockStoreKey(addr, keyName))
	k.deleteLockByPower(ctx, lock)
}

// setLockByPower sets a lock by power to the store.
func (k Keeper) setLockByPower(ctx sdk.Context, lock types.Lock) {
	ctx.KVStore(k.storeKey).Set(types.LockByPowerIndexKey(lock), []byte(lock.Key))
}

// deleteLockByPower deletes a lock by power from the store.
func (k Keeper) deleteLockByPower(ctx sdk.Context, lock types.Lock) {
	ctx.KVStore(k.storeKey).Delete(types.LockByPowerIndexKey(lock))
}
