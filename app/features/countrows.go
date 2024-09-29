package features

import (
	"os"

	btree "github/com/codecrafters-io/sqlite-starter-go/app/btree"
)

func CountRows(tableName string, dbFile *os.File) uint16 {
	leafPagesChannel := make(chan btree.LeafTablePage, 1)
	go btree.LoadAllLeafTablePages(tableName, dbFile, leafPagesChannel)

	cellsCount := uint16(0)
	for c := range leafPagesChannel {
		cellsCount += c.Header.CellsCount
	}

	return cellsCount
}
