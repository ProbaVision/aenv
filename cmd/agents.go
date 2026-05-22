package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/ProbaVision/aenv/internal/checker"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(agentsCmd)
}

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "List local agent environment files",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		entries, err := checker.ScanHome(home)
		if err != nil {
			return err
		}

		out, err := yaml.Marshal(entries)
		if err != nil {
			return err
		}

		_, err = cmd.OutOrStdout().Write(out)
		return err
	},
}
