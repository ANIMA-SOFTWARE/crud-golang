package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"main/api"
	"main/storage"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const fileName = "sqlite.db"

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()
	os.Remove(fileName)

	db, err := sql.Open("sqlite3", fileName)
	var ctx context.Context

	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewSQLiteStorage(db, ctx)
	store.Migrate()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())

}
