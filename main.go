package main

import (
	"context"
	"fmt"
	"log"
	api "menribardhi/micro-go-psql/api"
	"menribardhi/micro-go-psql/config"
	db "menribardhi/micro-go-psql/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var err error
	con, err := config.ReadConfig("./")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(con)

	dbUrl := con.DbUrl
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to ping db err:%v", err)
	}
	store := db.NewStore(pool)
	server := api.NewServer(store)
	adress := con.AppAdress
	err = server.Start(adress)
	if err != nil {
		log.Fatalf("could not start server at port ")
	}
	fmt.Println("Server started at adress : %s", adress)
}
