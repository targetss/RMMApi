package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	host = "localhost"
	port = 5432
)

type DBObject struct {
	DB  *sql.DB
	log *os.File
}

func (obj *DBObject) InitialConnectDB() {
	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	dbname := os.Getenv("DBName")
	strConn := fmt.Sprintf("user=%v host=%v port=%v password=%v dbname=%v sslmode=disable", user, host, port, password, dbname)
	conn, err := sql.Open("postgres", strConn)
	if err != nil {
		log.Println("Database connection error: ", err)
		panic(err)
	}
	obj.DB = conn
}

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

func (db *DBObject) WriteLog(strlog string) (int, error) {
	//strlog = append(strlog, byte('\n'))
	count, err := db.log.WriteString(strlog)
	db.log.WriteString("\n")
	return count, err
}

func (db *DBObject) CheckUserInTable(paramQuery ...string) {
	if countParam := len(paramQuery); countParam < 1 {
		return
	}
	str := make([]string, 0)
	for i, v := range paramQuery {
		if i == 1 {
			return
		}
		str = append(str, v)
	}
	db.DB.QueryRow(paramQuery[0])
}
