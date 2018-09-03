package router

import (
	"awesomeProject/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Server = echo.New()

func SetRouter() {
	// middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	// index
	Server.GET("/", controllers.Welcome)
	// user
	Server.GET("/users", controllers.ListUsers)
	Server.GET("/user/:id", controllers.FindUser)
	Server.POST("/users", controllers.CreateUser)
	Server.PUT("/user/:id", controllers.UpdateUser)
	Server.DELETE("/user/:id", controllers.DeleteUser)
}
