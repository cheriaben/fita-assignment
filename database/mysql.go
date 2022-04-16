package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)

func ConnectMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/assignment")

	if err != nil {
		log.Panic(err)
	}

	DBConn = db

	return db
}
