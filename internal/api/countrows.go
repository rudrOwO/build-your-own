package api

import (
	btree "github.com/rudrowo/sqlite/internal/btree"
)

func CountRows(tableName string) uint16 {
	go btree.LoadAllLeafTablePages(tableName, dbFile, leafPagesChannel)

	cellsCount := uint16(0)
	for page := range leafPagesChannel {
		cellsCount += page.Header.CellCount
	}

	return cellsCount
}
