package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fatih/color"
)

func makeResourcesValuesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "values",
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

			resource := terraform.FindResourceInModule(address, state.Values.RootModule)

			if nil == resource {
				fmt.Println(color.RedString("Resource with address %s was not found, aborting.", address))
				return
			}

			tbl := table.New("Key", "Value")
			tbl.WithHeaderFormatter(headerFmt)

			for k, v := range resource.Values {
				tbl.AddRow(k, v)
			}

			tbl.Print()
		},
	}
	cmd.PersistentFlags().StringVarP(&resourceAddress, ArgResourceAddress, ShortArgMap[ArgResourceAddress], "", ArgDescriptionMap[ArgResourceAddress])
	cmd.MarkFlagRequired(ArgResourceAddress)
	viper.BindPFlag(ArgResourceAddress, cmd.PersistentFlags().Lookup(ArgResourceAddress))
	return cmd
}
