package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "demo"
)

type DBObject struct {
	DB *sql.DB
}

func ConnectDB() (*sql.DB, error) {
	strConn := fmt.Sprintf("user=%v host=%v port=%v password=%v dbname=%v sslmode=disable", user, host, port, password, dbname)
	conn, err := sql.Open("postgres", strConn)
	if err != nil {
		log.Println("Database conniction error: ", err)
		return nil, err
	}

	return conn, nil
}

func (db *DBObject) CloseConnection() {
	db.DB.Close()
}
