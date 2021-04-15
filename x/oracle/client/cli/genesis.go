package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/spf13/cobra"
)

// GetCmdTxGenAddNominee returns genesis setup command.
func GetCmdTxGenAddNominee() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen-add-nominee [accAddress]",
		Short:   "Add nominee account to genesis.json (no validation)",
		Example: "gen-add-nominee wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse inputs
			nomineeAddr, err := pkg.ParseSdkAddressParam("accAddress", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			return updateModuleGenesis(cmd, func(genState *types.GenesisState) error {
				genState.Params.Nominees = append(genState.Params.Nominees, nomineeAddr.String())
				return nil
			})
		},
	}

	pkg.BuildCmdHelp(cmd, []string{
		"Nominee account address",
	})

	return cmd
}

// GetCmdTxGenAddOracle returns genesis setup command.
func GetCmdTxGenAddOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen-add-oracle [oracleAccAddress] [priceMaxBytes] [priceDecimals] [sourceDescription]",
		Short:   "Add Oracle source to genesis.json (no validation)",
		Example: "gen-add-oracle wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 8 8 \"Binance\"",
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse inputs
			oracleAddr, err := pkg.ParseSdkAddressParam("oracleAccAddress", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			priceMaxBytes, err := pkg.ParseUint8Param("priceMaxBytes", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			priceDecimals, err := pkg.ParseUint8Param("priceDecimals", args[2], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			oracleDesc := ""
			if len(args) > 3 {
				oracleDesc = args[3]
			}

			return updateModuleGenesis(cmd, func(genState *types.GenesisState) error {
				genState.Oracles = append(genState.Oracles, types.Oracle{
					AccAddress:    oracleAddr.String(),
					Description:   oracleDesc,
					PriceMaxBytes: uint32(priceMaxBytes),
					PriceDecimals: uint32(priceDecimals),
				})
				return nil
			})
		},
	}

	pkg.BuildCmdHelp(cmd, []string{
		"Oracle account address",
		"max number of bytes for ask/bid price values (limit)",
		"number of decimals for ask/bid price values (values send from Oracle are big.Int)",
		"Oracle description (optional)",
	})

	return cmd
}

// GetCmdTxGenAddAsset returns genesis setup command.
func GetCmdTxGenAddAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen-add-asset [assetCode] [rateDecimals] [oracleAddress1...]",
		Short:   "Add Asset to genesis.json (no validation)",
		Example: "gen-add-asset eth_usdt 8 wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8",
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse inputs
			assetCode, err := pkg.ParseAssetCodeParam("assetCode", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			decimals, err := pkg.ParseUint8Param("rateDecimals", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			var oracles []string
			if len(args) > 2 {
				addrs, err := pkg.ParseSpaceSepSdkAddressesParams("oracleAddresses", args[2:], pkg.ParamTypeCliArg)
				if err != nil {
					return err
				}
				for _, addr := range addrs {
					oracles = append(oracles, addr.String())
				}
			}

			return updateModuleGenesis(cmd, func(genState *types.GenesisState) error {
				genState.Assets = append(genState.Assets, types.Asset{
					AssetCode: assetCode,
					Oracles:   oracles,
					Decimals:  uint32(decimals),
				})
				return nil
			})
		},
	}

	pkg.BuildCmdHelp(cmd, []string{
		"asset code symbol (eth_usdt)",
		"number of decimals for ask/bid price values",
		"source (Oracle) account addresses (optional, if none - asset is disabled)",
	})

	return cmd
}

// updateModuleGenesis read, updates module state and saves application genesis state to default genesis.json location.
func updateModuleGenesis(cmd *cobra.Command, stateUpdater func(genState *types.GenesisState) error) error {
	serverCtx := server.GetServerContextFromCmd(cmd)
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	config := serverCtx.Config
	config.SetRoot(clientCtx.HomeDir)

	genFile := config.GenesisFile()
	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return fmt.Errorf("application genesis state: unmarshal failed (%s): %w", genFile, err)
	}

	genesisState := types.GenesisState{}
	if err := json.Unmarshal(appState[types.ModuleName], &genesisState); err != nil {
		return fmt.Errorf("module genesis state (%s): unmarshal failed: %w", types.ModuleName, err)
	}

	if err := stateUpdater(&genesisState); err != nil {
		return fmt.Errorf("module genesis state (%s): update failed: %w", types.ModuleName, err)
	}

	genesisStateBz, err := json.Marshal(genesisState)
	if err != nil {
		return fmt.Errorf("module genesis state (%s): marshal failed: %w", types.ModuleName, err)
	}
	appState[types.ModuleName] = genesisStateBz

	appStateBz, err := json.Marshal(appState)
	if err != nil {
		return fmt.Errorf("application genesis state: marshal failed: %w", err)
	}
	genDoc.AppState = appStateBz

	return genutil.ExportGenesisFile(genDoc, genFile)
}
