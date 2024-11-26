package main

import (
	"fmt"
	"os"

	"github.com/rudrowo/sqlite/internal/api"
	"github.com/rudrowo/sqlite/internal/sql"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	fileName := os.Args[1]
	userCommand := os.Args[2]

	dbFile := api.Init(fileName)
	defer dbFile.Close()

	// TODO  Remove switch case and use Regex matching
	switch userCommand {
	case ".dbinfo":
		fmt.Printf("database page size: %v\n", api.ReadPageSize())
		fmt.Printf("number of tables: %v", api.CountRows("sqlite_schema"))
	case ".tables":
		fmt.Print(sql.ExecuteSelect("SELECT tbl_name FROM sqlite_schema WHERE tbl_name != 'sqlite_sequence'"))
	default:
		fmt.Println("Unknown command", userCommand)
		os.Exit(1)
	}
}
