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
	r.GET("/list_site", connDB.GetListSite)
	r.GET("/pc_to_site/:id", connDB.GetPCToSite)
	r.GET("info_pc/:id", connDB.GetInfoComputer)
	//r.GET("/routes", connDB.GetRoutes)
	//r.POST("/aircrafts", connDB.PostAircrafts)

	r.Run("localhost:8080")
}
