package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Marib777/lifecrm/internal/httpapi"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://lifecrm:lifecrm@localhost:5432/lifecrm?sslmode=disable"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	mux := httpapi.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("LifeCRM API listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
