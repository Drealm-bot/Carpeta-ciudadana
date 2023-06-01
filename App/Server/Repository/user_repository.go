package repository

import (
	models "github.com/Drealm-bot/Carpeta-ciudadana.git/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUserByCivID(civId int) (*models.User, error) {
	var u models.User
	err := ur.db.Where("civ_id = ?", civId).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (ur *UserRepository) GetUserByCivIDAndEmail(civId int, email string) (*models.User, error) {
	var u models.User
	if err := ur.db.Where("civ_id = ? AND email = ?", civId, email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (ur *UserRepository) GetUserByCivIDAndPassword(civId int, password string) (*models.User, error) {
	var u models.User
	if err := ur.db.Where("civ_id = ? AND password = ?", civId, password).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (ur *UserRepository) UpdateUser(u models.User) error {
	if err := ur.db.Save(u).Error; err != nil {
		return err
	}
	return nil
}
