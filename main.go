package main

import (
	"github.com/abeer93/simple-bank.git/util"
	"github.com/abeer93/simple-bank.git/api"
	"github.com/abeer93/simple-bank.git/db/sqlc"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load cofnig: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}
