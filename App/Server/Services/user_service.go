package service

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	controller "github.com/Drealm-bot/Carpeta-ciudadana.git/Controllers/Utils"
	database "github.com/Drealm-bot/Carpeta-ciudadana.git/Database"
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	repository "github.com/Drealm-bot/Carpeta-ciudadana.git/Repository"
	"github.com/golang-jwt/jwt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}

func (us *UserService) RegisterUser(u *models.User) (int, error) {
	idStr := strconv.Itoa(u.CivID)
	resp, err := http.Get(controller.ValidateCitizen + idStr)
	if err != nil {
		log.Printf("ReqRequest Failed: %s", err)
		return http.StatusInternalServerError, err
	}

	status := resp.StatusCode

	switch status {
	case http.StatusOK:
		return status, nil

	case http.StatusNoContent:
		validateUser := &models.Citizen{
			Id:           u.CivID,
			Name:         u.Name,
			Address:      u.Address,
			Email:        u.Email,
			OperatorId:   controller.OperatorId,
			OperatorName: controller.OperatorName,
		}
		url := controller.RegisterCitizen
		jsonData, err := json.Marshal(validateUser)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		regResp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return http.StatusInternalServerError, err
		}

		if regResp.StatusCode != http.StatusCreated {
			return resp.StatusCode, err
		}

		database.DB.Create(u)

		return regResp.StatusCode, nil
	}
	return http.StatusBadRequest, nil
}

func (us *UserService) ReturnUser(civId int) (*models.User, error) {
	u, err := us.userRepository.GetUserByCivID(civId)
	if err != nil {
		return nil, err
	}

	database.DB.Model(&u).Association("Archives").Find(&u.Archives)

	return u, nil
}

func (us *UserService) GenerateUserPassword(gi models.GenerativeInfo) (int, error) {
	u, err := us.userRepository.GetUserByCivIDAndEmail(gi.CivID, gi.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if u.PasswordCreatedAt == nil || time.Since(*u.PasswordCreatedAt) > 5*time.Minute {
		pass := GenerateRandomPassword(16)
		password := string(pass)
		u.Password = &password
		now := time.Now()
		u.PasswordCreatedAt = &now

		if err := us.userRepository.UpdateUser(*u); err != nil {
			return http.StatusInternalServerError, err
		}

		SendEmail(password, u.Email)

		return http.StatusCreated, nil
	}

	return http.StatusOK, nil
}

func (us *UserService) LoginUser(li models.LoginInfo) (*models.User, error) {
	fmt.Println("tongo")
	u, err := us.userRepository.GetUserByCivIDAndPassword(li.CivID, li.Password)
	fmt.Println("tongo")
	if err != nil {
		return nil, err
	}

	if u.PasswordCreatedAt != nil && time.Since(*u.PasswordCreatedAt) <= 5*time.Minute {
		fmt.Println("tongo")
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = u.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		signedToken, err := token.SignedString([]byte("secret"))
		if err != nil {
			return nil, err
		}
		u.Token = signedToken
		u.Password = nil
		return u, nil
	}

	return nil, errors.New("El usuario ha ingresado exitosamente")
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
