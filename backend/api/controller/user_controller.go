package controller

import (
	"errors"
	"finora/api/interfaces"
	"finora/api/middleware"
	"finora/api/util"
	"finora/database"
	"finora/database/model"
	"fmt"
)

type UserController struct{}

type AuthResponse struct {
    User  *model.User `json:"user"`
    Token string      `json:"token"`
}


func NewUserController() interfaces.UserInterface {
	return &UserController{}
}

func (uc *UserController) CreateUser(user *model.User) (*model.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	if !util.IsValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}
	if !util.IsValidUsername(user.Username) {
		return nil, errors.New("invalid username format")
	}

	if !util.IsValidPassword(user.Password) {
		return nil, errors.New("password must be at least 8 characters long and contain at least one digit")
	}

	user.Password = hashedPassword

	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserController) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (uc *UserController) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %v", err)
	}
	return users, nil
}

func (uc *UserController) UpdateUser(id uint, updatedUser *model.User) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if updatedUser.Email != "" && !util.IsValidEmail(updatedUser.Email) {
		return nil, errors.New("invalid email format")
	}
	if updatedUser.Username != "" && !util.IsValidUsername(updatedUser.Username) {
		return nil, errors.New("invalid username format")
	}
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email

	if err := database.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}


func (uc *UserController) DeleteUser(id uint) error {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (uc *UserController) LogInUser(email, password string) (string, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return " ", errors.New("user does not exist")
	}
	if !util.CheckPassword(user.Password, password) {
		return " ", errors.New("invalid email or password")
	}

	token, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		return " ", fmt.Errorf("failed to generate token: %v", err)
	}

	
	return token, nil
}

func (uc *UserController) LogOutUser(token string) error {
	if err := middleware.InvalidateToken(token); err != nil {
		return fmt.Errorf("failed to log out: %v", err)
	}
	return nil
}

func (uc *UserController) ResetPassword(id uint, newPassword string) error {
	var user model.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return fmt.Errorf("reset failed: user with ID %d not found", id)
	}

	if !util.IsValidPassword(newPassword) {
		return errors.New("password must be at least 8 characters long and contain at least one digit")
	}

	hashedPassword, err := util.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	if err := database.DB.Model(&user).Update("password", hashedPassword).Error; err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

