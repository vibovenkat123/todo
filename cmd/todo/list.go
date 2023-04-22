package todo
import ( "fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vibovenkat123/todo/pkg/helpers"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "lists all the todos",
	Run:     list,
}

func list(cmd *cobra.Command, args []string) {
	todos, err := helpers.Application.GetAllTodos()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("No todos!")
		return
	}
	fmt.Println("Todos:")
	for _, todo := range todos {
		fmt.Printf("%d. %s\n", todo.ID, todo.Name)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
