package btree

import (
	"os"

	u "github/com/codecrafters-io/sqlite-starter-go/app/utils"
)

/*
Traverse a table btree and enumerate list of
starting poitions of leaf pages
*/
func LoadAllLeafTablePages(tableName string, dbFile *os.File, leafTablesChannel chan<- LeafTablePage) {
	loadAllLeafTablePages(getRootPageOffset(tableName), dbFile, leafTablesChannel)
	close(leafTablesChannel)
}

func getRootPageOffset(tableName string) int64 {
	if tableName == "sqlite_schema" || tableName == "sqlite_master" {
		return SQLITE_SCHEMA_ROOT_OFFSET
	} else {
		// TODO parse sqlite_schema table and find root offset
		return 0
	}
}

func loadAllLeafTablePages(rootPageOffset int64, dbFile *os.File, leafTablesChannel chan<- LeafTablePage) {
	fileBuffer := make([]byte, PAGE_SIZE)
	_, err := dbFile.Seek(rootPageOffset, 0)
	u.HandleError(err)
	_, err = dbFile.Read(fileBuffer)
	u.HandleError(err)

	if fileBuffer[0] == LEAF_TABLE_PAGE_TYPE {
		leafPage := new(LeafTablePage)
		leafPage.loadPageFromBuffer(fileBuffer)
		leafTablesChannel <- *leafPage
	} else {
		interiorPage := new(interiorTablePage)
		interiorPage.loadPageFromBuffer(fileBuffer)

		loadAllLeafTablePages(int64(interiorPage.header.rightmostPointer-1)*PAGE_SIZE, // * page offsets stored in db are 1 based
			dbFile, leafTablesChannel)

		for _, c := range interiorPage.cells {
			loadAllLeafTablePages(int64(c.leftChildPointer-1)*PAGE_SIZE, // * page offsets stored in db are 1 based
				dbFile, leafTablesChannel)
		}
	}
}
