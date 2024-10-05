package main

import (
	"log"
	"mybank/db"
	"mybank/routes"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db.InitDB("postgres://user:password@localhost:5432/mybank?sslmode=disable")
	db := db.DB

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully!")

	r := routes.SetupRoutes()
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
