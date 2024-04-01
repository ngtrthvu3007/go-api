package routes

import (
	"book-store/controllers" //add this

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/users", controllers.GetAllUsers)           // get all user
	e.POST("/user", controllers.CreateUser)            // create user
	e.GET("/user/:userId", controllers.GetAUser)       // get user detail
	e.PUT("/user/:userId", controllers.EditAUser)      // edit user detail
	e.DELETE("/user/:userId", controllers.DeleteAUser) // destroy user
}
