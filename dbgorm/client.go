package dbgorm

import (
	"RestApi/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DB struct {
	database *gorm.DB
}

const (
	host = "localhost"
	port = 5432
)

func (db *DB) Connect() {
	var err error
	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	dbname := os.Getenv("DBName")
	strConn := fmt.Sprintf("user=%v host=%v port=%v password=%v dbname=%v sslmode=disable", user, host, port, password, dbname)
	db.database, err = gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}

func (db *DB) Migrate() {
	db.database.AutoMigrate(&models.User{})
	log.Println("Database migration!")
}
