package router

import (
	"net/http"
	"os"
	"pomodoro-api/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uh handler.IUserHandler, th handler.ITaskHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
	}))
	e.POST("/signup", uh.SignUp)
	e.POST("/login", uh.LogIn)
	e.POST("/logout", uh.LogOut)
	e.GET("/csrf", uh.CsrfToken)

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
