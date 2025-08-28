package repository

import (
	"sensor-consumer/core/entity"

	"gorm.io/gorm"
)

type userRoleRepository struct {
	DB *gorm.DB
}

type UserRoleRepository interface {
	GetRoleByUserID(userID uint64) ([]string, error)
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		DB: db,
	}
}

func (r *userRoleRepository) GetRoleByUserID(userID uint64) ([]string, error) {
	var roles []string
	entity := entity.UserRole{}
	err := r.DB.Table(entity.TableName()).
		Model(&entity).
		Select("r.name").
		Joins("JOIN roles r ON r.id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Pluck("r.name", &roles).Error
	return roles, err
}
