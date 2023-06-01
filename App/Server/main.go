package main

import (
	"fmt"

	database "github.com/Drealm-bot/Carpeta-ciudadana.git/Database"
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	routes "github.com/Drealm-bot/Carpeta-ciudadana.git/Routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	database.DBConnection()
	database.DB.AutoMigrate(models.Archive{})
	database.DB.AutoMigrate(models.User{})

	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	/*app.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/generate" || c.Path() == "/signup" || c.Path() == "/repository/:id/upload" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))*/

	app.POST("/signup", routes.RegisterUser)
	app.POST("/generate", routes.GeneratePassword)
	app.POST("/login", routes.LoginUser)
	app.POST("/archive", routes.PostArchive)
	app.GET("/repository/:id", routes.ReturnUser)
	app.GET("/archives", routes.GetArchives)
	app.Static("/", "public")
	app.Static("/public", "Public")
	app.POST("/repository/:id/upload", routes.UploadArchive)
	app.Logger.Fatal(app.Start(":3000"))
	fmt.Print()
}
