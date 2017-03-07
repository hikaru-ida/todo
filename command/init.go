package command

import (
	"database/sql"
	"log"
	_ "os"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/mitchellh/cli"
)

type InitCommand struct {
}

func (c *InitCommand) Run(args []string) int {
	//isDatabaseExists := exists(dbPath())

	/*if isDatabaseExists {
		println("Database is alreadey exists")
		return 0
	    }
	*/
	//os.Remove(dbPath())

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	println(dbPath())

	sqlStmt := `
	CREATE TABLE todos (
	    id integer NOT NULL PRIMARY KEY, 
	    title text NOT NULL, 
	    is_done integer NOT NULL, 
	    created_at integer NOT NULL, 
	    updated_at integer NOT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return 0
	}
	return 0
}

func (c *InitCommand) Synopsis() string {
	return "Init database for todo"
}

func (c *InitCommand) Help() string {
	helpText := "Usage: todo add [option]"
	return helpText
}
