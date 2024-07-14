package delivery

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	user "github.com/ratheeshkumar25/pkg/user/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockAdminUseCase is a mock implementation of the AdminUseCase interface
type MockAdminUseCase struct {
	mock.Mock
}

func (m *MockAdminUseCase) RegisterAdmin(admin *user.AdminRegister) error {
	args := m.Called(admin)
	return args.Error(0)
}

func (m *MockAdminUseCase) Login(login *user.AdminLogin) (*user.AdminRegister, error) {
	args := m.Called(login)
	return args.Get(0).(*user.AdminRegister), args.Error(1)
}

func (m *MockAdminUseCase) GetUseList(name string) (*[]user.UserRegister, error) {
	args := m.Called(name)
	return args.Get(0).(*[]user.UserRegister), args.Error(1)
}

func (m *MockAdminUseCase) AddProduct(product *user.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockAdminUseCase) GetProducts(productname string) (*[]user.Product, error) {
	args := m.Called(productname)
	return args.Get(0).(*[]user.Product), args.Error(1)
}

func (m *MockAdminUseCase) FindProduct(id uint) (*user.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*user.Product), args.Error(1)
}

func (m *MockAdminUseCase) UpdateProduct(product *user.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockAdminUseCase) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestRegisterAdminHandler(t *testing.T) {
	mockUseCase := new(MockAdminUseCase)
	handler := NewAdminHandler(mockUseCase)

	admin := &user.AdminRegister{
		Username: "admin1",
		Password: "password123",
		Email:    "admin1@example.com",
	}
	mockUseCase.On("RegisterAdmin", admin).Return(nil)

	router := gin.Default()
	router.POST("/adminsignup", handler.RegisterAdminHandler)

	body, _ := json.Marshal(admin)
	req, _ := http.NewRequest("POST", "/adminsignup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"Status":"Admin registration done successfully"}`, w.Body.String())
}

func TestLoginAdminHandler(t *testing.T) {
	mockUseCase := new(MockAdminUseCase)
	handler := NewAdminHandler(mockUseCase)

	adminLogin := &user.AdminLogin{
		Username: "admin1",
		Password: "password123",
	}
	admin := &user.AdminRegister{
		Username: "admin1",
		Email:    "admin1@example.com",
	}
	mockUseCase.On("Login", adminLogin).Return(admin, nil)

	router := gin.Default()
	router.POST("/adminlogin", handler.LoginAdminHandler)

	body, _ := json.Marshal(adminLogin)
	req, _ := http.NewRequest("POST", "/adminlogin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	expectedResponse := gin.H{"Status": "Success", "admin": gin.H{
		"username": admin.Username,
		"name":     admin.Email,
	}}
	responseBody, _ := json.Marshal(expectedResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(responseBody), w.Body.String())
}

func TestAddProductHandler(t *testing.T) {
	mockUseCase := new(MockAdminUseCase)
	handler := NewAdminHandler(mockUseCase)

	product := &user.Product{
		ProductName: "Product1",
		Price:       18.30,
	}
	mockUseCase.On("AddProduct", product).Return(nil)

	router := gin.Default()
	router.POST("/addproduct", handler.AddProductHandler)

	body, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/addproduct", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message":"product added successfully"}`, w.Body.String())
}

func TestGetProductHandler(t *testing.T) {
	mockUseCase := new(MockAdminUseCase)
	handler := NewAdminHandler(mockUseCase)

	productName := "Product1"
	products := &[]user.Product{
		{ProductName: "Product1", Price: 18.30},
		{ProductName: "Product2", Price: 20.50},
	}
	mockUseCase.On("GetProducts", productName).Return(products, nil)

	router := gin.Default()
	router.GET("/getproduct", handler.GetProductHandler)

	req, _ := http.NewRequest("GET", "/getproduct", nil)
	q := req.URL.Query()
	q.Add("name", productName)
	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseBody, _ := json.Marshal(products)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(responseBody), w.Body.String())
}

func TestUpdateProductHandler(t *testing.T) {
	mockUseCase := new(MockAdminUseCase)
	handler := NewAdminHandler(mockUseCase)

	// Original product details
	product := &user.Product{
		Model:       gorm.Model{ID: 1},
		ProductName: "Product1",
		Description: "TEST",
		Price:       18.300,
		Quantity:    10,
		CategoryID:  1,
	}

	// Updated product details
	updatedProduct := &user.Product{
		Model:       gorm.Model{ID: 1},
		ProductName: "UpdatedProduct",
		Description: "Updated Description",
		Price:       25.500,
		Quantity:    15,
		CategoryID:  1,
	}

	// Mock the FindProduct and UpdateProduct methods
	mockUseCase.On("FindProduct", uint(1)).Return(product, nil)
	mockUseCase.On("UpdateProduct", mock.MatchedBy(func(p *user.Product) bool {
		return p.ProductName == updatedProduct.ProductName &&
			p.Description == updatedProduct.Description &&
			p.Price == updatedProduct.Price &&
			p.Quantity == updatedProduct.Quantity &&
			p.CategoryID == updatedProduct.CategoryID
	})).Return(nil)

	router := gin.Default()
	router.PUT("/productupdate", handler.UpdateProductHandler)

	body, _ := json.Marshal(updatedProduct)
	req, _ := http.NewRequest("PUT", "/productupdate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponse, _ := json.Marshal(updatedProduct)
	assert.JSONEq(t, string(expectedResponse), w.Body.String())
}

func TestDeleteProductHandler(t *testing.T) { // corrected function name
	mockUseCase := new(MockAdminUseCase)
	handler := NewAdminHandler(mockUseCase)

	productId := 1
	mockUseCase.On("DeleteProduct", productId).Return(nil)

	router := gin.Default()
	router.DELETE("/productdelete/:id",handler.DeletProductHandler)
	 // corrected endpoint name

	req, _ := http.NewRequest("DELETE", "/productdelete/1", nil) // corrected endpoint name

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"product deleted successfully"}`, w.Body.String())
}
