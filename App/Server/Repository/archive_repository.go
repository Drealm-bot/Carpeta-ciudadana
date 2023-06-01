package repository

import (
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	"gorm.io/gorm"
)

type ArchiveRepository struct {
	db *gorm.DB
}

func NewArchiveRepository(db *gorm.DB) *ArchiveRepository {
	return &ArchiveRepository{db: db}
}

func (ar *ArchiveRepository) CreateArchive(a *models.Archive) error {
	return ar.db.Create(&a).Error
}

func (ar *ArchiveRepository) GetArchiveByCivIDAndFileName(civId string, fileName string) (*models.Archive, error) {
	f := new(models.Archive)
	if err := ar.db.Where("owner = ? AND full_name = ?", civId, fileName).First(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}
