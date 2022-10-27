package main

import (
	"RestApi/dbgorm"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"net/http"
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

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NoSet"
			c.SetCookie("auth_jwt", "jwt_id1", 0, "/", "localhost", false, true)
		}
		fmt.Printf("cookie: %s \n", cookie)
	})
	r.GET("/test_cookie", func(c *gin.Context) {
		cookiejwt, err := c.Cookie("auth_jwt")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":   http.StatusForbidden,
				"response": "error cookie",
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"jwt":    cookiejwt,
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/register", connDB.RegisterUser)
		auth.POST("/token", connDB.GenerateToken)
		api := auth.Group("/api").Use(connDB.Auth())
		{
			api.GET("/users", connDB.GetAccountsUser)
			api.GET("/site", connDB.GetListSite)
			api.GET("/pc-site/:id", connDB.GetPCToSite)
			api.GET("/pc-info/:id", connDB.GetInfoComputer)
		}
	}

	r.Run("localhost:8080")
}
