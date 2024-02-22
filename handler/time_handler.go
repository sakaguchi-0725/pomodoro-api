package handler

import (
	"net/http"
	"pomodoro-api/domain"
	"pomodoro-api/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITimeHandler interface {
	GetAllTimes(c echo.Context) error
	StoreTime(c echo.Context) error
}

type timeHandler struct {
	tu usecase.ITimeUsecase
}

func NewTimeHandler(tu usecase.ITimeUsecase) ITimeHandler {
	return &timeHandler{tu}
}

func (th *timeHandler) GetAllTimes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	timeRes, err := th.tu.GetAllTimes(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, timeRes)
}

func (th *timeHandler) StoreTime(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	time := domain.Time{}
	if err := c.Bind(&time); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	time.UserId = uint(userId.(float64))
	timeRes, err := th.tu.StoreTime(time)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, timeRes)
}
