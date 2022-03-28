package repository

import (
	"github.com/alextnetto/pkg/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *entity.User) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (u userRepository) Create(newUser *entity.User) (*entity.User, error) {
	result := u.db.Create(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}
