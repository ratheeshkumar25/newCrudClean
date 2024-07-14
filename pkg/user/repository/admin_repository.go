package repository

import (
	"fmt"
	"log"

	user "github.com/ratheeshkumar25/pkg/user/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin *user.AdminRegister) error
	GetAdminByUsername(username string) (*user.AdminRegister, error)
	FindAdmin(username string)(*user.AdminRegister,error)
	GetUserList(username string) (*[]user.UserRegister, error)
	AddProduct(product *user.Product) error
	GetProducts(productname string) (*[]user.Product, error)
	FindProduct(id uint) (*user.Product, error)
	UpdateProduct(product *user.Product) error
	DeleteProduct(id int) error
	
}

type AdminDataBaseInteraction struct {
	DB *gorm.DB
}

func (admn *AdminDataBaseInteraction) CreateAdmin(admin *user.AdminRegister) error {
	//HAsh password before storing into DB

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password %W", err)
	}

	admin.Password = string(hashedPassword)

	//**create the admin in Database
	result := admn.DB.Create(&admin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (admn *AdminDataBaseInteraction) GetAdminByUsername(username string) (*user.AdminRegister, error) {
	var admin user.AdminRegister
	if err := admn.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, fmt.Errorf("unable to find admin: %w", err)
	}
	return &admin, nil
}


func (admn *AdminDataBaseInteraction)FindAdmin(username string)(*user.AdminRegister,error){
	var admin *user.AdminRegister
	if err := admn.DB.Where("username =?",username).First(&admin).Error;err != nil{
		return nil,err
	}
	return admin,nil
}


func (admn *AdminDataBaseInteraction) GetUserList(username string) (*[]user.UserRegister, error) {
	var users []user.UserRegister
	log.Printf("Executing query with name: %s", username)
	if err := admn.DB.Where("name LIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("unable to find userlist: %w", err)
	}
	log.Printf("Found users: %v", users)
	return &users, nil
}

func (admn *AdminDataBaseInteraction) AddProduct(product *user.Product) error {
	result := admn.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (admn *AdminDataBaseInteraction) GetProducts(productname string) (*[]user.Product, error) {
	var products []user.Product
	log.Printf("Executing query with name: %s", productname)
	if err := admn.DB.Where("product_name LIKE ?","%"+productname+"%").Find(&products).Error; err != nil {
		return nil, fmt.Errorf("unable to find products: %w", err)
	}
	return &products, nil
}

func (admn *AdminDataBaseInteraction)FindProduct(id uint) (*user.Product, error){
	var product user.Product

	if err := admn.DB.First(&product,id).Error;err!= nil{
		return nil, fmt.Errorf("unable to find product by ID:%w",err)
	}
	return &product, nil
}

func (admn *AdminDataBaseInteraction)UpdateProduct(product *user.Product) error{
	if product.ID == 0{
		return fmt.Errorf("productID is not set")
	}

	result := admn.DB.Model(&product).Where("id = ?",product.ID).Updates(product)
	if result.Error != nil{
		return fmt.Errorf("updating the product is faile")
	}
	return nil
}

func (admn *AdminDataBaseInteraction)DeleteProduct(id int) error{
	if err := admn.DB.Delete(&user.Product{},id).Error; err != nil{
		return fmt.Errorf("unable to delete product:%W",err)
	}
	return nil 
}



func NewAdminUserRepository(db *gorm.DB)AdminRepository{
	return &AdminDataBaseInteraction{
		DB: db,
	}
} 


