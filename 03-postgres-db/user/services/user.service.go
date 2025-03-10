package services

import (
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/user/models"
	"github.com/lgmamaniap/gorm-tutorial/03-postgres-db/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user models.User) (models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetByID(id uint64) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(user models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
