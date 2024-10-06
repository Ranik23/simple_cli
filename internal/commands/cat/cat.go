package read

import (
	"cli/internal/usecase"
	
	"github.com/spf13/cobra"
)


func CatCommand(uc *usecase.UserOperator) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cat",
		Short:   "Output the content of the file",
		Example: "cli cat file.txt",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath := args[0]
			err := uc.Print(filePath)
			return err
		},
	}

	return cmd
}