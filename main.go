package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/omaciel/gobutsu/api"
	db "github.com/omaciel/gobutsu/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "user=root password=secret host=localhost dbname=gobutsu sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	app := db.NewApp(conn)
	server := api.NewServer(app)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}