package godb

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type selectStatement struct {
	db *DB

	columns    []string
	fromTables []string
	joins      []Join
	orderBy    []string
	limit      int
	offset     int
}

type Join struct {
	JoinType
	tableName string
	localOn   string
	remoteOn  string
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

type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	OuterJoin
)

func (j JoinType) String() string {
	switch j {
	case InnerJoin:
		return "INNER JOIN"
	case LeftJoin:
		return "OUTER JOIN"
	case RightJoin:
		return "RIGHT JOIN"
	case OuterJoin:
		return "FULL OUTER JOIN"
	default:
		return ""
	}
}

func (ss *selectStatement) Join(jt JoinType, table string, localOn string, remoteOn string) *selectStatement {
	ss.joins = append(ss.joins, Join{
		JoinType:  jt,
		tableName: table,
		localOn:   localOn,
		remoteOn:  remoteOn,
	})
	return ss
}

func (ss *selectStatement) OrderBy(columns ...string) *selectStatement {
	ss.orderBy = append(ss.orderBy, columns...)
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
	colStr := strings.Join(ss.columns, ",\n\t")
	q := fmt.Sprintf("SELECT %s\n", colStr)

	tableStr := strings.Join(ss.fromTables, ", ")
	q += fmt.Sprintf("FROM %s\n", tableStr)

	for _, j := range ss.joins {
		q += fmt.Sprintf("%s %s ON %s.%s = %s.%s\n", j.JoinType.String(), j.tableName, j.tableName, j.localOn, ss.fromTables[0], j.remoteOn)
	}

	orderStr := strings.Join(ss.orderBy, ",\n")
	q += fmt.Sprintf("ORDER BY %s\n", orderStr)

	if ss.limit != 0 {
		q += fmt.Sprintf("LIMIT %v\n", ss.limit)
	}

	if ss.offset != 0 {
		q += fmt.Sprintf("OFFSET %v\n", ss.offset)
	}

	log.Println(q)

	return q
}
