package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testPool *pgxpool.Pool
var testStore *Store

func TestMain(m *testing.M) {
	var err error
	dbUrl := "postgres://postgres:passwrod123@localhost:5432/simplebank?sslmode=disable"
	ctx := context.Background()
	testPool, err = pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to ping db err:%v", err)
	}
	defer testPool.Close()
	testStore = NewStore(testPool)
	if testStore.db == nil {
		log.Panic("something went wrong")
	}
	testQueries = New(testPool)
	if testQueries == nil {
		log.Println("something went wrong")
	}
	log.Println("Database connected for tests")
	code := m.Run()
	testPool.Close()

	os.Exit(code)

}
