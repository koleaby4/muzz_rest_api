package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/koleaby4/muzz_rest_api/config"
	"log"
)

func GetQueries() *Queries {
	dsn := config.GetConfig("DSN")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalln("error connecting to the database", err)
	}
	return New(conn)
}
