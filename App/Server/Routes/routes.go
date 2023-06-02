package routes

import (
	"context"

	controller "github.com/Drealm-bot/Carpeta-ciudadana.git/Controllers/Api"
	"github.com/labstack/echo"
)

func UserRoutes(app *echo.Echo, c context.Context, controller *controller.UserController) {
	app.POST("/signup", controller.RegisterUser)
	app.GET("/repository/:id", controller.ReturnUser)
	app.POST("/generate", controller.GenerateUserPassword)
	app.POST("/login", controller.LoginUser)
}

func ArchiveRoutes(app *echo.Echo, c context.Context, controller *controller.ArchiveController) {
	app.POST("/repository/:id/upload", controller.UploadArchive)
	app.GET("/repository/:id/*", controller.ServeArchive)
	app.GET("/repository/:id/download/*", controller.DownloadArchive)
	app.GET("/repository/:id/authenticate/*", controller.AuthenticateArchive)
}
