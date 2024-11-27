package btree

import (
	"log"
	"os"
)

/*
Traverse a table btree and enumerate list of
starting poitions of leaf pages
*/
func LoadAllLeafTablePages(tableName string, dbFile *os.File, leafPagesChannel chan<- LeafTablePage) {
	loadAllLeafTablePages(getRootPageOffset(tableName), dbFile, leafPagesChannel)
	close(leafPagesChannel)
}

func getRootPageOffset(tableName string) int64 {
	if tableName == "sqlite_schema" || tableName == "sqlite_master" {
		return SQLITE_SCHEMA_ROOT_OFFSET
	} else {
		return 0
	}
}

func loadAllLeafTablePages(rootPageOffset int64, dbFile *os.File, leafTablesChannel chan<- LeafTablePage) {
	fileBuffer := make([]byte, PAGE_SIZE)
	_, err := dbFile.Seek(rootPageOffset, 0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = dbFile.Read(fileBuffer)
	if err != nil {
		log.Fatal(err)
	}

	var isLeafPage bool
	var isSchemaPage bool

	if rootPageOffset == 0 {
		isSchemaPage = true
		isLeafPage = fileBuffer[SQLITE3_HEADER_SIZE] == LEAF_TABLE_PAGE_TYPE
	} else {
		isSchemaPage = false
		isLeafPage = fileBuffer[0] == LEAF_TABLE_PAGE_TYPE
	}

	if isLeafPage {
		leafPage := LeafTablePage{}
		leafPage.loadFromBuffer(fileBuffer, isSchemaPage)
		leafTablesChannel <- leafPage
	} else {
		interiorPage := interiorTablePage{}
		interiorPage.loadFromBuffer(fileBuffer, isSchemaPage)

		loadAllLeafTablePages(int64(interiorPage.header.rightmostPointer-1)*PAGE_SIZE, // * page offsets stored in db are 1 based
			dbFile, leafTablesChannel)

		for _, c := range interiorPage.cells {
			loadAllLeafTablePages(int64(c.leftChildPointer-1)*PAGE_SIZE, // * page offsets stored in db are 1 based
				dbFile, leafTablesChannel)
		}
	}
}
