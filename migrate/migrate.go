package main

import (
	"fmt"
	"pomodoro-api/db"
	"pomodoro-api/domain"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.User{}, &domain.Task{}, &domain.Time{})
}
