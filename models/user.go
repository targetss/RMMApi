package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Username string `form:"username", json:"username", gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `form:"password", json:"password"`
}

func (user *User) HashPassword(pswd string) error {
	hashPswd, err := bcrypt.GenerateFromPassword([]byte(pswd), 12)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = string(hashPswd)
	return nil
}

func (user *User) CheckPassword(pswd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pswd)); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
