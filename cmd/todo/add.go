package todo

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vibovenkat123/todo/pkg/helpers"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "touch", "make", "mk"},
	Short:   "Create a todo",
	Args: cobra.MinimumNArgs(1),
	Run:     add,
}

func add(cmd *cobra.Command, args []string) {
	name := strings.Join(args, " ")
	err := helpers.Application.InsertTodo(name)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
