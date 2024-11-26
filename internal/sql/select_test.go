package sql

import "testing"

func TestParseWhereClause(t *testing.T) {
	whereClause := "rootpage = 69"
	row1 := []any{"", "oranges", "", int64(69), ""}
	row2 := []any{"", "apples", "", int64(76), ""}

	callback := parseWhereClause(whereClause, parseSchema(SQLITE_MASTER_SCHEMA))

	t.Log(callback(row1))
	t.Log(callback(row2))
}

func TestExecuteSelect(t *testing.T) {
	t.Log(ExecuteSelect("SELECT sql FROM sqlite_schema"))
	t.Log(ExecuteSelect("SELECT name, rootpage FROM sqlite_schema WHERE tbl_name = 'apples'"))
}
