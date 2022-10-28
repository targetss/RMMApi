package main

import (
	"RestApi/auth"
	"RestApi/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DBObject) GenerateToken(c *gin.Context) {
	strConn := "select name, username, password from users where username = $1"
	var userReq models.User
	var userDB models.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"response": "Incorrect data auth",
		})
		db.WriteLog("Incorrect bind json (GenerateToken)")
		c.Abort()
	}
	result := db.DB.QueryRow(strConn, userReq.Username)
	err := result.Scan(&userDB.Name, &userDB.Username, &userDB.Password)
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNoContent, gin.H{
			"status":   http.StatusNoContent,
			"response": "User not found",
		})
	case nil:
		if err := userDB.CheckPassword(userReq.Password); err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":   http.StatusForbidden,
				"response": "Password incorrect",
			})
			c.Abort()
		}
		tokenString, err := auth.GenerateJWT(userReq.Email, userReq.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.SetCookie("JWTAuth", tokenString, 3600, "/", "localhost", false, true)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"response": "Error work database",
		})
	}

}

func (db *DBObject) RegisterUser(c *gin.Context) {
	strAddData := "INSERT INTO users (name, username, email, password) values($1, $2, $3, $4)"

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   http.StatusBadRequest,
			"response": "Incorrect data register",
		})
		db.WriteLog("Incorrect bind json (RegisterUser)")
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"response": "incorrect password",
		})
		db.WriteLog(fmt.Sprintf("Incorrect hashPassword, password:%v", user.Password))
		c.Abort()
	}

	//var usr models.User
	strSearchUserRMM := "select username, email from accounts_user where username = '$1' and email = '$1'"
	userResult := db.DB.QueryRow(strSearchUserRMM, user.Username, user.Email)
	err := userResult.Scan()
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNoContent, gin.H{
			"status":   http.StatusNoContent,
			"response": "Account not fount server TacticalRMM",
		})
		c.Abort()
	case nil:

		res, err := db.DB.Exec(strAddData, user.Name, user.Username, user.Email, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":   http.StatusInternalServerError,
				"response": "InternalServerError",
			})
			db.WriteLog("Ошибка запроса в базу данных!")
			c.Abort()
		}
		countAdd, _ := res.RowsAffected()
		c.JSON(http.StatusOK, gin.H{
			"username":    user.Username,
			"email":       user.Email,
			"user":        user.Name,
			"countAddStr": countAdd,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   http.StatusInternalServerError,
			"response": "Server internal Error",
		})
		c.Abort()
	}
}
