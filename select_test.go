package godb

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nadama95/godb/pkg/adapters/postgresql"
)

type TestSelectRecord struct {
	Id        int
	IPaddress string
	UpdatedAt time.Time
}

func TestSelect(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Errorf("error loading dotenv variables %s", err)
	}

	db, err := Open(postgresql.Adapter, os.Getenv("DATABASE_URL"), &Config{PrintQueries: true})

	if err != nil {
		t.Errorf("error opening database %s", err)
		return
	}

	defer db.Close()

	q := db.Select()
	q = q.From("ipam_ipaddress")
	q = q.Columns("id::int", "ipaddress::text", "updated_at")
	q = q.Where("ipaddress::text", Like, "10.0.36.%")
	q = q.Limit(1).Offset(1)

	record := TestSelectRecord{}

	rs, err := Execute(q, record)

	if err != nil {
		t.Errorf("error executing query %s", err)
		return
	}

	fmt.Println(rs)

}
