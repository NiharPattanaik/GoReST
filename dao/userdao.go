package dao

import (
	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
)

var (
	PGUserDAOInterface IUserDAO
)

type IUserDAO interface {
	CreateUser(*model.User) (*model.User, *error.RestError)
	GetUser(int64) (*model.User, *error.RestError)
	GetUsersList() (*[]model.User, *error.RestError)
	UpdateUser(int64, model.User) (*model.User, *error.RestError)
}

func init() {
	PGUserDAOInterface = &UserDAOPG{}
}
