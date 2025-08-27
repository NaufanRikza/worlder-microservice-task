package repository

import (
	"sensor-consumer/core/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetByUsername(username string) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetByUsername(username string) (entity.User, error) {
	var user entity.User
	err := r.DB.Model(&user).Where("username = ?", username).First(&user).Error
	return user, err
}
