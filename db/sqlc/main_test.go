package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:donghao@localhost:5432/bank?sslmode=disable"
)

var TestStore Store

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	TestStore = NewStore(conn)
	if err != nil {
		log.Fatal("connect error", err.Error())
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
