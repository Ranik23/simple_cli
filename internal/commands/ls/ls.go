package ls

import (
	"cli/internal/usecase"
	"github.com/spf13/cobra"
)

func LsCommand(uc usecase.UserOperator) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Short:   "Output the content of the directory",
		Example: "cli ls <directory>",
		RunE: func(cmd *cobra.Command, args[]string) error {
			var path string
			if len(args) == 0 {
				path = "."
			}else {
				path = args[0]
			}
			err := uc.Ls(path)
			
			return err
		},
	}
	return cmd
}