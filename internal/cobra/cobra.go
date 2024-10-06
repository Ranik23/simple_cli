package cobra

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use:   "cli",
		Short: "Command Line Interface",
		Long:  "Command Line Interface. Commands: cat, count",
	}
)