package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	// Available if you need it!
	// "github.com/xwb1989/sqlparser"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	dbFilePath := os.Args[1]
	userCommand := os.Args[2]

	dbFile, err := os.Open(dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	dbHeader := make([]byte, 100)

	_, err = dbFile.Read(dbHeader)
	if err != nil {
		log.Fatal(err)
	}

	switch userCommand {
	case ".dbinfo":
		var pageSize uint16
		// var numberOfTables uint16

		if err := binary.Read(bytes.NewReader(dbHeader[16:18]), binary.BigEndian, &pageSize); err != nil {
			fmt.Println("Failed to read integer:", err)
			return
		}

		fmt.Printf("database page size: %v", pageSize)
	default:
		fmt.Println("Unknown command", userCommand)
		os.Exit(1)
	}
}