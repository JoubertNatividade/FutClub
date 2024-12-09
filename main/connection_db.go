package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func ConnectionMySQL() *sql.DB {
	log.Infoln("Trying to connect to database...")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/futtart")
	if err != nil {
		log.Errorf("Error on connect to database: %s", err)
	}
	db.Close()
	log.Infoln("Connection Success!")
	return db
}
