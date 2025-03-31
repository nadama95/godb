package godb

import (
	"database/sql"

	"github.com/nadama95/godb/pkg/adapters"
)

type DB struct {
	adapter adapters.Adapter
	sql     *sql.DB
}

func initalize(adapter adapters.Adapter, sql *sql.DB) *DB {
	return &DB{
		adapter: adapter,
		sql:     sql,
	}
}

func Open(adapter adapters.Adapter, dataSourceName string) (*DB, error) {
	db, err := sql.Open(adapter.DriverName(), dataSourceName)

	if err != nil {
		return nil, err
	}

	return initalize(adapter, db), nil
}

func (db *DB) Ping() bool {
	err := db.sql.Ping()

	return err == nil

}

func (db *DB) Close() error {
	return db.sql.Close()
}
