package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "lbysfyyk"
	password = "zXIRi3RrROF4OeUW9Sfd"
	dbname   = "tacticalrmm"
)

type DBObject struct {
	DB  *sql.DB
	log *os.File
}

func (obj *DBObject) InitialConnectDB() {
	strConn := fmt.Sprintf("user=%v host=%v port=%v password=%v dbname=%v sslmode=disable", user, host, port, password, dbname)
	conn, err := sql.Open("postgres", strConn)
	if err != nil {
		log.Println("Database connection error: ", err)
		panic(err)
	}
	obj.DB = conn
}

/*
func ConnectDB() (*sql.DB, error) {
	strConn := fmt.Sprintf("user=%v host=%v port=%v password=%v dbname=%v sslmode=disable", user, host, port, password, dbname)
	conn, err := sql.Open("postgres", strConn)
	if err != nil {
		log.Println("Database connection error: ", err)
		return nil, err
	}

	return conn, nil
}
*/

func (db *DBObject) CloseConnection() {
	db.DB.Close()
}

func (db *DBObject) InitialLogFile() {
	pathTemp, err := os.UserCacheDir()
	fullPathLog := filepath.Join(pathTemp, DirConfig, nameLogFile)
	if err != nil {
		log.Println("Error read path UserCacheDir")
	}
	if err := os.Chdir(filepath.Join(pathTemp, DirConfig)); err != nil {
		os.Mkdir(filepath.Join(pathTemp, DirConfig), 0755)
		os.Create(fullPathLog)
	}
	(*db).log, err = os.OpenFile(fullPathLog, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("Error open file Log")
	}
}
