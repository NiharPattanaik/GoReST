package controller

import (
	"net/http"
	"strconv"

	"github.com/NiharPattanaik/GoReST/error"
	"github.com/NiharPattanaik/GoReST/model"
	"github.com/NiharPattanaik/GoReST/service"

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

	newUser, err := service.CreateUser(&user)
	if err != nil {
		c.JSON(err.StatusCode, err)
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
	user, restErr := service.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetUserList(c *gin.Context) {
	users, err := service.GetUsersList()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, &users)
}
