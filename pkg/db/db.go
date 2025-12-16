package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func Connect(dns string) *sql.DB {
	db. err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatalf("Failed to prepare database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	fmt.Println("✅ Database connected successfully")
	return db
}