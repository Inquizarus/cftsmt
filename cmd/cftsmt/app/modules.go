package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/inquizarus/cftsmt/pkg/terraform"
	"github.com/spf13/cobra"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func makeModulesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "modules",
		Short: "Lists all modules from the state in a neat table.",
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

			tbl := table.New("Address", "Module qty", "Resource qty", "Level")
			tbl.WithHeaderFormatter(headerFmt)

			addModuleRows(state.Values.RootModule, tbl, 0)
			addModulesModulesRows(state.Values.RootModule.ChildModules, tbl, 1)
			tbl.Print()
		},
	}
}
