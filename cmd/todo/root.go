package todo

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLi to manage your todos",
	Long: `A simple todo manager using sqlite3
that allows you to manage all your todos
in the command line`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running todo: %s", err)
		os.Exit(1)
	}
}
