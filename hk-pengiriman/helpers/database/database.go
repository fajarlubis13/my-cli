package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

// +===========================================================================+
// | Manage Database Connection												   |
// +===========================================================================+

// DB ...
type DB struct {
	SQL *sqlx.DB
}

var (
	dbConn = &DB{}
)

// Init ...
func Init() *DB {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Duration(300 * time.Second))

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	dbConn.SQL = db

	return dbConn
}

// +===========================================================================+
