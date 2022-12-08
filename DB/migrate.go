package main

import (
	"database/sql"
	"gateway_api/Helper"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		log.Fatal("Nothing happened")
		return
	}
	Helper.LoadDB()

	dataSourceName := Helper.DBConnection + "://" + Helper.DBUsername + ":" + Helper.DBPassword + "@" + Helper.DBHost + ":" + Helper.DBPort + "/" + Helper.DBName + "?sslmode=" + Helper.DBSSLMode
	db, err := sql.Open(Helper.DBConnection, dataSourceName)
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://DB/migrations",
		Helper.DBConnection, driver)
	if err != nil {
		log.Fatal(err)
	}

	arg1 := os.Args[1]
	if arg1 == "up" || arg1 == "UP" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Migrated")
	} else if arg1 == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Dropped")
	}
	defer m.Close()
}
