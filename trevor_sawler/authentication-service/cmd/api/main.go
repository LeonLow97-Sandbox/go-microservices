package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	Repo data.Repository
	Client *http.Client
}

func main() {
	log.Println("Starting authentication service")

	// connecting to database
	db := openDB()
	if db == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// set up config
	app := Config{
		Client: &http.Client{},
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB() *sql.DB {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("failed to open database connection")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to verify database connection")
	}

	return db
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}