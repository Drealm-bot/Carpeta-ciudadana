package controller

import (
	"net/http"
	"strconv"

	service "github.com/Drealm-bot/Carpeta-ciudadana.git/Services"
	"github.com/labstack/echo"
)

type ArchiveController struct {
	archiveService *service.ArchiveService
}

func NewArchiveController(archiveService *service.ArchiveService) *ArchiveController {
	return &ArchiveController{archiveService: archiveService}
}

func (ac *ArchiveController) UploadArchive(c echo.Context) error {
	civId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	resp, err := ac.archiveService.UploadArchive(civId, file)
	if err != nil {
		return c.JSON(resp, err)
	}
	return c.JSON(resp, map[string]interface{}{
		"message":  "Archivo cargado exitosamente",
		"filename": file.Filename,
		"size":     file.Size,
	})
}

func (ac *ArchiveController) ServeArchive(c echo.Context) error {
	id := c.Param("id")
	fileName := c.Param("*")
	fullPath, _ := ac.archiveService.FindArchive(id, fileName)
	return c.File(fullPath)
}

func (ac *ArchiveController) DownloadArchive(c echo.Context) error {
	id := c.Param("id")
	fileName := c.Param("*")
	fullPath, name := ac.archiveService.FindArchive(id, fileName)
	return c.Attachment(fullPath, name)
}
