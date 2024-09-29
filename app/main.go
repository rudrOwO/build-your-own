package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"github/com/codecrafters-io/sqlite-starter-go/app/features"
	u "github/com/codecrafters-io/sqlite-starter-go/app/utils"
	// Available if you need it!
	// "github.com/xwb1989/sqlparser"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	dbFile, err := os.Open(os.Args[1])
	u.HandleError(err)
	defer dbFile.Close()

	userCommand := os.Args[2]
	dbHeader := make([]byte, 100)
	_, err = dbFile.Read(dbHeader)
	u.HandleError(err)

	switch userCommand {
	case ".dbinfo":
		pageSize := binary.BigEndian.Uint16(dbHeader[16:18])
		numberOfTables := features.CountRows("sqlite_schema", dbFile)

		fmt.Printf("database page size: %v\n", pageSize)
		fmt.Printf("number of tables: %v", numberOfTables)
	default:
		fmt.Println("Unknown command", userCommand)
		os.Exit(1)
	}
}
