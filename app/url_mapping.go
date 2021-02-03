package app

import "github.com/NiharPattanaik/GoReST/controller"

func mapURLS() {
	router.POST("/users", controller.CreateUser)
	router.GET("/users/:userID", controller.GetUser)
	router.GET("/users", controller.GetUserList)
}
