package config

import (
	"database/sql"
	"fmt"
	"log"

	"atomisu.com/ocg-statics/infoInsert/env"

	"github.com/samber/do"
)

func TestDbConnection(i *do.Injector) (*DbConn, error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		env.GetAsString("PG_DB_USER", "postgres"),
		env.GetAsString("PG_DB_PASSWORD", "postgres"),
		env.GetAsString("PG_DB_HOST_PORT", "localhost:5555"),
		env.GetAsString("PG_DB_DATABASE_TEST", "test_ocg_statics"),
	)

	// Open the database
	database, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	// Connectivity check
	if err := database.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
		return nil, err
	}

	// Set schema
	_, err = database.Exec("SET search_path TO public;")
	if err != nil {
		return nil, err
	}

	return &DbConn{database}, nil
}
