package main

import (
	"book-store/configs"
	"book-store/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	configs.ConnectDB()
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":3000"))
}
