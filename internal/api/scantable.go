package api

import (
	btree "github.com/rudrowo/sqlite/internal/btree"
)

// TODO Use a StringBuilder https://pkg.go.dev/strings#Builder
func ScanTable(columnIndices []int, tableName string, filter func(row []any) bool) string {
	go btree.LoadAllLeafTablePages(tableName, dbFile, leafPagesChannel)

	return ""
}
