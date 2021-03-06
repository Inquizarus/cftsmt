package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inquizarus/cftsmt/pkg/formatting"
	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fatih/color"
)

func makeResourceStatementsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "statements",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := os.Stdin.Stat()
			if err != nil {
				panic(err)
			}

			if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
				fmt.Println("The command is intended to work with pipes.")
				return
			}

			reader := bufio.NewReader(os.Stdin)

			state, _ := terraform.FromReader(reader)

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

			address := viper.GetString(ArgResourceAddress)

			resources := []terraform.StateResource{}

			if "" != address {
				resource := terraform.FindResourceInModule(address, state.Values.RootModule)
				if nil == resource {
					fmt.Println(color.RedString("Resource with address %s was not found, aborting.", address))
					return
				}
				resources = append(resources, *resource)
			}

			if "" == address {
				depth := viper.GetInt(ArgDepth)
				filter := createResourceFilter(viper.GetViper())
				resources = terraform.FindResourcesInModule(state.Values.RootModule, depth, filter)
			}

			f := formatting.NewFormatter(map[string]formatting.Provider{})
			f.SetProvider(&formatting.DefaultProvider{})

			tbl := table.New("Remove", "Import")
			tbl.WithHeaderFormatter(headerFmt)

			for _, r := range resources {
				tbl.AddRow(f.Remove(r), f.Import(r))
			}

			tbl.Print()
		},
	}
	cmd.PersistentFlags().StringVarP(&resourceAddress, ArgResourceAddress, ShortArgMap[ArgResourceAddress], "", ArgDescriptionMap[ArgResourceAddress])
	cmd.Flags().StringVarP(&modeFilter, ArgModeFilter, ShortArgMap[ArgModeFilter], ArgDefaultValueMap[ArgModeFilter].(string), ArgDescriptionMap[ArgModeFilter])
	cmd.Flags().StringVarP(&typeFilter, ArgTypeFilter, ShortArgMap[ArgTypeFilter], ArgDefaultValueMap[ArgTypeFilter].(string), ArgDescriptionMap[ArgTypeFilter])
	cmd.Flags().StringVarP(&valuesFilter, ArgValuesFilter, ShortArgMap[ArgValuesFilter], ArgDefaultValueMap[ArgValuesFilter].(string), ArgDescriptionMap[ArgValuesFilter])
	cmd.Flags().IntVarP(&depth, ArgDepth, ShortArgMap[ArgDepth], ArgDefaultValueMap[ArgDepth].(int), ArgDescriptionMap[ArgDepth])

	viper.BindPFlag(ArgResourceAddress, cmd.PersistentFlags().Lookup(ArgResourceAddress))
	viper.BindPFlag(ArgModeFilter, cmd.Flags().Lookup(ArgModeFilter))
	viper.BindPFlag(ArgTypeFilter, cmd.Flags().Lookup(ArgTypeFilter))
	viper.BindPFlag(ArgValuesFilter, cmd.Flags().Lookup(ArgValuesFilter))
	viper.BindPFlag(ArgDepth, cmd.Flags().Lookup(ArgDepth))

	return cmd
}
