package v10

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	v076 "github.com/dfinance/dstation/x/migration/migrations/v076"
	"github.com/gogo/protobuf/proto"
)

// migrateAuthModule migrates old auth module GenesisState to the new auth module.
// Transfers only BaseAccounts, skipping ModuleAccounts.
func migrateAuthModule(cdc codec.Marshaler, oldStateBz, newStateBz json.RawMessage) (proto.Message, error) {
	oldCodec := codec.NewLegacyAmino()
	v076.RegisterAuthLegacyAminoCodec(oldCodec)

	// Input checks
	if cdc == nil {
		return nil, fmt.Errorf("cdc: nil")
	}
	if oldStateBz == nil {
		return nil, fmt.Errorf("oldStateBz: nil")
	}
	if newStateBz == nil {
		return nil, fmt.Errorf("newStateBz: nil")
	}

	// Unmarshal states
	oldState := v076.AuthGenesisState{}
	if err := oldCodec.UnmarshalJSON(oldStateBz, &oldState); err != nil {
		return nil, fmt.Errorf("old state: legacy codec JSON unmarshal: %w", err)
	}

	newState := &authTypes.GenesisState{}
	if err := cdc.UnmarshalJSON(newStateBz, newState); err != nil {
		return nil, fmt.Errorf("new state: proto codec JSON unmarshal: %w", err)
	}

	// Migration routine
	newAccs := make([]authTypes.GenesisAccount, 0, len(oldState.Accounts))
	for _, oldAccRaw := range oldState.Accounts {
		switch oldAcc := oldAccRaw.(type) {
		case *v076.ModuleAccount:
			continue
		case *v076.BaseAccount:
			newAcc := authTypes.NewBaseAccount(oldAcc.Address, oldAcc.PubKey, oldAcc.AccountNumber, oldAcc.Sequence)
			newAccs = append(newAccs, newAcc)
		default:
			return nil, fmt.Errorf("account (%s): unsupported account type (%T)", oldAccRaw.GetAddress(), oldAcc)
		}
	}

	// Convert []authTypes.GenesisAccount into []proto.Any
	newAccAnys := make([]*codecTypes.Any, 0, len(newAccs))
	for _, newAcc := range newAccs {
		any, err := codecTypes.NewAnyWithValue(newAcc)
		if err != nil {
			return nil, fmt.Errorf("account (%s): converting to proto.Any: %w", newAcc.GetAddress(), err)
		}
		newAccAnys = append(newAccAnys, any)
	}
	newState.Accounts = newAccAnys

	return newState, nil
}

// migrateBankModule migrates old auth module GenesisState to the new bank module.
// Transfers BaseAccounts balances for "xfi" token.
// Estimates a new supply amount.
func migrateBankModule(cdc codec.Marshaler, oldStateBz, newStateBz json.RawMessage) (proto.Message, error) {
	targetDenoms := []string{"xfi"}

	oldCodec := codec.NewLegacyAmino()
	v076.RegisterAuthLegacyAminoCodec(oldCodec)

	// Input checks
	if cdc == nil {
		return nil, fmt.Errorf("cdc: nil")
	}
	if oldStateBz == nil {
		return nil, fmt.Errorf("oldStateBz: nil")
	}
	if newStateBz == nil {
		return nil, fmt.Errorf("newStateBz: nil")
	}

	// Unmarshal states
	oldState := v076.AuthGenesisState{}
	if err := oldCodec.UnmarshalJSON(oldStateBz, &oldState); err != nil {
		return nil, fmt.Errorf("old state: legacy codec JSON unmarshal: %w", err)
	}

	newState := &bankTypes.GenesisState{}
	if err := cdc.UnmarshalJSON(newStateBz, newState); err != nil {
		return nil, fmt.Errorf("new state: proto codec JSON unmarshal: %w", err)
	}

	// Migration routine
	newBalances := make([]bankTypes.Balance, 0)
	newSupply := sdk.NewCoins()
	for _, oldAccRaw := range oldState.Accounts {
		switch oldAcc := oldAccRaw.(type) {
		case *v076.BaseAccount:
			coins := sdk.NewCoins()
			for _, targetDenom := range targetDenoms {
				amt := oldAcc.Coins.AmountOf(targetDenom)
				if !amt.IsZero() {
					coins = coins.Add(sdk.NewCoin(targetDenom, amt))
				}
			}
			if coins.IsZero() {
				continue
			}

			newBalances = append(newBalances, bankTypes.Balance{
				Address: oldAcc.GetAddress().String(),
				Coins:   coins,
			})
			newSupply = newSupply.Add(coins...)
		default:
		}
	}
	newState.Balances = newBalances
	newState.Supply = newSupply

	return newState, nil
}
