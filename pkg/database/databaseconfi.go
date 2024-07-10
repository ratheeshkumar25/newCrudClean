package database

import (
	"fmt"

	user "github.com/ratheeshkumar25/pkg/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	host := "localhost"
	userr := "postgres"
	password := "revathy25"
	dbname := "newusers"
	port := "5432"
	sslmode := "disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, userr, password, dbname, port, sslmode)

	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection to the database failed:", err)
	}
	DB.AutoMigrate(user.UserRegister{})
	return DB

}
