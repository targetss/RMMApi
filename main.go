package main

import (
	"RestApi/dbgorm"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"os"
)

func main() {
	fmt.Println("DB Project")
	dbGorm := new(dbgorm.DB)
	dbGorm.Connect()
	dbGorm.Migrate()

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
		GInfo.GET("/pc-info/:id", connDB.GetInfoComputer)
	}
	auth := r.Group("/auth")
	{
		auth.POST("/register", connDB.RegisterUser)
		auth.POST("/token", connDB.GenerateToken)
		api := auth.Group("/api").Use(Auth())
		{
			api.GET("/users", connDB.GetAccountsUser)
			api.GET("/site", connDB.GetListSite)
			api.GET("/pc-site/:id", connDB.GetPCToSite)
			api.GET("/pc-info/:id", connDB.GetInfoComputer)
		}
	}

	r.Run("localhost:8080")
}
