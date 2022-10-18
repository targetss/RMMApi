package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("DB Project")

	connDB := new(DBObject)
	connDB.InitialConnectDB()
	connDB.InitialLogFile()

	defer connDB.CloseConnection()

	r := gin.Default()
	r.GET("/account_user", connDB.GetAccountsUser)
	r.GET("/list_site:id", connDB.GetListSite)
	//r.GET("/routes", connDB.GetRoutes)
	//r.POST("/aircrafts", connDB.PostAircrafts)

	r.Run("localhost:8080")
}
