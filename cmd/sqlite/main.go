package main

import (
	"fmt"
	"os"

	"github.com/rudrowo/sqlite/internal/api"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	dbFile := api.Init()
	defer dbFile.Close()
	userCommand := os.Args[2]

	// TODO  Remove switch case and call HandleDotCommands() and HandleSelectQuery()
	switch userCommand {
	case ".dbinfo":
		fmt.Printf("database page size: %v\n", api.ReadPageSize())
		fmt.Printf("number of tables: %v", api.CountRows("sqlite_schema"))
	default:
		fmt.Println("Unknown command", userCommand)
		os.Exit(1)
	}
}
