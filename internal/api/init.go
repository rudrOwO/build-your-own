package api

import (
	"log"
	"os"

	btree "github.com/rudrowo/sqlite/internal/btree"
)

var (
	leafPagesChannel chan btree.LeafTablePage
	dbFile           *os.File
)

func Init(fileName string) *os.File {
	var err error
	dbFile, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	leafPagesChannel = make(chan btree.LeafTablePage, 1)
	return dbFile
}
