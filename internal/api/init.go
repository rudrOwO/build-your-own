package api

import (
	"log"
	"os"

	"github.com/rudrowo/sqlite/internal/btree"
)

var dbFile *os.File

func Init(fileName string) *os.File {
	var err error
	dbFile, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	btree.PAGE_SIZE = int64(readPageSize())
	return dbFile
}
