package postgresql

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgreSQL struct{}

var Adapter = PostgreSQL{}

func (PostgreSQL) DriverName() string {
	return "pgx"
}
