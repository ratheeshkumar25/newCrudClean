package delivery

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	user "github.com/ratheeshkumar25/pkg/user/entity"
	"github.com/ratheeshkumar25/pkg/user/usecase"
)

type AdminHandler struct {
	adminUseCase usecase.AdminUseCase
}

type AdminUseCases interface {
	RegisterAdminHandler(c *gin.Context)
	LoginAdminHandler(c *gin.Context)
	GetUserListHandler(c *gin.Context)
	AddProductHandler(c *gin.Context)
	GetProductHandler(c *gin.Context)
	UpdateProductHandler(c *gin.Context)
	DeletProductHandler(c *gin.Context)
}

func (a *AdminHandler) RegisterAdminHandler(c *gin.Context) {
	var admin user.AdminRegister

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(400, gin.H{"Error": "bindig error"})
		return
	}

	err := a.adminUseCase.RegisterAdmin(&admin)
	if err != nil {
		c.JSON(500, gin.H{"Error": "admin already register"})
		return
	}
	c.JSON(201, gin.H{"Status": "Admin registration done successfully"})
}

func (a *AdminHandler) LoginAdminHandler(c *gin.Context) {
	var adminLogin user.AdminLogin

	if err := c.BindJSON(&adminLogin); err != nil {
		c.JSON(400, gin.H{"Error": "Binding Error"})
		return
	}

	admin, err := a.adminUseCase.Login(&adminLogin)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Wrong UserName and Password"})
		return
	}

	c.JSON(200, gin.H{"Status": "Success", "admin": gin.H{
		"username": admin.Username,
		"name":     admin.Email,
	}})
}

func (a *AdminHandler) GetUserListHandler(c *gin.Context) {
	user := c.Query("name")
	users, err := a.adminUseCase.GetUseList(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (a *AdminHandler) AddProductHandler(c *gin.Context) {
	var product user.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := a.adminUseCase.AddProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(201, gin.H{"message": "Product added successfully"})

}

func (a *AdminHandler) GetProductHandler(c *gin.Context) {
	productname := c.Query("name")
	products, err := a.adminUseCase.GetProducts(productname)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, products)
}

func (h *AdminHandler) UpdateProductHandler(c *gin.Context) {
	var product user.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existingProduct, err := h.adminUseCase.FindProduct(product.ID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	// Update the existing product fields with the new values
	existingProduct.ProductName = product.ProductName
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.Quantity = product.Quantity
	existingProduct.CategoryID = product.CategoryID

	if err := h.adminUseCase.UpdateProduct(existingProduct); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(200, existingProduct)
}

func (a *AdminHandler) DeletProductHandler(c *gin.Context) {
	idStr := c.Param("id")
	log.Printf("Executing query with id %s", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := a.adminUseCase.DeleteProduct(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "product deleted successfully"})
}

func NewAdminHandler(adminUseCase usecase.AdminUseCase) *AdminHandler {
	return &AdminHandler{adminUseCase: adminUseCase}
}
