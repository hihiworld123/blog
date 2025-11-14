package domain

import "blog/entity"

type UserData struct {
	entity.User
	Token string `json:"token"`
}
