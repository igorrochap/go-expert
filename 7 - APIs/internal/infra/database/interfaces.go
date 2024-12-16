package database

import "github.com/igorrochap/goexpert/7-APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
