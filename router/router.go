package router

import (
	"net/http"
	"os"
	"pomodoro-api/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uh handler.IUserHandler, taskHandler handler.ITaskHandler, timeHandler handler.ITimeHandler) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", os.Getenv("FE_URL")},
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
	e.GET("/is-auth", uh.IsAuthenticated)
	e.GET("/csrf", uh.CsrfToken)

	// Task
	tasks := e.Group("/tasks")
	tasks.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	tasks.GET("", taskHandler.GetAllTasks)
	tasks.GET("/:taskId", taskHandler.GetTaskById)
	tasks.POST("", taskHandler.CreateTask)
	tasks.PUT("/:taskId", taskHandler.UpdateTask)
	tasks.DELETE("/:taskId", taskHandler.DeleteTask)

	// Time
	times := e.Group("/times")
	times.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	times.GET("/report", timeHandler.GetReport)
	times.POST("", timeHandler.StoreTime)

	return e
}
