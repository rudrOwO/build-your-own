package api

import (
	btree "github.com/rudrowo/sqlite/internal/btree"
)

func CountRows(rootPageOffset int64) uint16 {
	go btree.LoadAllLeafTablePages(rootPageOffset, dbFile, leafPagesChannel)

	cellsCount := uint16(0)
	for page := range leafPagesChannel {
		cellsCount += page.Header.CellCount
	}

	return cellsCount
}
