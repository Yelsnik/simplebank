package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Yelsnik/simplebank/util"
	_ "github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

// var testDB *pgxpool.Pool
var testStore Store

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("could not connect")
	}

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	os.Exit(m.Run())
}
