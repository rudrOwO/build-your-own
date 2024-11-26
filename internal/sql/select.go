package sql

import (
	"regexp"
	"strconv"

	"github.com/rudrowo/sqlite/internal/api"
	"github.com/rudrowo/sqlite/internal/dataformat"
)

const (
	SELECT_STATEMENT_REGEX = `(?i)^SELECT\s+(.*?)\s+FROM\s+(\w+)\s*(?:\s+WHERE\s+(.*))?$`
	WHERE_CLAUSE_REGEX     = `([a-zA-Z_][a-zA-Z0-9_]*)\s*(=|!=|<=|>=|<|>)\s*'?(\w+)'?\s*`
	COUNT_REGEX            = `(?i)COUNT\(\*\)`
)

var (
	whereClauseRegex      = regexp.MustCompile(WHERE_CLAUSE_REGEX)
	selectStatementRegex  = regexp.MustCompile(SELECT_STATEMENT_REGEX)
	commaSeparatorRegex   = regexp.MustCompile(`\s*,\s*`)
	countExpresseionRegex = regexp.MustCompile(COUNT_REGEX)
)

// BUG  Fails if where clause LHS is not present in select cluase
func ExecuteSelect(query string) string {
	matches := selectStatementRegex.FindStringSubmatch(query)

	columnTokens := commaSeparatorRegex.Split(matches[1], -1)
	tableName := matches[2]
	whereClause := matches[3]

	if countExpresseionRegex.MatchString(matches[1]) {
		rowCount := api.CountRows(tableName)
		return strconv.FormatInt(int64(rowCount), 10)
	}

	schemaSql := getTableSchema(tableName)
	parsedSchema := parseSchema(schemaSql)
	columnIndices := make([]int, len(columnTokens))

	for i, columnName := range columnTokens {
		for _, parsedColumn := range parsedSchema {
			if columnName == parsedColumn.columnName {
				columnIndices[i] = parsedColumn.columnIndex
			}
		}
	}

	filter := parseWhereClause(whereClause, parsedSchema)
	return api.ScanTable(columnIndices, len(parsedSchema), tableName, filter)
}

func parseWhereClause(whereClause string, parsedSchema []parsedColumn) func(row []any) bool {
	if whereClause == "" {
		return func(row []any) bool {
			return true
		}
	}

	tokens := whereClauseRegex.FindStringSubmatch(whereClause)
	lhsToken, operatorToken, rhsToken := tokens[1], tokens[2], tokens[3]

	var intPrimitive func(int64, int64) bool
	var floatPrimitive func(float64, float64) bool
	var stringPrimitive func(string, string) bool

	switch operatorToken {
	case "=":
		intPrimitive = equalToPrimitive[int64]
		floatPrimitive = equalToPrimitive[float64]
		stringPrimitive = equalToPrimitive[string]
	case "!=":
		intPrimitive = notEqualToPrimitive[int64]
		floatPrimitive = notEqualToPrimitive[float64]
		stringPrimitive = notEqualToPrimitive[string]
	case ">":
		intPrimitive = strictlyGreaterThanPrimitive[int64]
		floatPrimitive = strictlyGreaterThanPrimitive[float64]
		stringPrimitive = strictlyGreaterThanPrimitive[string]
	case ">=":
		intPrimitive = greaterThanOrEqualToPrimitive[int64]
		floatPrimitive = greaterThanOrEqualToPrimitive[float64]
		stringPrimitive = greaterThanOrEqualToPrimitive[string]
	case "<":
		intPrimitive = strictlyLessThanPrimitive[int64]
		floatPrimitive = strictlyLessThanPrimitive[float64]
		stringPrimitive = strictlyLessThanPrimitive[string]
	}

	var targetColumn parsedColumn

	for _, parsedColumn := range parsedSchema {
		if lhsToken == parsedColumn.columnName {
			targetColumn = parsedColumn
			break
		}
	}

	i := targetColumn.columnIndex
	switch targetColumn.columnType {
	case "integer":
		rhsArg, err := strconv.ParseInt(rhsToken, 10, 64)
		if err != nil {
			panic("Type conversion failed for deserialized int64 value")
		}

		return func(row []any) bool {
			return intPrimitive(row[i].(int64), rhsArg)
		}
	case "real":
		rhsArg, err := strconv.ParseFloat(rhsToken, 64)
		if err != nil {
			panic("Type conversion failed for deserialized float64 value")
		}

		return func(row []any) bool {
			return floatPrimitive(row[i].(float64), rhsArg)
		}
	case "text":
		rhsArg := rhsToken

		return func(row []any) bool {
			return stringPrimitive(row[i].(string), rhsArg)
		}
	default:
		panic("Malformed schema passed to where clause")
	}
}

// Comparison Primitives
type Constraint dataformat.DeserializedTypes

func equalToPrimitive[T Constraint](lhs T, rhs T) bool {
	return lhs == rhs
}

func notEqualToPrimitive[T Constraint](lhs T, rhs T) bool {
	return lhs != rhs
}

func strictlyGreaterThanPrimitive[T Constraint](lhs T, rhs T) bool {
	return lhs > rhs
}

func greaterThanOrEqualToPrimitive[T Constraint](lhs T, rhs T) bool {
	return lhs >= rhs
}

func strictlyLessThanPrimitive[T Constraint](lhs T, rhs T) bool {
	return lhs < rhs
}

func lessThanOrEqualToPrimitive[T Constraint](lhs T, rhs T) bool {
	return lhs <= rhs
}
