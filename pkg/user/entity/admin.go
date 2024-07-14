package user

import "gorm.io/gorm"

type AdminRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	gorm.Model
	ProductName  string  `gorm:"type:varchar(255);not null" json:"product_name"`
	Description  string  `gorm:"type:text" json:"description"`
	Quantity     int     `gorm:"type:int" json:"quantity"`
	Price        float32 `gorm:"type:decimal(10,2)" json:"price"`
	CategoryID   uint    `gorm:"not null" json:"category_id"`
}

