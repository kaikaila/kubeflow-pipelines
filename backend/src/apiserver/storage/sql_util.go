// Package storage provides database-agnostic SQL utility functions,
// including dialect-aware pagination, identifier quoting, and builder selection.
package storage

import (
	sq "github.com/Masterminds/squirrel"
)

var (
	// While MySQL & SQLite uses ? for placeholder
	sqBuilderDefault = sq.StatementBuilder
	// PostgreSQL placeholders uses $
	sqBuilderDollar = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

// Builder returns the appropriate squirrel StatementBuilderType for the given dialect.
func Builder(dialect SQLDialect) sq.StatementBuilderType {
	switch dialect.(type) {
	case PostgreDialect:
		return sqBuilderDollar
	default: // MySQLDialect, SQLiteDialect
		return sqBuilderDefault
	}
}

// QuoteIdentifier returns a quoted identifier for each dialect.
func QuoteIdentifier(dialect SQLDialect, id string) string {
	switch dialect.(type) {
	case PostgreDialect:
		// double quotes, preserves case for Postgres
		return `"` + id + `"`
	case MySQLDialect:
		// backticks for MySQL
		return "`" + id + "`"
	case SQLiteDialect:
		// double quotes for SQLite
		return `"` + id + `"`
	default:
		return id
	}
}

func QuoteColumns(dialect SQLDialect, cols []string) []string {
	q := make([]string, len(cols))
	for i, c := range cols {
		q[i] = QuoteIdentifier(dialect, c)
	}
	return q
}
