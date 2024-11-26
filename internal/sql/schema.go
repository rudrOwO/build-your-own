package sql

import (
	"regexp"
	"strings"
)

const (
	SCHEMA_REGEX = `.+?\(([^\)]+)`
	COLUMN_REGEX = `(?i)(\w+)\s+(NULL|INTEGER|REAL|TEXT|BLOB)`
)

var (
	schemaRegex = regexp.MustCompile(SCHEMA_REGEX)
	columnRegex = regexp.MustCompile(COLUMN_REGEX)
)

const SQLITE_MASTER_SCHEMA = `
    CREATE TABLE sqlite_schema(
    type text,
    name text,
    tbl_name text,
    rootpage integer,
    sql text
  );`

type parsedColumn struct {
	columnName  string
	columnType  string
	columnIndex int
}

func parseSchema(schemaSql string) []parsedColumn {
	matches := schemaRegex.FindStringSubmatch(schemaSql)
	columns := commaSeparatorRegex.Split(matches[1], -1)
	parsedSchema := make([]parsedColumn, len(columns))

	for i, column := range columns {
		matches := columnRegex.FindStringSubmatch(column)
		parsedSchema[i] = parsedColumn{
			columnName:  matches[1],
			columnType:  strings.ToLower(matches[2]),
			columnIndex: i,
		}
	}

	return parsedSchema
}

func getTableSchema(tableName string) string {
	// TODO ExecuteSelect("SELECT sql FROM sqlite_schema WHERE name = " + tableName)

	if tableName == "sqlite_schema" || tableName == "sqlite_master" {
		return SQLITE_MASTER_SCHEMA
	}

	return ""
}
