package router

import (
	"os"
	"pomodoro-api/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uh handler.IUserHandler, th handler.ITaskHandler) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uh.SignUp)
	e.POST("/login", uh.LogIn)
	e.POST("/logout", uh.LogOut)

	// Task
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", th.GetAllTasks)
	t.GET("/:taskId", th.GetTaskById)
	t.POST("", th.CreateTask)
	t.PUT("/:taskId", th.UpdateTask)
	t.DELETE("/:taskId", th.DeleteTask)

	return e
}
