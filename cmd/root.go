package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// running the app without a subcommand just displays the help message
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Serves information on HtN users",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
