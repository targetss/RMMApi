package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JWTKey = []byte("TacticalRMM")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, username string) (string, error) {
	lifetimeJWT := time.Now().Add(1 * time.Hour)
	userJWT := &JWTClaim{
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(lifetimeJWT),
			Issuer:    "TacticalRMM",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJWT)
	ss, err := token.SignedString(JWTKey)
	return ss, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
			err = errors.New("Token expired")
			return
		}
	} else {
		err = errors.New("Error parse token")
		return
	}
	return
}
