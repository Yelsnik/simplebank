package db

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

// var testDB *pgxpool.Pool
var testStore *Store

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	testDB, err := pgxpool.New(context.Background(), dbsource)

	if err != nil {
		log.Fatal("could not connect")
	}

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	os.Exit(m.Run())
}
