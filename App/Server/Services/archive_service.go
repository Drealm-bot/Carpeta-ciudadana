package service

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	repository "github.com/Drealm-bot/Carpeta-ciudadana.git/Repository"
)

type ArchiveService struct {
	archiveRepository *repository.ArchiveRepository
}

func NewArchiveService(ar *repository.ArchiveRepository) *ArchiveService {
	return &ArchiveService{
		archiveRepository: ar,
	}
}

func (as *ArchiveService) UploadArchive(civId int, file *multipart.FileHeader) (int, error) {
	src, err := file.Open()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer src.Close()

	dirPath := "repository/" + strconv.Itoa(civId)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	filePath := dirPath + "/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return http.StatusInternalServerError, err
	}

	name, ext := SplitFilenameAndExtension(file.Filename)
	a := models.Archive{
		Owner:    civId,
		FullName: file.Filename,
		Name:     name,
		Type:     ext,
		Path:     filePath,
	}

	as.archiveRepository.CreateArchive(&a)
	return http.StatusCreated, nil
}

func (as *ArchiveService) FindArchive(id string, fileName string) (string, string) {
	f, err := as.archiveRepository.GetArchiveByCivIDAndFileName(id, fileName)
	if err != nil {
		return "El archivo no existe.", ""
	}
	return f.Path, fileName
}

func SplitFilenameAndExtension(filename string) (name string, ext string) {
	parts := strings.Split(filename, ".")
	name = parts[0]
	ext = parts[1]
	return name, ext
}
