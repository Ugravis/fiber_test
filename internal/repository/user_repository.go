package repository

import (
	"fiber-test/internal/database"
	"fiber-test/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {}


func NewUserRepository() *UserRepository {
	return &UserRepository{}
}


func (r *UserRepository) Exists(id uint) bool {
	var count int64
	database.DB.Model(&models.User{}).Where("id = ?", id).Count(&count)
	return count > 0
}


func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}


func (r *UserRepository) GetById(id uint) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return user, result.Error
}


func (r *UserRepository) Create(userData models.User) (models.User, error) {
	result := database.DB.Create(&userData)
	return userData, result.Error
}


func (r *UserRepository) Update(id uint, updatedUser models.User) (models.User, error) {
	user, err := r.GetById(id)
	if err != nil {
		return models.User{}, err
	}

	user.Name = updatedUser.Name
	user.Age = updatedUser.Age

	result := database.DB.Save(&user)
	return user, result.Error
}


func (r *UserRepository) Delete(id uint) error {
	if !r.Exists(id) {
		return gorm.ErrRecordNotFound
	}

	result := database.DB.Delete(&models.User{}, id)
	return result.Error
}