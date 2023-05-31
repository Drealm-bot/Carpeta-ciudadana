package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	database.DB.Find(&u)
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
		url := "http://169.51.195.62:30174/apis/registerCitizen"
		jsonData, err := json.Marshal(validateUser)
		if err != nil {
			return err
		}

		regResp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}

		if regResp.StatusCode != http.StatusCreated {
			return fmt.Errorf("la solicitud POST no fue exitosa. Código de estado: %d", resp.StatusCode)
		}

		database.DB.Create(u)

		return c.JSON(regResp.StatusCode, u)
	}
	return c.JSON(http.StatusBadRequest, u)
}

func GetArchives(c echo.Context) error {
	var archives []models.Archive
	database.DB.Find(&archives)
	return c.JSON(http.StatusOK, archives)
}

func PostArchive(c echo.Context) error {
	a := new(models.Archive)
	if err := c.Bind(a); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	database.DB.Create(&a)
	return c.JSON(http.StatusOK, a)
}

func UploadArchive(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("/Public/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Archivo cargado exitosamente",
		"filename": file.Filename,
		"size":     file.Size,
	})
}
