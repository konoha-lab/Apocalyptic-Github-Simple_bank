package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
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
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func getRandomAccountID(t *testing.T) int64 {
	accountId1, err := testQueries.GetRandomId(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, accountId1)

	return accountId1
}
