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
	app.GET("/user", routes.ReturnUser)
	app.POST("/user", routes.RegisterUser)
	app.POST("/archive", routes.PostArchive)
	app.GET("/archives", routes.GetArchives)
	app.POST("/upload", routes.UploadArchive)
	app.Logger.Fatal(app.Start(":3000"))
	fmt.Print()
}
