package api

import (
	"sort"
	"strconv"
	"strings"

	btree "github.com/rudrowo/sqlite/internal/btree"
	"github.com/rudrowo/sqlite/internal/dataformat"
)

func ScanTable(columnIndices []int, rowLength int, tableName string, filter func(row []any) bool) string {
	go btree.LoadAllLeafTablePages(tableName, dbFile, leafPagesChannel)

	if !sort.IntsAreSorted(columnIndices) {
		sort.Ints(columnIndices)
	}

	var result strings.Builder

	for page := range leafPagesChannel {
		for _, cell := range page.Cells { // each cell corresponds to a row
			row := make([]any, rowLength)
			j, k := 0, 0

			for i, columnType := range cell.Payload.ColumnTypes {
				recordBody := cell.Payload.RecordBody
				contentSize := int(dataformat.GetContentSize(columnType))

				// Lazy serializer: serialze only the selected columns
				if i == columnIndices[j] {
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
					j += 1

					if filter(row) {
						switch content := row[i].(type) {
						case int64:
							result.WriteString(strconv.FormatInt(content, 10))
						case float64:
							result.WriteString(strconv.FormatFloat(content, 'f', 2, 64))
						case string: // string
							result.WriteString(content)
						}
						result.WriteString("\n")
					}

				} else {
					row[i] = nil
				}

				k += contentSize

				if len(columnIndices) == j {
					break
				}

			}
		}
	}

	return result.String()
}
