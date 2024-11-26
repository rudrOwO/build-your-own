package api

import (
	"log"
	"os"
)

var dbFile *os.File

func Init(fileName string) *os.File {
	var err error
	dbFile, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return dbFile
}
