package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "CFTSMT"
)

var rootCmd = &cobra.Command{
	Use: "cftsmt",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the cli app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(makeOutputsCommand())
	rootCmd.AddCommand(makeResourcesCommand())
	rootCmd.AddCommand(makeModulesCommand())
}

func initConfig() {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
}
