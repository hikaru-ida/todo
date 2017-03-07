package command

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type AddCommand struct {
}

func (c *AddCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("[ERROR] Must set task")
		return 0
	}
	title := strings.Join(args, " ")
	fmt.Println("Added new task : ", title)

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err = db.Exec("INSERT INTO todos(title, is_done, created_at, updated_at) values('" + title + "', 0, " + now + ",  " + now + ")")
	if err != nil {
		log.Fatal(err)
	}

	return 0
}

func (c *AddCommand) Synopsis() string {
	return ""
}

func (c *AddCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
