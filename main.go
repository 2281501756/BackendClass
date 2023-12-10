package main

import (
	"database/sql"
	"github.com/2281501756/BackendClass/api"
	db "github.com/2281501756/BackendClass/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:donghao@localhost:5432/bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Print("db can't connect", err.Error())
		return
	}
	server := api.NewServer(db.NewStore(conn))
	_ = server.Router.Run(":8080")
}
