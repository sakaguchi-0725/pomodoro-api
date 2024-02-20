package main

import (
	"pomodoro-api/db"
	"pomodoro-api/handler"
	"pomodoro-api/repository"
	"pomodoro-api/router"
	"pomodoro-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e := router.NewRouter(userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
