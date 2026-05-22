package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "aenv",
	Short: "Cockpit view of agent env setup",
	Long:  "aenv provides a quick view of local agent environment setup.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: false,
	},
}

func init() {
	rootCmd.InitDefaultCompletionCmd()
}

func Execute() error {
	return rootCmd.Execute()
}
