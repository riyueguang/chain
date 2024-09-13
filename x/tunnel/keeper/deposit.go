package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/tunnel/types"
)

// validateDepositDenom validates if the deposit denom is accepted by the tunnel module.
func (k Keeper) validateDepositDenom(ctx sdk.Context, depositAmount sdk.Coins) error {
	params := k.GetParams(ctx)
	denoms := make([]string, 0, len(params.MinDeposit))
	acceptedDenoms := make(map[string]bool, len(params.MinDeposit))
	for _, coin := range params.MinDeposit {
		acceptedDenoms[coin.Denom] = true
		denoms = append(denoms, coin.Denom)
	}

	for _, coin := range depositAmount {
		if _, ok := acceptedDenoms[coin.Denom]; !ok {
			return errors.Wrapf(
				types.ErrInvalidDepositDenom,
				"deposited %s, but tunnel accepts only the following denom(s): %v",
				depositAmount,
				denoms,
			)
		}
	}

	return nil
}

// AddDeposit adds a deposit to a tunnel
func (k Keeper) AddDeposit(
	ctx sdk.Context,
	tunnelID uint64,
	depositorAddr sdk.AccAddress,
	depositAmount sdk.Coins,
) error {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		return err
	}

	if err := k.validateDepositDenom(ctx, depositAmount); err != nil {
		return err
	}

	// Transfer the deposit from the depositor to the tunnel module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		depositorAddr,
		types.ModuleName,
		depositAmount,
	); err != nil {
		return err
	}

	// Update the depositor's deposit
	deposit, found := k.GetDeposit(ctx, tunnelID, depositorAddr)
	if !found {
		deposit = types.NewDeposit(tunnelID, depositorAddr, depositAmount)
	} else {
		deposit.Amount = deposit.Amount.Add(depositAmount...)
	}
	k.SetDeposit(ctx, deposit)

	// Update the tunnel's total deposit
	tunnel.TotalDeposit = tunnel.TotalDeposit.Add(depositAmount...)
	k.SetTunnel(ctx, tunnel)

	return nil
}

// SetDeposit sets a deposit in the store
func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	ctx.KVStore(k.storeKey).
		Set(types.TunnelDepositStoreKey(deposit.TunnelID, sdk.MustAccAddressFromBech32(deposit.Depositor)), k.cdc.MustMarshal(&deposit))
}

// DeleteDeposit deletes a deposit from the store
func (k Keeper) DeleteDeposit(ctx sdk.Context, deposit types.Deposit) {
	ctx.KVStore(k.storeKey).
		Delete(types.TunnelDepositStoreKey(deposit.TunnelID, sdk.MustAccAddressFromBech32(deposit.Depositor)))
}

// GetDeposit retrieves a deposit by its tunnel ID and depositor address
func (k Keeper) GetDeposit(
	ctx sdk.Context,
	tunnelID uint64,
	depositorAddr sdk.AccAddress,
) (deposit types.Deposit, found bool) {
	bz := ctx.KVStore(k.storeKey).Get(types.TunnelDepositStoreKey(tunnelID, depositorAddr))
	if bz == nil {
		return types.Deposit{}, false
	}

	k.cdc.MustUnmarshal(bz, &deposit)
	return deposit, true
}

// GetDeposits retrieves all deposits for a tunnel
func (k Keeper) GetDeposits(ctx sdk.Context, tunnelID uint64) []types.Deposit {
	var deposits []types.Deposit
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.TunnelDepositsStoreKey(tunnelID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var deposit types.Deposit
		k.cdc.MustUnmarshal(iterator.Value(), &deposit)
		deposits = append(deposits, deposit)
	}

	return deposits
}

// WithdrawDeposit withdraws a deposit from a tunnel
func (k Keeper) WithdrawDeposit(ctx sdk.Context, tunnelID uint64, amount sdk.Coins, withdrawer sdk.AccAddress) error {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		return err
	}

	deposit, found := k.GetDeposit(ctx, tunnelID, withdrawer)
	if !found {
		return types.ErrDepositNotFound
	}

	// Check if the withdrawer has enough deposit
	if deposit.Amount.IsAllLT(amount) {
		return types.ErrInsufficientDeposit
	}

	// Transfer the deposit from the tunnel module account to the withdrawer
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		withdrawer,
		amount,
	); err != nil {
		return err
	}

	// Update the withdrawer's deposit
	deposit.Amount = deposit.Amount.Sub(amount...)
	if deposit.Amount.IsZero() {
		k.DeleteDeposit(ctx, deposit)
	} else {
		k.SetDeposit(ctx, deposit)
	}

	// Update the tunnel's total deposit
	tunnel.TotalDeposit = tunnel.TotalDeposit.Sub(amount...)
	k.SetTunnel(ctx, tunnel)

	// Deactivate the tunnel if the total deposit is less than the min deposit
	minDeposit := k.GetParams(ctx).MinDeposit
	if tunnel.TotalDeposit.IsAllLT(minDeposit) {
		k.DeactivateTunnel(ctx, tunnelID)
	}

	return nil
}
