package dao

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/NiharPattanaik/GoReST/datasource/postgress"
	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
)

const (
	updateUserQry   = "UPDATE USERS SET FIRST_NAME = $1, LAST_NAME = $2, EMAIL = $3, DATE_UPDATED = NOW() WHERE ID = $4 RETURNING ID,FIRST_NAME, LAST_NAME, EMAIL, to_char(DATE_CREATED, 'DD-MM-YYYY'), to_char(DATE_UPDATED, 'DD-MM-YYYY');"
	createUSerQry   = "INSERT INTO USERS (FIRST_NAME, LAST_NAME, EMAIL, PASSWORD, DATE_CREATED, DATE_UPDATED) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING ID,FIRST_NAME, LAST_NAME, EMAIL, PASSWORD, to_char(DATE_CREATED, 'DD-MM-YYYY'), to_char(DATE_UPDATED, 'DD-MM-YYYY');"
	getUserQry      = "SELECT ID,FIRST_NAME, LAST_NAME, EMAIL, to_char(DATE_CREATED, 'DD-MM-YYYY'), to_char(DATE_UPDATED, 'DD-MM-YYYY') FROM USERS WHERE ID = $1"
	getUsersListQry = "SELECT ID,FIRST_NAME, LAST_NAME, EMAIL, to_char(DATE_CREATED, 'DD-MM-YYYY'), to_char(DATE_UPDATED, 'DD-MM-YYYY') FROM USERS"
	noResultErrMsg  = "no rows in result set"
)

type UserDAOPG struct{}

func (u *UserDAOPG) CreateUser(user *model.User) (*model.User, *error.RestError) {
	var newUser model.User
	err := postgress.DBPool.QueryRow(context.Background(), createUSerQry, user.FirstName, user.LastName, user.Email, user.Password).Scan(&newUser.Id, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.Password, &newUser.DateCreated, &newUser.DateUpdated)
	if err != nil {
		return nil, &error.RestError{
			Message:    "Error while saving user in database",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}

	}
	(&newUser).Password = ""
	return &newUser, nil
}

func (u *UserDAOPG) GetUser(userId int64) (*model.User, *error.RestError) {
	var user model.User
	err := postgress.DBPool.QueryRow(context.Background(), getUserQry, userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.DateUpdated)
	if err != nil {
		if strings.Contains(err.Error(), noResultErrMsg) {
			return nil, &error.RestError{
				Message:    fmt.Sprintf("No user found having user Id : %d", userId),
				StatusCode: http.StatusNotFound,
				Error:      err.Error(),
			}
		} else {
			return nil, &error.RestError{
				Message:    fmt.Sprintf("Could not retrieve the user details for user having Id : %d", userId),
				StatusCode: http.StatusInternalServerError,
				Error:      err.Error(),
			}
		}
	}
	return &user, nil
}

func (u *UserDAOPG) GetUsersList() (*[]model.User, *error.RestError) {
	users := make([]model.User, 0)
	rows, err := postgress.DBPool.Query(context.Background(), getUsersListQry)
	if err != nil {
		return nil, &error.RestError{
			Message:    "Error while fetching the users",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.DateUpdated)
		if err != nil {
			return nil, &error.RestError{
				Message:    "Error while fetching the users",
				StatusCode: http.StatusInternalServerError,
				Error:      err.Error(),
			}
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, &error.RestError{
			Message:    "No users found",
			StatusCode: http.StatusNotFound,
			Error:      "Users not found",
		}
	}
	return &users, nil
}

func (u *UserDAOPG) UpdateUser(userId int64, user model.User) (*model.User, *error.RestError) {
	var newUser model.User
	err := postgress.DBPool.QueryRow(context.Background(), updateUserQry, user.FirstName, user.LastName, user.Email, userId).Scan(&newUser.Id, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.DateCreated, &newUser.DateUpdated)
	if err != nil {
		return nil, &error.RestError{
			Message:    "Error while updating user in database",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}

	}
	return &newUser, nil
}
