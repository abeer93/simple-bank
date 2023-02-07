package db

import (
	"github.com/abeer93/simple-bank.git/util"
	"os"
	"log"
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not load cofnig: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to DB", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
