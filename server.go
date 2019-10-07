package main

import (
	"blog/pkg/controllers/auth"
	"blog/pkg/controllers/registrationuser"
	db "blog/pkg/sqlite"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db.Initialize()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/auth/login", auth.Authentication)
	e.POST("/api/registration", registrationuser.Registration)

	e.Logger.Fatal(e.Start(":9999"))
}
