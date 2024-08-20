package cmd

import (
	"github.com/mauroandocilla/cli-squeezy-pix/help"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spix",
	Short: help.RootShortHelpText,
	Long:  help.RootHelpText,
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands to the root command
	rootCmd.AddCommand(OptimizeCmd)
	rootCmd.AddCommand(ConvertCmd)
	rootCmd.AddCommand(ResponsiveCmd)

	// Add help template to the root command
	rootCmd.SetHelpTemplate(help.HelpTemplate)
}
