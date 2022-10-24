package main

import (
	"RestApi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (db *DBObject) GenerateToken(c *gin.Context) {

}

func (db *DBObject) RegisterUser(c *gin.Context) {
	strAddData := "INSERT INTO user (name, username, email, password) values($1, $2, $3, $4)"

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
}
