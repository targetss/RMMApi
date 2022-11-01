package main

import (
	"RestApi/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DBObject) Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		cookie, err := c.Cookie("JWTAuth")
		if err != nil || auth.ValidateToken(cookie) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":   http.StatusUnauthorized,
				"response": "NoAuth",
			})
			c.Abort()
			return
		}
		/*
			if err := auth.ValidateToken(cookie); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":   http.StatusUnauthorized,
					"response": "Token no Valid",
				})
				c.Abort()
				return
			}
			c.Next()
		*/
	}

	//Проверка по Хедеру, проверяет json с учетками авторизации
	/*
		return func(c *gin.Context) {
			strSearch := "select name, username, password from users where username = $1"
			var userReq models.User
			strAuth := c.GetHeader("Authorization")
			if strAuth != "Authorization" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":   http.StatusInternalServerError,
					"response": "Error header parameter authorization",
				})
				c.Abort()
			}
			err := c.ShouldBindJSON(&userReq)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":   http.StatusInternalServerError,
					"response": "Error parameter json users",
				})
				c.Abort()
			}
			var userDB models.User
			resultSearch := db.DB.QueryRow(strSearch, userReq.Username)
			err = resultSearch.Scan(&userDB.Name, &userDB.Username, &userDB.Password)
			switch err {
			case sql.ErrNoRows:
				c.JSON(http.StatusNoContent, gin.H{
					"status":   http.StatusNoContent,
					"response": "User not found or user not register",
				})
				c.Abort()
				return
			case nil:
				if err := userDB.CheckPassword(userReq.Password); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":   http.StatusInternalServerError,
						"response": "Password incorrect, check the entered data",
					})
					c.Abort()
					return
				}
				c.Next()
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":   http.StatusInternalServerError,
					"response": "User not found, check request data",
				})
				c.Abort()
				return
			}
		}
	*/
}
