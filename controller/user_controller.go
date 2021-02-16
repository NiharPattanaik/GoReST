package controller

import (
	"net/http"
	"strconv"

	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
	"github.com/NiharPattanaik/GoReST/service"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := error.RestError{
			Message:    "Invalid paylod",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	ok, errString := user.IsUserValid()
	if !ok {
		c.JSON(http.StatusBadRequest, error.RestError{
			Message:    "Mandatory fields are not present in the request",
			StatusCode: http.StatusBadRequest,
			Error:      errString,
		})
		return
	}

	password := user.Password
	hashedPass, err := hashPassword(password)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	user.Password = hashedPass
	newUser, err := service.UserServiceInterface.CreateUser(&user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)

}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error.RestError{
			Message:    "User ID passed is not in correct format",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		})
		return
	}
	user, restErr := service.UserServiceInterface.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetUserList(c *gin.Context) {
	users, err := service.UserServiceInterface.GetUsersList()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, &users)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	userId, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error.RestError{
			Message:    "User ID passed is not in correct format",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		})
		return
	}
	_, restErr := service.UserServiceInterface.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := error.RestError{
			Message:    "Invalid paylod",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	newUser, restErr := service.UserServiceInterface.UpdateUser(userId, user)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, &newUser)
}

func hashPassword(password string) (string, *error.RestError) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		restErr := &error.RestError{
			Message:    "Invalid paylod",
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
		return "", restErr
	}

	return string(bytes), nil
}
