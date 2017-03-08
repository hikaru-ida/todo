package command

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
)

type DeleteCommand struct {
}

func (c *DeleteCommand) Run(args []string) int {
	var all bool
	flags := flag.NewFlagSet("delete", flag.ContinueOnError)
	flags.BoolVar(&all, "all", false, "delete all task")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	var id string
	if len(args) == 1 {
		id = args[0]
	}

	var s string
	var m string
	if all {
		s = "DELETE FROM todos"
		m = "Deleted all tasks"
	} else {
		s = "DELETE FROM todos WHERE id = " + id
		m = "Deleted task id = " + id
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)

	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return ""
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
