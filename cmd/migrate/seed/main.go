package main

import (
	"log"

	"github.com/BMPaiba/Go-Backend-Engineering/internal/db"
	"github.com/BMPaiba/Go-Backend-Engineering/internal/env"
	"github.com/BMPaiba/Go-Backend-Engineering/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 30, 30, "15m")

	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStorage(conn)

	db.Seed(store)
	defer conn.Close()
}
