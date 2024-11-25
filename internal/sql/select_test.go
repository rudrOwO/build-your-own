package sql

import "testing"

func TestSelect(t *testing.T) {
	testQuery := "SELECT sql FROM sqlite_schema"
	ExecuteSelect(testQuery)
}

func TestParseWhereClause(t *testing.T) {
	whereClause := "rootpage < 69"
	tableName := "sqlite_schema"
	row1 := []any{"", "oranges", "", int64(69), ""}
	row2 := []any{"", "apples", "", int64(76), ""}

	callback := parseWhereClause(whereClause, tableName)

	t.Log(callback(row1))
	t.Log(callback(row2))
}
