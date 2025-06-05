package interfaces

import (
	"finora/database/model"
)

type UserInterface interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(id uint, user *model.User) (*model.User, error)
	DeleteUser(id uint) error
	LogInUser(email, password string) (string, error)
	LogOutUser(token string) error
	ResetPassword(id uint, newPassword string) error
}