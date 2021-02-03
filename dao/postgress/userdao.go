package postgress

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NiharPattanaik/GoReST/datasource/postgress"
	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
)

const (
	createUSerQry   = "INSERT INTO USERS (FIRST_NAME, LAST_NAME, EMAIL, DATE_CREATED, DATE_UPDATED) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING ID,FIRST_NAME, LAST_NAME, EMAIL, to_char(DATE_CREATED, 'DD/MM/YYYY'), to_char(DATE_UPDATED, 'DD/MM/YYYY');"
	getUserQry      = "SELECT ID,FIRST_NAME, LAST_NAME, EMAIL, to_char(DATE_CREATED, 'DD/MM/YYYY'), to_char(DATE_UPDATED, 'DD/MM/YYYY') FROM USERS WHERE ID = $1"
	getUsersListQry = "SELECT ID,FIRST_NAME, LAST_NAME, EMAIL, to_char(DATE_CREATED, 'DD/MM/YYYY'), to_char(DATE_UPDATED, 'DD/MM/YYYY') FROM USERS"
)

func CreateUser(user *model.User) (*model.User, *error.RestError) {
	var newUser model.User
	err := postgress.DBPool.QueryRow(context.Background(), createUSerQry, user.FirstName, user.LastName, user.Email).Scan(&newUser.Id, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.DateCreated, &newUser.DateUpdated)
	if err != nil {
		return nil, &error.RestError{
			Message:    "Error while saving user in database",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}

	}
	return &newUser, nil
}

func GetUser(userId int64) (*model.User, *error.RestError) {
	var user model.User
	err := postgress.DBPool.QueryRow(context.Background(), getUserQry, userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.DateUpdated)
	if err != nil {
		return nil, &error.RestError{
			Message:    fmt.Sprintf("Could not retrieve the user details for user having Id : %d", userId),
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
	}
	return &user, nil
}

func GetUsersList() (*[]model.User, *error.RestError) {
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

	return &users, nil
}
