package cli

import (
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/spf13/cobra"
)

// GetCmdTxGenAddNominee returns genesis setup command.
func GetCmdTxGenAddNominee() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen-add-nominee [accAddress]",
		Short:   "Add x/staker module nominee account to genesis.json (no validation)",
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
