package database

import (
	"database/sql"
	"github.com/carrot/burrow/environment"
	_ "github.com/lib/pq"
	"log"
)

var database *sql.DB

func Open() {
	// Pulling environment vars
	databaseUrl := environment.GetEnvVar(environment.PSQL_DATABASE_URL)

	// Opening + storing the connection
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Pinging the database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	database = db
}

func Close() {
	database.Close()
}

func Get() *sql.DB {
	return database
}
