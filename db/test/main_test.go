package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	_repo "simple_bank/db/repository"
	"simple_bank/util"
)

var testQueries *_repo.Queries
var testDB *sql.DB

// func TestMain(m *testing.M) {
// 	conn, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	testQueries = New(conn)

// 	os.Exit(m.Run())
// }

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = _repo.New(testDB)

	os.Exit(m.Run())
}
