package main 
import (
	"fmt"
	"os"
	"flag"
	"github.com/marviigrey/todo-cli"

)
const (
	todoFile = ".todos.Json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {
	case *add:
		todos.Add("sample todo")
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}