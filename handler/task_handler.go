package handler

import (
	"net/http"
	"pomodoro-api/domain"
	"pomodoro-api/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskHandler interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskHandler struct {
	tu usecase.ITaskUseCase
}

func NewTaskHandler(tu usecase.ITaskUseCase) ITaskHandler {
	return &taskHandler{tu}
}

func (th *taskHandler) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := th.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (th *taskHandler) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	taskRes, err := th.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (th *taskHandler) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := domain.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId.(float64))
	taskRes, err := th.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, taskRes)
}

func (th *taskHandler) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := domain.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := th.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (th *taskHandler) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := th.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
