package command

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type DoneCommand struct {
}

func (c *DoneCommand) Run(args []string) int {
	if len(args) != 1 {
		return 0
	}

	id := args[0]
	fmt.Printf("Task %s is done\n", id)

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET is_done = 1 WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}
	return 0
}

func (c *DoneCommand) Synopsis() string {
	return ""
}

func (c *DoneCommand) Help() string {
	helpText := ""
	return strings.TrimSpace(helpText)
}
