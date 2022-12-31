package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(context.Background(), config)
	defer conn.Close(context.Background())
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
