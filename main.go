package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("DB Project")
	var (
		err error
	)

	connDB := new(DBObject)

	connDB.DB, err = ConnectDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/aircraft", connDB.GetAircraft)
	r.GET("/routes", connDB.GetRoutes)

	r.Run("localhost:8080")
}
