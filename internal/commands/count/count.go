package count

import (
	"cli/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

func CountCommand(uc *usecase.UserOperator) *cobra.Command {
	cmd := &cobra.Command{
		Use:		"count",
		Short:  	"Count all the words in the file",
		Example: 	"count file.txt",
		Args: 		cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			filepath := args[0]
			
			n, err := uc.CountWords(filepath, -1, -1)
			
			if err == nil {
				fmt.Println(n)
			}

			return err
		},	
	}

	return cmd
}