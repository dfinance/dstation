package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

// GetVmCurrencyInfo implements types.CurrencyInfoResProvider interface.
func (k Keeper) GetVmCurrencyInfo(ctx sdk.Context, denom string) *dvmTypes.CurrencyInfo {
	// Check if cached
	if ccInfo, found := k.cache.currencyInfo[denom]; found {
		return k.enrichCurrencyInfoWithTotalSupply(ctx, ccInfo)
	}

	// Request from the Bank and cache the value
	_, denomUnit, found := k.getDenomMetadataFromBank(ctx, denom)
	if !found {
		return nil
	}

	ccInfo := dvmTypes.CurrencyInfo{
		Denom:       []byte(denomUnit.Denom),
		Decimals:    denomUnit.Exponent,
		IsToken:     false,
		Address:     types.StdLibAddress,
		TotalSupply: nil,
	}
	k.cache.currencyInfo[denom] = ccInfo

	return k.enrichCurrencyInfoWithTotalSupply(ctx, ccInfo)
}

// getDenomMetadataFromBank iterates over Bank denom metadata entries looking for specified denom-unit.
func (k Keeper) getDenomMetadataFromBank(ctx sdk.Context, denom string) (retMeta *bankTypes.Metadata, retUnit *bankTypes.DenomUnit, found bool) {
	k.bankKeeper.IterateAllDenomMetaData(ctx, func(metadata bankTypes.Metadata) bool {
		for _, unit := range metadata.DenomUnits {
			if unit == nil || unit.Denom != denom {
				continue
			}

			if unit.Denom == denom {
				retMeta, retUnit, found = &metadata, unit, true
			}

			return true
		}

		return false
	})

	return
}

// enrichCurrencyInfoWithTotalSupply sets the TotalSupply field of dvmTypes.CurrencyInfo.
func (k Keeper) enrichCurrencyInfoWithTotalSupply(ctx sdk.Context, ccInfo dvmTypes.CurrencyInfo) *dvmTypes.CurrencyInfo {
	denom := string(ccInfo.Denom)
	supply := k.bankKeeper.GetSupply(ctx).GetTotal().AmountOf(denom)

	value, err := types.SdkIntToVmU128(supply)
	if err != nil {
		panic(fmt.Errorf("converting totalSupply (%s) for denom (%s) to VM U128 type: %w", supply.String(), denom, err))
	}
	ccInfo.TotalSupply = value

	return &ccInfo
}
