package controller

import (
	"go-temp/go-sample/api/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserController struct {
	uc usecase.UserUsecase
}

func NewUserController(uc usecase.UserUsecase) *UserController {
	return &UserController{
		uc: uc,
	}
}

type (
	SignUpReq struct {
		Name string `json:"name" validate:"required"`
		//
	}
	SignUpResp User
)

func (ct *UserController) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	user, err := ct.uc.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
