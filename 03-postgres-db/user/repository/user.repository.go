package repository

import (
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/user/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(data models.User) (models.User, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return models.User{}, err
	}
	return data, nil
}

func (r *UserRepository) GetByID(id uint64) error {
	return r.db.First(&models.User{}, id).Error
}

func (r *UserRepository) Update(data models.User) error {
	return r.db.Save(&data).Error
}

func (r *UserRepository) Delete(id uint64) error {
	return r.db.Delete(&models.User{}, id).Error
}
