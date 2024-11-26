package api

import (
	"sort"
	"strconv"
	"strings"

	btree "github.com/rudrowo/sqlite/internal/btree"
	"github.com/rudrowo/sqlite/internal/dataformat"
)

func ScanTable(columnIndices []int, rowLength int, rootPageOffset int64, filter func(row []any) bool) string {
	leafPagesChannel := make(chan btree.LeafTablePage, 1)
	go btree.LoadAllLeafTablePages(rootPageOffset, dbFile, leafPagesChannel, true)

	if !sort.IntsAreSorted(columnIndices) {
		sort.Ints(columnIndices)
	}

	var result strings.Builder

	for page := range leafPagesChannel {
		for _, cell := range page.Cells { // each cell corresponds to a row
			row := make([]any, rowLength)
			j, k := 0, 0
			firstPrintInRow := true

			for i, columnType := range cell.Payload.ColumnTypes {
				recordBody := cell.Payload.RecordBody
				contentSize := int(dataformat.GetContentSize(columnType))

				// Lazy serializer: serialze only the selected columns
				var content any

				switch {
				case columnType == 0: // NULL
					content = nil
				case columnType >= 1 && columnType <= 6: // int
					content = dataformat.DeserializeInteger(recordBody[k : k+contentSize])
				case columnType == 7: // float
					content = dataformat.DeserializeFloat(recordBody[k : k+contentSize])
				default: // string
					content = string(recordBody[k : k+contentSize])
				}

				row[i] = content
				k += contentSize

				if j < len(columnIndices) && i == columnIndices[j] {
					if filter(row) {
						if !firstPrintInRow {
							result.WriteString("|")
						}

						switch content := row[i].(type) {
						case int64:
							result.WriteString(strconv.FormatInt(content, 10))
						case float64:
							result.WriteString(strconv.FormatFloat(content, 'f', 2, 64))
						case string: // string
							result.WriteString(content)
						}
						firstPrintInRow = false
					}
					j += 1
				}
			}
			firstPrintInRow = true

			if filter(row) {
				result.WriteString("\n")
			}
		}
	}

	return result.String()
}
