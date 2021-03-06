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

func makeOutputsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "outputs",
		Short: "Lists all outputs from the state in a neat table.",
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

			tbl := table.New("Name", "Sensitive", "Value")
			tbl.WithHeaderFormatter(headerFmt)

			for k, v := range state.Values.Outputs {
				tbl.AddRow(k, v.Sensitive, v.Value)
			}

			tbl.Print()
		},
	}
}
