package godb

import (
	"database/sql"
	"fmt"
	"strings"
)

type selectStatement struct {
	db *DB

	columns    []string
	fromTables []string
	limit      int
	offset     int
}

func (ss *selectStatement) DB() *sql.DB {
	return ss.db.sql
}

// Create Statement Struct
func (db *DB) Select() *selectStatement {
	return &selectStatement{
		db: db,
	}
}

// Builder Functions
func (ss *selectStatement) From(tableNames ...string) *selectStatement {
	ss.fromTables = append(ss.fromTables, tableNames...)
	return ss
}

func (ss *selectStatement) Columns(columns ...string) *selectStatement {
	ss.columns = append(ss.columns, columns...)
	return ss
}

func (ss *selectStatement) Limit(limit int) *selectStatement {
	ss.limit = limit
	return ss
}

func (ss *selectStatement) Offset(offset int) *selectStatement {
	ss.offset = offset
	return ss
}

// Execution Functions
func (ss *selectStatement) buildQuery() string {
	colStr := strings.Join(ss.columns, ", ")
	q := fmt.Sprintf("SELECT %s\n", colStr)

	tableStr := strings.Join(ss.fromTables, ", ")
	q += fmt.Sprintf("FROM %s\n", tableStr)

	if ss.limit != 0 {
		q += fmt.Sprintf("LIMIT %v\n", ss.limit)
	}

	if ss.offset != 0 {
		q += fmt.Sprintf("OFFSET %v\n", ss.offset)
	}

	return q
}
