package sql

import (
	"testing"

	"github.com/rudrowo/sqlite/internal/api"
)

func TestParseWhereClause(t *testing.T) {
	whereClause := "rootpage = 69"
	row1 := []any{"", "oranges", "", int64(69), ""}
	row2 := []any{"", "apples", "", int64(76), ""}

	callback, _ := parseWhereClause(whereClause, parseSchema(SQLITE_MASTER_SCHEMA))

	t.Log(callback(row1))
	t.Log(callback(row2))
}

func TestExecuteSelect(t *testing.T) {
	dbFile := api.Init("../../superheroes.db")
	defer dbFile.Close()
	// t.Log(ExecuteSelect("SELECT rootpage FROM sqlite_schema WHERE tbl_name = 'superheroes'"))
	t.Log(ExecuteSelect("SELECT name FROM superheroes WHERE hair_color = 'Reddish Brown Hair'"))
}
