package common

import (
	"blog/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 生成 JWT
func GenToken(storedUser entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.Id,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("your_secret_key"))
}
