package repository

import (
	"context"
	"sensor-consumer/core/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetUserIDByUsername(ctx context.Context, username string) (uint64, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Model(&user).Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) GetUserIDByUsername(ctx context.Context, username string) (uint64, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Model(&user).Where("username = ?", username).First(&user).Error
	return user.ID, err
}
