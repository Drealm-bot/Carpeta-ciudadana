package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	database "github.com/Drealm-bot/Carpeta-ciudadana.git/Database"
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	"github.com/labstack/echo"
)

type validateCitizen struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	OperatorId   int    `json:"operatorId"`
	OperatorName string `json:"operatorName"`
}

func ReturnUser(c echo.Context) error {
	u := &models.User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

func RegisterUser(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id := u.Cc
	idStr := strconv.Itoa(int(id))
	resp, err := http.Get("http://169.51.195.62:30174/apis/validateCitizen/" + idStr)
	if err != nil {
		log.Printf("ReqRequest Failed: %s", err)
		return err
	}
	status := resp.StatusCode
	switch status {
	case http.StatusOK:
		return c.JSON(status, u)
	case http.StatusNoContent:

		validateUser := validateCitizen{
			Id:           u.Cc,
			Name:         u.Name,
			Address:      u.Address,
			Email:        u.Email,
			OperatorId:   5,
			OperatorName: "ciucarp",
		}
		log.Print(validateUser)
		url := "http://169.51.195.62:30174/apis/registerCitizen"
		log.Print("tongo")
		jsonData, err := json.Marshal(validateUser)
		if err != nil {
			return err
		}
		log.Print("rongo")
		regResp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		log.Print("fongo")
		log.Print(regResp.StatusCode)
		if regResp.StatusCode != http.StatusCreated {
			return fmt.Errorf("la solicitud POST no fue exitosa. Código de estado: %d", resp.StatusCode)
		}
		log.Print("nongo")
		database.DB.Create(u)

		return c.JSON(regResp.StatusCode, u)
	}
	return nil
}
