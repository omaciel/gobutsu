package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/omaciel/gobutsu/api"
	db "github.com/omaciel/gobutsu/db/sqlc"
	"github.com/omaciel/gobutsu/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	app := db.NewApp(conn)
	server := api.NewServer(config, app)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
