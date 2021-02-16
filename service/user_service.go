package service

import (
	"github.com/NiharPattanaik/GoReST/dao"
	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
)

var (
	UserServiceInterface IUserServiceInterface
)

func init() {
	UserServiceInterface = &UserService{}
}

type IUserServiceInterface interface {
	GetUser(int64) (*model.User, *error.RestError)
	CreateUser(*model.User) (*model.User, *error.RestError)
	GetUsersList() (*[]model.User, *error.RestError)
	UpdateUser(int64, model.User) (*model.User, *error.RestError)
}

type UserService struct {
}

func (u *UserService) GetUser(userID int64) (*model.User, *error.RestError) {
	return dao.PGUserDAOInterface.GetUser(userID)
}

func (u *UserService) CreateUser(user *model.User) (*model.User, *error.RestError) {
	return dao.PGUserDAOInterface.CreateUser(user)
}

func (u *UserService) GetUsersList() (*[]model.User, *error.RestError) {
	return dao.PGUserDAOInterface.GetUsersList()
}

func (u *UserService) UpdateUser(userId int64, user model.User) (*model.User, *error.RestError) {
	return dao.PGUserDAOInterface.UpdateUser(userId, user)
}
