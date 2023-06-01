package routes

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	database "github.com/Drealm-bot/Carpeta-ciudadana.git/Database"
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	"github.com/golang-jwt/jwt"
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
	civid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	u := new(models.User)

	/*if err := database.DB.Preload("Archive").Where("civ_id = ?", civid).First(&u).Error; err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}*/

	if err := database.DB.Where("civ_id = ?", civid).First(&u).Error; err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	database.DB.Model(&u).Association("Archives").Find(&u.Archives)

	return c.JSON(http.StatusOK, u)
}

func RegisterUser(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id := u.CivID
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
			Id:           u.CivID,
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

func GeneratePassword(c echo.Context) error {
	type GenerativeInfo struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}

	l := new(GenerativeInfo)
	if err := c.Bind(l); err != nil {
		return err
	}

	user := new(models.User)
	if err := database.DB.Where("user_id = ? AND email = ?", l.Id, l.Email).First(user).Error; err != nil {
		return err
	}

	if user.PasswordCreatedAt != nil && time.Since(*user.PasswordCreatedAt) <= 5*time.Minute {
		// Se ha generado una contraseña en el período de tiempo deseado
		// Utilizar la contraseña existente
		return c.JSON(http.StatusCreated, l)
	} else {
		pass := GenerateRandomPassword(16)
		password := string(pass)
		user.Password = &password
		now := time.Now()
		user.PasswordCreatedAt = &now

		if err := database.DB.Save(user).Error; err != nil {
			return err
		}

		SendEmail(password, l.Email)
		return c.JSON(http.StatusCreated, l)
	}
}

func LoginUser(c echo.Context) error {
	type LoginInfo struct {
		Id       int    `json:"id"`
		Password string `json:"password"`
	}

	l := new(LoginInfo)
	if err := c.Bind(l); err != nil {
		return err
	}

	u := new(models.User)
	if err := database.DB.Where("user_id = ? AND password = ?", l.Id, l.Password).First(u).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, "Documento o contraseña invalidos.")
	}

	if u.PasswordCreatedAt != nil && time.Since(*u.PasswordCreatedAt) <= 5*time.Minute {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = u.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		signedToken, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		u.Token = signedToken
		u.Password = nil
		return c.JSON(http.StatusOK, u)
	} else {
		return c.JSON(http.StatusNotAcceptable, "Contraseña expirada. Vuelva a generar una nueva.")
	}
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
	civId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dirPath := "./Repository/" + strconv.Itoa(civId)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := dirPath + "/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	name, ext := SplitFilenameAndExtension(file.Filename)
	a := models.Archive{
		CivID: civId,
		Name:  name,
		Type:  ext,
		Path:  filePath,
	}

	database.DB.Create(&a)
	fmt.Print(a)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Archivo cargado exitosamente",
		"filename": file.Filename,
		"size":     file.Size,
	})
}

func SendEmail(body, emailTo string) {

	from := "darangoh23@gmail.com"
	pass := "vvspgbbxlfskzlqi"

	to := emailTo

	subject := mime.QEncoding.Encode("utf-8", "Contraseña Carpeta Ciudadana")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent to " + to)
}

func GenerateRandomPassword(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	password := base64.URLEncoding.EncodeToString(randomBytes)
	return password[:length]
}

func SplitFilenameAndExtension(filename string) (name string, ext string) {
	parts := strings.Split(filename, ".")
	name = parts[0]
	ext = parts[1]
	return name, ext
}
