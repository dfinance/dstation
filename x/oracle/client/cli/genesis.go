package cli

import (
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/spf13/cobra"
)

// GetCmdTxGenAddNominee returns genesis setup command.
func GetCmdTxGenAddNominee() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen-add-nominee [accAddress]",
		Short:   "Add x/oracle module nominee account to genesis.json (no validation)",
		Example: "gen-add-nominee wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Parse inputs
			nomineeAddr, err := pkg.ParseSdkAddressParam("accAddress", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			genState := &types.GenesisState{}
			return pkg.UpdateModuleGenesis(cmd, types.ModuleName, genState, func(handlerState interface{}) error {
				moduleState := handlerState.(*types.GenesisState)
				moduleState.Params.Nominees = append(moduleState.Params.Nominees, nomineeAddr.String())
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

			genState := &types.GenesisState{}
			return pkg.UpdateModuleGenesis(cmd, types.ModuleName, genState, func(handlerState interface{}) error {
				moduleState := handlerState.(*types.GenesisState)
				moduleState.Oracles = append(moduleState.Oracles, types.Oracle{
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

			genState := &types.GenesisState{}
			return pkg.UpdateModuleGenesis(cmd, types.ModuleName, genState, func(handlerState interface{}) error {
				moduleState := handlerState.(*types.GenesisState)
				moduleState.Assets = append(moduleState.Assets, types.Asset{
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
