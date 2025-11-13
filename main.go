package main

import (
	"context"
	"fmt"
	"log"
	api "menribardhi/micro-go-psql/api"
	db "menribardhi/micro-go-psql/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var err error
	dbUrl := "postgres://postgres:passwrod123@localhost:5432/simplebank?sslmode=disable"
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to ping db err:%v", err)
	}
	store := db.NewStore(pool)
	server := api.NewServer(store)
	adress := "0.0.0.0:8084"
	err = server.Start(adress)
	if err != nil {
		log.Fatalf("could not start server at port ")
	}
	fmt.Println("Server started at adress : %s", adress)
}
