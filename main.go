package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"os"
)

func main() {
	fmt.Println("DB Project")

	connDB := new(DBObject)
	connDB.InitialConnectDB()
	connDB.InitialLogFile()

	defer connDB.CloseConnection()

	gin.DefaultWriter = io.MultiWriter(connDB.log, os.Stdout)
	r := gin.Default()

	GInfo := r.Group("/info")
	{
		GInfo.GET("/users", connDB.GetAccountsUser)
		GInfo.GET("/site", connDB.GetListSite)
		GInfo.GET("/pc-site/:id", connDB.GetPCToSite)
		GInfo.GET("pc-info/:id", connDB.GetInfoComputer)
	}

	r.Run("localhost:8080")
}
