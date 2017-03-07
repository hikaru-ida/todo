package command

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type ListCommand struct {
}

func (c *ListCommand) Run(args []string) int {
	var all bool
	flags := flag.NewFlagSet("list", flag.ContinueOnError)
	flags.BoolVar(&all, "all", false, "show all task")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	var s string
	if all {
		s = "SELECT id, title, is_done, created_at FROM todos"
	} else {
		s = "SELECT id, title, is_done, created_at FROM todos WHERE is_done = 0"
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data := [][]string{}

	for rows.Next() {
		var id int
		var title string
		var isDone int
		var createdAt int64
		rows.Scan(&id, &title, &isDone, &createdAt)
		data = append(data, []string{strconv.Itoa(id), title, doneLabel(isDone), dateForView(createdAt)})
	}

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"NO", "Title", "status", "Created"})
		table.SetBorder(true)
		table.AppendBulk(data)
		table.Render()
	}

	return 0
}

func doneLabel(isDone int) string {
	if isDone == 0 {
		return "-"
	}
	return "Done"
}

func dateForView(at int64) string {
	return time.Unix(at, 0).Format("2006-01-02 15:04:05")
}

func (c *ListCommand) Synopsis() string {
	return ""
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
