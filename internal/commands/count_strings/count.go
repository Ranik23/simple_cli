package count_strings

import (
	"cli/internal/usecase"
	"fmt"

	"github.com/spf13/cobra"
)

func CountCommand2(uc *usecase.UserOperator) *cobra.Command {
	cmd := &cobra.Command{
		Use:		"countS",
		Short:  	"Count all the words in the file",
		Example: 	"count file.txt",
		Args: 		cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			filepath := args[0]
			
			n, err := uc.CountStrings(filepath)
			
			if err == nil {
				fmt.Println(n)
			}

			return err
		},	
	}

	return cmd
}