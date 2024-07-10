package usecase

import (
	"fmt"

	user "github.com/ratheeshkumar25/pkg/user/entity"
	"github.com/ratheeshkumar25/pkg/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(user *user.UserRegister) error
	Login(login *user.UserLogin) (*user.UserRegister, error)
	UpdateUser(user *user.UserRegister) error
	GetUserDetail(id uint) (*user.UserRegister, error)
	RemoveUser(id uint) error
}

type userInteraction struct {
	userRepo repository.UserRepository
}

func (u *userInteraction) RegisterUser(user *user.UserRegister) error {
	return u.userRepo.CreateUser(user)
}

func (u *userInteraction) Login(login *user.UserLogin) (*user.UserRegister, error) {
	user, err := u.userRepo.FindUserByName(login.UserName)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password:%w", err)
	}
	return user, nil
}


func (u *userInteraction) UpdateUser(user *user.UserRegister) error {
	return u.userRepo.UpdateUser(user)
}

func (u *userInteraction) GetUserDetail(id uint) (*user.UserRegister, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userInteraction) RemoveUser(id uint) error {
	return u.userRepo.DeleteUser(int(id))
}

func NewUserUsecase(userRepo repository.UserRepository) UserUseCase {
	return &userInteraction{
		userRepo: userRepo,
	}
}
