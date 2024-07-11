package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	user "github.com/ratheeshkumar25/pkg/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
 
// Attempt to load .env file, for connecting DB
_ = godotenv.Load()

// Retrieve DSN from environment variables
dsn := os.Getenv("DSN")
if dsn == "" {
	log.Fatal("DSN environment variable not set")
}

// Debug print to check if DSN is correctly loaded
fmt.Println("DSN:", dsn)

// Open a connection to the database
DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatalf("Connection to the database failed: %v", err)
}

DB.AutoMigrate(&user.UserRegister{})
return DB

}

// func getEnv(key,fallback string)string{
// 	if value,exists := os.LookupEnv(key);exists{
// 		return value
// 	}
// 	return fallback
// }
