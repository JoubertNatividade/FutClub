package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func ConnectionMySQL() *sql.DB {
	log.Infoln("Trying to connect to database...")

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/futclub")
	if err != nil {
		log.Fatalf("Error on connecting to database: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error on verifying connection: %s", err)
	}

	log.Infoln("Connection Success!")
	return db
}
