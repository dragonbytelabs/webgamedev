package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dragonbytelabs/webgamedev/internal/dbx"
	"github.com/dragonbytelabs/webgamedev/internal/routes"
)

const port = ":3000"

func main() {
	db := setupDB()
	mux := http.NewServeMux()
	setupRoutes(mux, db)

	log.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func setupDB() *dbx.DB {
	ctx := context.Background()
	db, err := dbx.OpenSQLite("basicrouter.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.ApplyMigrations(ctx); err != nil {
		log.Fatal(err)
	}
	return db
}

func setupRoutes(mux *http.ServeMux, db *dbx.DB) {
	routes.RegisterStatic(mux)
	routes.RegisterAPI(mux)
	routes.RegisterAuth(mux, db)
}
