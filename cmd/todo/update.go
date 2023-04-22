package todo

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vibovenkat123/todo/pkg/helpers"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"change", "mv", "chng", "move"},
	Short:   "Update a todo",
	Args:    cobra.MinimumNArgs(2),
	Run:     update,
}

func update(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprint(os.Stderr, "Enter a valid ID (integer)\n")
		os.Exit(1)
	}
	args = args[1:]
	name := strings.Join(args, " ")
	err = helpers.Application.UpdateTodo(name, id)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
