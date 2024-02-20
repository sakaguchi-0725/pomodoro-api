package main

import (
	"pomodoro-api/db"
	"pomodoro-api/handler"
	"pomodoro-api/repository"
	"pomodoro-api/router"
	"pomodoro-api/usecase"
	"pomodoro-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userHandler := handler.NewUserHandler(userUsecase)

	// Task
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := router.NewRouter(userHandler, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
