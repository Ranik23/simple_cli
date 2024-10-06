package ls

import (
	"cli/internal/usecase"
	"github.com/spf13/cobra"
)

func LsCommand(uc usecase.UserOperator) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Short:   "Output the content of the current directory",
		Example: "cli ls",
		RunE: func(cmd *cobra.Command, args[]string) error {
			err := uc.Ls()
			return err
		},
	}
	return cmd
}