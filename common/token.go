package common

import (
	"blog/entity"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

func ParseToken(tokenString string) (*entity.User, error) {
	// 解析Token
	token1, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查Token的签名方法是否是我们所期望的算法，这里我们期望的是HS256算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于验证Token的签名
		getString := viper.GetString("jwt.password")
		return []byte(getString), nil
	})
	if err != nil || !token1.Valid {
		log.Error("auth token is invalid")
	}

	user := entity.User{}
	claims := token1.Claims
	if claims != nil && claims.Valid() == nil {
		mapClaims := claims.(jwt.MapClaims)
		if mapClaims != nil {
			userId := mapClaims["id"].(float64)
			user.Id = int64(userId)
			user.Username = mapClaims["username"].(string)
		}
	}

	return &user, nil
}
