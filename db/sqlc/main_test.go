package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testpool *pgxpool.Pool

func TestMain(m *testing.M) {

	dbUrl := "postgres://postgres:passwrod123@localhost:5432/simplebank?sslmode=disable"
	ctx := context.Background()
	testpool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to ping db err:%v", err)
	}
	defer testpool.Close()

	testQueries = New(testpool)
	if testQueries == nil {
		log.Println("something went wrong")
	}
	log.Println("Database connected for tests")
	code := m.Run()
	testpool.Close()

	os.Exit(code)

}
