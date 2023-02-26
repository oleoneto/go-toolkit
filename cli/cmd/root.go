package cmd

import (
	"os"

	"github.com/oleoneto/go-toolkit/cli/cmd/postman"
	"github.com/oleoneto/go-toolkit/logger"
	"github.com/spf13/cobra"
)

var (
	format   = "plain"
	template = ""

	rootCmd = &cobra.Command{
		Use:   "go-toolkit",
		Short: "go-toolkit, a cli tool...",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the version of the CLI",
		Run: func(cmd *cobra.Command, args []string) {
			logg := logger.NewLogger(logger.LoggerOptions{Format: format})
			logg.Log(&version, os.Stdout, template)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {}

func init() {
	// CLI configuration
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&format, "output-format", "o", format, "output format")
	rootCmd.PersistentFlags().StringVarP(&template, "output-template", "y", template, "template (used when output format is 'gotemplate')")

	// Commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(postman.Cmd)
}
