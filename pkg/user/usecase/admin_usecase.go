package usecase

import (
	"fmt"

	user "github.com/ratheeshkumar25/pkg/user/entity"
	"github.com/ratheeshkumar25/pkg/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	RegisterAdmin(admin *user.AdminRegister) error
	Login(login *user.AdminLogin) (*user.AdminRegister, error)
	GetUseList(user string) (*[]user.UserRegister, error)
	AddProduct(product *user.Product) error
	GetProducts(productname string) (*[]user.Product, error)
	FindProduct(id uint) (*user.Product, error)
	UpdateProduct(product *user.Product) error
	DeleteProduct(id int) error
}

type adminInteraction struct {
	adminRepo repository.AdminRepository
}

func (admn *adminInteraction) RegisterAdmin(admin *user.AdminRegister) error {
	return admn.adminRepo.CreateAdmin(admin)
}

func (admn *adminInteraction) Login(login *user.AdminLogin) (*user.AdminRegister, error) {
	admin, err := admn.adminRepo.FindAdmin(login.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(login.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}
	return admin, nil
}

func (admn *adminInteraction) GetUseList(user string) (*[]user.UserRegister, error) {
	users, err := admn.adminRepo.GetUserList(user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}
	return users, nil
}

func (admn *adminInteraction) AddProduct(product *user.Product) error {
	err := admn.adminRepo.AddProduct(product)
	if err != nil {
		return fmt.Errorf("failed to add product: %w", err)
	}
	return nil
}

func (admn *adminInteraction) GetProducts(productname string) (*[]user.Product, error) {
	products, err := admn.adminRepo.GetProducts(productname)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	return products, nil
}

func (admn *adminInteraction) FindProduct(id uint) (*user.Product, error) {
	product, err := admn.adminRepo.FindProduct(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find product: %w", err)
	}
	return product, nil
}

func (admn *adminInteraction) UpdateProduct(product *user.Product) error {
	return admn.adminRepo.UpdateProduct(product)
}

func (admn *adminInteraction) DeleteProduct(id int) error {
	return admn.adminRepo.DeleteProduct(id)
}





func NewAdminUseCase(adminRepo repository.AdminRepository) AdminUseCase {
	return &adminInteraction{
		adminRepo: adminRepo,
	}
}
