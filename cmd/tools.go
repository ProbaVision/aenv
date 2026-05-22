package cmd

import (
	"github.com/ProbaVision/aenv/internal/toolcheck"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Check commonly used agent CLI tools",
	RunE: func(cmd *cobra.Command, args []string) error {
		tools, err := toolcheck.CheckTools()
		if err != nil {
			return err
		}
		out, err := yaml.Marshal(tools)
		if err != nil {
			return err
		}
		_, err = cmd.OutOrStdout().Write(out)
		return err
	},
}
