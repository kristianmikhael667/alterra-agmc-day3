package routes

import (
	"main/constants"
	"main/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// noauth
	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBook)

	e.POST("/login", controllers.LoginUserController)

	eAuth := e.Group("/v1")
	eAuth.Use(middleware.JWT([]byte(constants.SecretKey())))
	eAuth.POST("/books", controllers.CreateBooks)
	eAuth.PUT("/books/:id", controllers.UpdateBook)
	eAuth.DELETE("/books/:id", controllers.DeleteBooks)

	return e
}
