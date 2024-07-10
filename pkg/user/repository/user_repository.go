package repository

import (
	"fmt"

	user "github.com/ratheeshkumar25/pkg/user/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *user.UserRegister) error
	FindUserByName(username string) (*user.UserRegister, error)
	UpdateUser(user *user.UserRegister) error
	GetUserByID(id uint) (*user.UserRegister, error)
	DeleteUser(id int) error
}

type UserDataBaseInteraction struct {
	DB *gorm.DB
}

func (u *UserDataBaseInteraction) CreateUser(user *user.UserRegister) error {
	//**Hash the password before storing to DB
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password %W", err)
	}
	user.Password = string(hashedPassword)

	//**Create the user in the Database
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserDataBaseInteraction) UpdateUser(user *user.UserRegister) error {
//Ensure that the ID is set in the User Object
	if user.ID == 0{
		return fmt.Errorf("user ID is not set")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password %W", err)
	}
	user.Password = string(hashedPassword)
//Use Model and specify the ID explicity 
	result := u.DB.Model(&user).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("updating the user: %w", result.Error)
	}
	return nil
}

func (u *UserDataBaseInteraction) FindUserByName(username string) (*user.UserRegister, error) {
	var user *user.UserRegister
	if err := u.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDataBaseInteraction) GetUserByID(id uint) (*user.UserRegister, error) {
	var user user.UserRegister
	if err := u.DB.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("getting user by ID :%w", err)
	}
	return &user, nil
}

func (u *UserDataBaseInteraction) DeleteUser(id int) error {
	result := u.DB.Delete(&user.UserRegister{}, id)

	if result.Error != nil {
		return fmt.Errorf("deleting the user:%w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with ID: %d", id)
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserDataBaseInteraction{
		DB: db,
	}
}
