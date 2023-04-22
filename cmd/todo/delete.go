package todo

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vibovenkat123/todo/pkg/helpers"
)

var deleteCmd = &cobra.Command{
	Use:     "finish",
	Aliases: []string{"delete", "remove", "del", "rm"},
	Short:   "removes a todo",
	Args: cobra.ExactArgs(1),
	Run:     delete,
}

func delete(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, "Enter a valid ID (integer)\n")
		os.Exit(1)
	}
	err = helpers.Application.DeleteTodo(id)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
