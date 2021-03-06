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

			resources := terraform.FindResourcesInModule(state.Values.RootModule, depth)

			addResourceRows(resources, tbl)

			tbl.Print()
		},
	}
	cmd.Flags().IntVarP(&depth, ArgDepth, ShortArgMap[ArgDepth], ArgDefaultValueMap[ArgDepth].(int), ArgDescriptionMap[ArgDepth])
	viper.BindPFlag(ArgDepth, cmd.Flags().Lookup(ArgDepth))
	cmd.AddCommand(makeResourcesValuesCommand())
	cmd.AddCommand(makeResourceStatementsCommand())
	return cmd
}
