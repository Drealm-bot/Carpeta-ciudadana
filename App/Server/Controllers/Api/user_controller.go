package controller

import (
	"net/http"
	"strconv"

	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	service "github.com/Drealm-bot/Carpeta-ciudadana.git/Services"
	"github.com/labstack/echo"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(&u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	status, err := uc.userService.RegisterUser(u)
	if err != nil {
		return c.JSON(status, err)
	}
	return c.JSON(status, u)
}

func (uc *UserController) ReturnUser(c echo.Context) error {
	civid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	u, err := uc.userService.ReturnUser(civid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, u)
}

func (uc *UserController) GenerateUserPassword(c echo.Context) error {
	gi := new(models.GenerativeInfo)
	if err := c.Bind(gi); err != nil {
		return err
	}
	resp, err := uc.userService.GenerateUserPassword(*gi)
	if err != nil {
		return c.JSON(resp, err)
	}

	return c.JSON(http.StatusCreated, gi)
}

func (uc *UserController) LoginUser(c echo.Context) error {
	li := new(models.LoginInfo)
	if err := c.Bind(li); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	u, err := uc.userService.LoginUser(*li)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, u)
}
