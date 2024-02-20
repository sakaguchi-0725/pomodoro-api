package router

import (
	"pomodoro-api/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(uh handler.IUserHandler) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uh.SignUp)
	e.POST("/login", uh.LogIn)
	e.POST("/logout", uh.LogOut)

	return e
}
