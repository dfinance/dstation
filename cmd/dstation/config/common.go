package config

import (
	"encoding/json"
	"fmt"
	stdlog "log"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/dfinance/dstation/app"
)

// Chain defaults
const (
	MainDenom = "xfi" // 12 decimals

	// Min TX fee
	FeeAmount = "100000000000000" // 0.0001
	// Governance: deposit amount
	GovMinDepositAmount = "1000000000000000000000" // 1000.0
	// Crisis: invariants check TX fee
	InvariantCheckAmount = "1000000000000000000000" // 1000.0

	MaxGas = 10000000
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	FeeCoin            sdk.Coin
	GovMinDepositCoin  sdk.Coin
	InvariantCheckCoin sdk.Coin
)

// SetGenesisDefaults takes default app genesis state and overwrites Cosmos SDK / Dfinance params.
func SetGenesisDefaults(cdc codec.Marshaler, appStateBz json.RawMessage) (json.RawMessage, error) {
	var genState app.GenesisState
	if err := json.Unmarshal(appStateBz, &genState); err != nil {
		return nil, fmt.Errorf("appStateBz json unmarshal: %w", err)
	}

	// Bank module genesis
	{
		moduleName, moduleState := bankTypes.ModuleName, bankTypes.GenesisState{}
		if err := cdc.UnmarshalJSON(genState[moduleName], &moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON unmarshal: %v", moduleName, err)
		}

		moduleState.DenomMetadata = []bankTypes.Metadata{
			{
				Base:        "axfi",
				Display:     "xfi",
				Description: "XFI currency: native staking and rewards token",
				DenomUnits: []*bankTypes.DenomUnit{
					{
						Denom:    "axfi",
						Exponent: 0,
						Aliases: []string{
							"attoxfi",
						},
					},
					{
						Denom:    "cxfi",
						Exponent: 2,
						Aliases: []string{
							"centixfi",
						},
					},
					{
						Denom:    "mxfi",
						Exponent: 3,
						Aliases: []string{
							"millixfi",
						},
					},
					{
						Denom:    "uxfi",
						Exponent: 6,
						Aliases: []string{
							"microxfi",
						},
					},
					{
						Denom:    "nxfi",
						Exponent: 9,
						Aliases: []string{
							"nanoxfi",
						},
					},
					{
						Denom:    "pxfi",
						Exponent: 12,
						Aliases: []string{
							"picoxfi",
						},
					},
					{
						Denom:    "xfi",
						Exponent: 18,
						Aliases: []string{
							"xfi",
						},
					},
				},
			},
			{
				Base:        "aeth",
				Display:     "eth",
				Description: "Ethereum currency",
				DenomUnits: []*bankTypes.DenomUnit{
					{
						Denom:    "aeth",
						Exponent: 0,
						Aliases: []string{
							"attoeth",
						},
					},
					{
						Denom:    "ceth",
						Exponent: 2,
						Aliases: []string{
							"centieth",
						},
					},
					{
						Denom:    "meth",
						Exponent: 3,
						Aliases: []string{
							"millieth",
						},
					},
					{
						Denom:    "ueth",
						Exponent: 6,
						Aliases: []string{
							"microeth",
						},
					},
					{
						Denom:    "neth",
						Exponent: 9,
						Aliases: []string{
							"nanoeth",
						},
					},
					{
						Denom:    "peth",
						Exponent: 12,
						Aliases: []string{
							"picoeth",
						},
					},
					{
						Denom:    "eth",
						Exponent: 18,
						Aliases: []string{
							"eth",
						},
					},
				},
			},
			{
				Base:        "satoshi",
				Display:     "btc",
				Description: "BitCoin currency",
				DenomUnits: []*bankTypes.DenomUnit{
					{
						Denom:    "satoshi",
						Exponent: 0,
						Aliases: []string{
							"satoshi",
						},
					},
					{
						Denom:    "cbtc",
						Exponent: 2,
						Aliases: []string{
							"centibtc",
						},
					},
					{
						Denom:    "mbtc",
						Exponent: 3,
						Aliases: []string{
							"millibtc",
						},
					},
					{
						Denom:    "ubtc",
						Exponent: 6,
						Aliases: []string{
							"microbtc",
						},
					},
					{
						Denom:    "btc",
						Exponent: 8,
						Aliases: []string{
							"btc",
						},
					},
				},
			},
			{
				Base:        "uusdt",
				Display:     "usdt",
				Description: "USD currency",
				DenomUnits: []*bankTypes.DenomUnit{
					{
						Denom:    "uusdt",
						Exponent: 0,
						Aliases: []string{
							"uusdt",
						},
					},
					{
						Denom:    "cusdt",
						Exponent: 2,
						Aliases: []string{
							"centiusdt",
						},
					},
					{
						Denom:    "musdt",
						Exponent: 3,
						Aliases: []string{
							"milliusdt",
						},
					},
					{
						Denom:    "usdt",
						Exponent: 6,
						Aliases: []string{
							"usdt",
						},
					},
				},
			},
		}
		moduleState.Params.SendEnabled = []*bankTypes.SendEnabled{
			{
				Denom:   "xfi",
				Enabled: true,
			},
		}

		if moduleStateBz, err := cdc.MarshalJSON(&moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON marshal: %v", moduleName, err)
		} else {
			genState[moduleName] = moduleStateBz
		}
	}

	// Mint module params
	{
		moduleName, moduleState := mintTypes.ModuleName, mintTypes.GenesisState{}
		if err := cdc.UnmarshalJSON(genState[moduleName], &moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON unmarshal: %v", moduleName, err)
		}
		moduleState.Params.MintDenom = MainDenom

		if moduleStateBz, err := cdc.MarshalJSON(&moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON marshal: %v", moduleName, err)
		} else {
			genState[moduleName] = moduleStateBz
		}
	}

	// Staking module params
	{
		moduleName, moduleState := stakingTypes.ModuleName, stakingTypes.GenesisState{}
		if err := cdc.UnmarshalJSON(genState[moduleName], &moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON unmarshal: %v", moduleName, err)
		}

		moduleState.Params.BondDenom = MainDenom

		if moduleStateBz, err := cdc.MarshalJSON(&moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON marshal: %v", moduleName, err)
		} else {
			genState[moduleName] = moduleStateBz
		}
	}

	// Crisis module params
	{
		moduleName, moduleState := crisisTypes.ModuleName, crisisTypes.GenesisState{}
		if err := cdc.UnmarshalJSON(genState[moduleName], &moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON unmarshal: %v", moduleName, err)
		}

		moduleState.ConstantFee.Denom = InvariantCheckCoin.Denom   // xfi
		moduleState.ConstantFee.Amount = InvariantCheckCoin.Amount // 1000.0

		if moduleStateBz, err := cdc.MarshalJSON(&moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON marshal: %v", moduleName, err)
		} else {
			genState[moduleName] = moduleStateBz
		}
	}

	// Gov module params
	{
		moduleName, moduleState := govTypes.ModuleName, govTypes.GenesisState{}
		if err := cdc.UnmarshalJSON(genState[moduleName], &moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON unmarshal: %v", moduleName, err)
		}

		moduleState.DepositParams.MinDeposit = sdk.NewCoins(GovMinDepositCoin) // 1000.0xfi

		if moduleStateBz, err := cdc.MarshalJSON(&moduleState); err != nil {
			return nil, fmt.Errorf("%s module: JSON marshal: %v", moduleName, err)
		} else {
			genState[moduleName] = moduleStateBz
		}
	}

	genStateBz, err := json.MarshalIndent(genState, "", " ")
	if err != nil {
		return nil, fmt.Errorf("genState json marshal: %w", err)
	}

	return genStateBz, nil
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		stdlog.Printf("Failed to get home dir: %v\n", err)
	}
	DefaultNodeHome = filepath.Join(userHomeDir, ".dstation")

	if value, ok := sdk.NewIntFromString(FeeAmount); !ok {
		panic("defaults: FeeAmount conversion failed")
	} else {
		FeeCoin = sdk.NewCoin(MainDenom, value)
	}

	if value, ok := sdk.NewIntFromString(GovMinDepositAmount); !ok {
		panic("governance defaults: GovMinDepositAmount conversion failed")
	} else {
		GovMinDepositCoin = sdk.NewCoin(MainDenom, value)
	}

	if value, ok := sdk.NewIntFromString(InvariantCheckAmount); !ok {
		panic("crisis defaults: InvariantCheckAmount conversion failed")
	} else {
		InvariantCheckCoin = sdk.NewCoin(MainDenom, value)
	}
}
