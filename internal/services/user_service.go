package services

import (
	"fiber-test/internal/models"
	"fiber-test/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}


func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}


func (s *UserService) GetUserById(id uint) (models.User, error) {
	return s.repo.GetById(id)
}


func (s *UserService) CreateUser(newUser models.User) (models.User, error) {
	return s.repo.Create(newUser)
}


func (s *UserService) UpdateUser(id uint, user models.User) (models.User, error) {
	return s.repo.Update(id, user)
}


func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}