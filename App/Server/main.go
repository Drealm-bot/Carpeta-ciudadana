package main

import (
	"fmt"

	"context"

	controller "github.com/Drealm-bot/Carpeta-ciudadana.git/Controllers/Api"
	database "github.com/Drealm-bot/Carpeta-ciudadana.git/Database"
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	repository "github.com/Drealm-bot/Carpeta-ciudadana.git/Repository"
	routes "github.com/Drealm-bot/Carpeta-ciudadana.git/Routes"
	service "github.com/Drealm-bot/Carpeta-ciudadana.git/Services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	c := context.Background()
	database.DBConnection()
	database.DB.AutoMigrate(models.Archive{})
	database.DB.AutoMigrate(models.User{})

	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	archiveRepository := repository.NewArchiveRepository(database.DB)
	archiveService := service.NewArchiveService(archiveRepository)
	archiveController := controller.NewArchiveController(archiveService)

	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/generate" || c.Path() == "/signup" || c.Path() == "/repository/:id/upload" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))

	routes.UserRoutes(app, c, userController)
	routes.ArchiveRoutes(app, c, archiveController)
	app.Logger.Fatal(app.Start(":3000"))
	fmt.Print()
}
