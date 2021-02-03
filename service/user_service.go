package service

import (
	"github.com/NiharPattanaik/GoReST/dao/postgress"
	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
)

func GetUser(userID int64) (*model.User, *error.RestError) {
	return postgress.GetUser(userID)
}

func CreateUser(user *model.User) (*model.User, *error.RestError) {
	return postgress.CreateUser(user)
}

func GetUsersList() (*[]model.User, *error.RestError) {
	return postgress.GetUsersList()
}
