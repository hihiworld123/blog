package userservice

import (
	"blog/common"
	"blog/entity"
	"blog/service/domain"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Register(request domain.RegisterRequest) (domain.UserData, error) {

	user := entity.User{Username: request.Username, Email: request.Email}
	userData := domain.UserData{}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Register err: ", err)
		return userData, err
	}
	user.Password = string(hashedPassword)

	if err = common.Db.Create(&user).Error; err != nil {
		log.Error("Failed to create user: ", err)
		return userData, err
	}

	userData.User = user
	userData.Token, err = common.GenToken(user)
	if err != nil {
		log.Error("Register GenToken err: ", err)
		return userData, err
	}

	return userData, nil
}

func Login(request domain.LoginRequest) (domain.UserData, error) {

	userData := domain.UserData{}
	var storedUser entity.User
	if err := common.Db.Where("username = ?", request.Username).First(&storedUser).Error; err != nil {
		log.Error("query user err: ", err)
		return userData, err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(request.Password)); err != nil {
		log.Error("Login password err: ", err)
		return userData, err
	}

	var err error
	userData.User = storedUser
	storedUser.Password = ""
	userData.Token, err = common.GenToken(storedUser)
	if err != nil {
		log.Error("Login GenToken err: ", err)
		return userData, err
	}

	return userData, nil
}
