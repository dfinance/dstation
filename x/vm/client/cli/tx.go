package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	govCli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cobra"

	"github.com/dfinance/dstation/pkg"
	vmClient "github.com/dfinance/dstation/x/vm/client"
	"github.com/dfinance/dstation/x/vm/types"
)

// GetTxCmd returns a root CLI command handler for all module transaction commands.
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "VM transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdTxExecuteScript(),
		GetCmdTxDeployContract(),
		GetCmdTxSendUpdateStdlibProposal(),
	)

	return txCmd
}

// GetCmdTxExecuteScript returns tx command that implement keeper handler.
func GetCmdTxExecuteScript() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "execute [moveFile] [arg1...]",
		Short:   "Execute Move script",
		Example: "execute ./script.move.json wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 100 true \"my string\" \"68656c6c6f2c20776f726c6421\" #\"XFI_ETH\" --from my_account --fees 10000xfi --gas 500000",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			compItems, err := getCompiledItemFromFileArg("moveFile", args[0], true)
			if err != nil {
				return err
			}

			// Extract script arguments meta
			meta, err := queryClient.Metadata(cmd.Context(), &types.QueryMetadataRequest{Code: compItems.CompiledItems[0].ByteCode})
			if err != nil {
				return fmt.Errorf("extracting script arguments meta: %w", err)
			}
			if meta.Metadata.GetScript() == nil {
				return fmt.Errorf("extracting script arguments meta: requested byteCode is not a script")
			}
			typedArgs := meta.Metadata.GetScript().Arguments

			// Build msg
			scriptArgs, err := vmClient.ConvertStringScriptArguments(args[1:], typedArgs)
			if err != nil {
				return fmt.Errorf("converting input args to typed args: %w", err)
			}
			msg := types.NewMsgExecuteScript(fromAddr, compItems.CompiledItems[0].ByteCode, scriptArgs...)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	pkg.BuildCmdHelp(cmd, []string{
		"path to compiled Move file containing byteCode (one script)",
		"space separated VM script arguments (optional)",
	})

	return cmd
}

// GetCmdTxDeployContract returns tx command that implement keeper handler.
func GetCmdTxDeployContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "publish [moveFile]",
		Short:   "Publish Move module",
		Example: "publish ./my_module.move.json --from my_account --fees 10000xfi --gas 500000",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			compItems, err := getCompiledItemFromFileArg("moveFile", args[0], true)
			if err != nil {
				return err
			}

			// Build msg
			contractsCode := make([][]byte, 0, len(compItems.CompiledItems))
			for _, item := range compItems.CompiledItems {
				contractsCode = append(contractsCode, item.ByteCode)
			}
			msg := types.NewMsgDeployModule(fromAddr, contractsCode...)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	pkg.BuildCmdHelp(cmd, []string{
		"path to compiled Move file containing byteCode (one / several modules)",
	})

	return cmd
}

// GetCmdTxSendUpdateStdlibProposal returns tx command that sends governance update stdlib VM module proposal.
func GetCmdTxSendUpdateStdlibProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-stdlib-proposal [moveFile] [plannedBlockHeight] [sourceUrl] [updateDescription]",
		Short:   "Submit a DVM stdlib update proposal",
		Example: "update-stdlib-proposal ./update.move.json 1000 http://github.com/repo 'fix for Foo module' --deposit 10000xfi --from my_account --fees 10000xfi",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			compItems, err := getCompiledItemFromFileArg("moveFile", args[0], false)
			if err != nil {
				return err
			}

			plannedBlockHeight, err := pkg.ParseInt64Param("plannedBlockHeight", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			sourceUrl, updateDesc := args[2], args[3]

			deposit, err := pkg.ParseDepositFlag(cmd.Flags())
			if err != nil {
				return err
			}

			// Build and validate PlannedProposal
			code := make([][]byte, 0, len(compItems.CompiledItems))
			for _, compItem := range compItems.CompiledItems {
				code = append(code, compItem.ByteCode)
			}

			pProposal, err := types.NewPlannedProposal(plannedBlockHeight, types.NewStdLibUpdateProposal(sourceUrl, updateDesc, code...))
			if err != nil {
				return fmt.Errorf("planned proposal: build: %w", err)
			}
			if err := pProposal.ValidateBasic(); err != nil {
				return fmt.Errorf("planned proposal: validation: %w", err)
			}

			// Build and validate Gov proposal message
			msg, err := govTypes.NewMsgSubmitProposal(pProposal, deposit, fromAddr)
			if err != nil {
				return fmt.Errorf("gov proposal message: build: %w", err)
			}
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("gov proposal message: validation: %w", err)
			}

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(govCli.FlagDeposit, "", "deposit of proposal")
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	_ = cmd.MarkFlagRequired(govCli.FlagDeposit)

	pkg.BuildCmdHelp(cmd, []string{
		"path to compiled Move file containing byteCode (one / several modules)",
		"blockHeight at which update should occur [int]",
		"URL containing proposal source code",
		"proposal description (version, short changelist)",
	})

	return cmd
}

// getCompiledItemFromFileArg reads .move file and performs basic code type checks.
func getCompiledItemFromFileArg(argName, argValue string, oneItem bool) (*types.QueryCompileResponse, error) {
	jsonContent, err := pkg.ParseFilePath(argName, argValue, pkg.ParamTypeCliArg)
	if err != nil {
		return nil, err
	}

	compItems := types.QueryCompileResponse{}
	if err := json.Unmarshal(jsonContent, &compItems); err != nil {
		return nil, pkg.BuildError(argName, argValue, pkg.ParamTypeCliArg, fmt.Sprintf("file json unmarshal: %v", err))
	}

	if len(compItems.CompiledItems) == 0 || (oneItem && len(compItems.CompiledItems) != 1) {
		return nil, pkg.BuildError(argName, argValue, pkg.ParamTypeCliArg, fmt.Sprintf("file contains wrong number of items: %d", len(compItems.CompiledItems)))
	}

	itemsCodeType := compItems.CompiledItems[0].CodeType
	for _, item := range compItems.CompiledItems {
		if itemsCodeType != item.CodeType {
			return nil, pkg.BuildError(argName, argValue, pkg.ParamTypeCliArg, "file contains different code types (only similar types are allowed)")
		}
	}

	return &compItems, nil
}
