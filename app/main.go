package main

import (
	"encoding/binary"
	"fmt"
	"os"

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
		var pageSize uint16
		// var numberOfTables uint16

		pageSize = binary.BigEndian.Uint16(dbHeader[16:18])

		fmt.Printf("database page size: %v", pageSize)
	default:
		fmt.Println("Unknown command", userCommand)
		os.Exit(1)
	}
}
