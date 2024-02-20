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

	// Task
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := router.NewRouter(userHandler, taskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
