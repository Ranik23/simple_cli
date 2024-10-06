package clear

import (
	"cli/internal/usecase"
	"github.com/spf13/cobra"
)

func ClearCommand(uc *usecase.UserOperator) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clear",
		Short:   "Clear the terminal",
		Example: "cli clear",
		Run: func(cmd *cobra.Command, args []string) {
			uc.Clear()
		},
	}

	return cmd
}