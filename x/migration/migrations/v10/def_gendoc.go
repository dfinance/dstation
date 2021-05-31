package v10

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/dfinance/dstation/app"
	"github.com/dfinance/dstation/cmd/dstation/config"
	tmTypes "github.com/tendermint/tendermint/types"
)

// buildNewDefaultGenDoc builds a new tmTypes.GenesisDoc with defaults and Dfinance overwrites.
func buildNewDefaultGenDoc(cdc codec.Marshaler, chainId string, genTime time.Time) (*tmTypes.GenesisDoc, error) {
	appState, err := config.SetGenesisDefaults(cdc, app.NewDefaultGenesisState())
	if err != nil {
		return nil, fmt.Errorf("appState: overwrite defaults: %w", err)
	}
	appStateBz, err := json.MarshalIndent(appState, "", " ")
	if err != nil {
		return nil, fmt.Errorf("appState: JSON marshal: %w", err)
	}

	genDoc := &tmTypes.GenesisDoc{
		GenesisTime:     genTime,
		ChainID:         chainId,
		InitialHeight:   0,
		ConsensusParams: nil,
		Validators:      nil,
		AppHash:         nil,
		AppState:        appStateBz,
	}
	if err := genDoc.ValidateAndComplete(); err != nil {
		return nil, fmt.Errorf("genDoc: ValidateAndComplete: %w", err)
	}

	consParams, err := config.SetConsensusDefaults(genDoc.ConsensusParams)
	if err != nil {
		return nil, fmt.Errorf("consesnsus params: overwrite defaults: %w", err)
	}
	genDoc.ConsensusParams = consParams

	return genDoc, nil
}
