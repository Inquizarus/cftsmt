package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func makeResourcesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resources",
		Short: "Lists all resources from the state in a neat table.",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := os.Stdin.Stat()
			if err != nil {
				panic(err)
			}

			if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
				fmt.Println("The command is intended to work with pipes.")
				return
			}

			depth := viper.GetViper().GetInt(ArgDepth)

			reader := bufio.NewReader(os.Stdin)

			state, _ := terraform.FromReader(reader)

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

			tbl := table.New("Type", "Name", "Address", "Mode", "Level")
			tbl.WithHeaderFormatter(headerFmt)

			filter := createResourceFilter(viper.GetViper())

			resources := terraform.FindResourcesInModule(state.Values.RootModule, depth, filter)

			addResourceRows(resources, tbl)

			tbl.Print()
		},
	}
	cmd.Flags().IntVarP(&depth, ArgDepth, ShortArgMap[ArgDepth], ArgDefaultValueMap[ArgDepth].(int), ArgDescriptionMap[ArgDepth])
	cmd.Flags().StringVarP(&modeFilter, ArgModeFilter, ShortArgMap[ArgModeFilter], ArgDefaultValueMap[ArgModeFilter].(string), ArgDescriptionMap[ArgModeFilter])
	cmd.Flags().StringVarP(&typeFilter, ArgTypeFilter, ShortArgMap[ArgTypeFilter], ArgDefaultValueMap[ArgTypeFilter].(string), ArgDescriptionMap[ArgTypeFilter])
	cmd.Flags().StringVarP(&valuesFilter, ArgValuesFilter, ShortArgMap[ArgValuesFilter], ArgDefaultValueMap[ArgValuesFilter].(string), ArgDescriptionMap[ArgValuesFilter])
	viper.BindPFlag(ArgDepth, cmd.Flags().Lookup(ArgDepth))
	viper.BindPFlag(ArgModeFilter, cmd.Flags().Lookup(ArgModeFilter))
	viper.BindPFlag(ArgTypeFilter, cmd.Flags().Lookup(ArgTypeFilter))
	viper.BindPFlag(ArgValuesFilter, cmd.Flags().Lookup(ArgValuesFilter))
	cmd.AddCommand(makeResourcesValuesCommand())
	cmd.AddCommand(makeResourceStatementsCommand())
	return cmd
}
