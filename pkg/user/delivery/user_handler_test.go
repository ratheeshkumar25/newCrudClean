package delivery

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	user "github.com/ratheeshkumar25/pkg/user/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserUseCase is a mock implementation of the UserUseCase interface
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) RegisterUser(user *user.UserRegister) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) Login(login *user.UserLogin) (*user.UserRegister, error) {
	args := m.Called(login)
	return args.Get(0).(*user.UserRegister), args.Error(1)
}

func (m *MockUserUseCase) UpdateUser(user *user.UserRegister) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) GetUserByID(id uint) (*user.UserRegister, error) {
	args := m.Called(id)
	return args.Get(0).(*user.UserRegister), args.Error(1)
}

func (m *MockUserUseCase) RemoveUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserUseCase) GetUserDetail(id uint) (*user.UserRegister, error) {
    args := m.Called(id)
    return args.Get(0).(*user.UserRegister), args.Error(1)
}

func TestRegisterUserHandler(t *testing.T){
	mockUseCase := new(MockUserUseCase)
    handler := NewUserHandler(mockUseCase)

    r := gin.Default()
    r.POST("/signup", handler.RegisterUserHandler)
    
	  // Mock user signup data
    user := user.UserRegister{
        UserName: "ratheeshgk",
        Name:     "Ratheesh G",
        Email:    "ratheeshgk@live1.com",
        Phone:    "9961429911",
        Password: "rathee@123",
    }

	// Setup the mock to return the expected user data
    mockUseCase.On("RegisterUser", &user).Return(nil)

	// Create the request
    jsonValue, _ := json.Marshal(user)
    req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

	 // Perform the request
    r.ServeHTTP(w, req)

	// Check the response status code
    assert.Equal(t, http.StatusOK, w.Code)

	 // Define the expected JSON response
    assert.JSONEq(t, `{"Status":"User registration done successfully"}`, w.Body.String())

}

func TestLoginUserHandler(t *testing.T) {
    mockUseCase := new(MockUserUseCase)
    handler := NewUserHandler(mockUseCase)

    r := gin.Default()
    r.POST("/login", handler.LoginUserHandler)

    // Mock user login data
    userLogin := user.UserLogin{
        UserName: "ratheeshgk",
        Password: "rathee@123",
    }

    // Mock user data returned by the login use case
    user := user.UserRegister{
        UserName: "ratheeshgk",
        Name:     "Ratheesh G",
        Email:    "ratheeshgk@live1.com",
        Phone:    "9961429911",
        Password: "$2a$10$IgtDVCIs6Tx07/0IQ3A5f.UYWOvbw4CEGyukAnESd8rgI8Bc",
    }

    // Setup the mock to return the expected user data
    mockUseCase.On("Login", &userLogin).Return(&user, nil)

    // Create the request
    jsonValue, _ := json.Marshal(userLogin) // Serialize `userLogin` for the request body
    req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)

    // Check the response status code
    assert.Equal(t, http.StatusOK, w.Code)

    // Define the expected JSON response for 
    expectedResponse := `{"Status":"Success","user":{"username":"ratheeshgk","name":"Ratheesh G","email":"ratheeshgk@live1.com","phone":"9961429911"}}`
    assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestUpdateUserHandler(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.PUT("/userupdate", handler.UpdateUserHandler)

	existingUser := user.UserRegister{
		Model:    gorm.Model{ID: 1},
		UserName: "ratheeshgk",
		Name:     "Ratheesh G",
		Email:    "ratheeshgk@live1.com",
		Phone:    "9961429911",
		Password: "rathee@123",
	}

	updateUser := user.UserRegister{
		Model:    gorm.Model{ID: 1},
		UserName: "ratheeshgku",
		Name:     "Ratheesh GK",
		Email:    "ratheeshgk@live12.com",
		Phone:    "9961429921",
		Password: "rathee@1234",
	}

	// Mock the UpdateUser method
	mockUseCase.On("UpdateUser", &existingUser).Return(nil)
	// Mock the GetUserDetail method (corrected from GetUserDetails to GetUserDetail)
	mockUseCase.On("GetUserDetail", uint(1)).Return(&updateUser, nil)

	// Convert existingUser to JSON
	jsonValue, _ := json.Marshal(existingUser)
	req, _ := http.NewRequest("PUT", "/userupdate", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Ensure the status code is OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Expected response (corrected the JSON format and content)
	expectedResponse := `{"Status":"User details updated successfully","user":{"username":"ratheeshgku","name":"Ratheesh GK","email":"ratheeshgk@live12.com","phone":"9961429921"}}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestDeleteUserHandler(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	handler := NewUserHandler(mockUseCase)

	r := gin.Default()
	r.DELETE("/userdelete/:id", handler.DeleteUserHandler)

	mockUseCase.On("RemoveUser", uint(1)).Return(nil)

	req, _ := http.NewRequest("DELETE", "/userdelete/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"Status":"User details deleted successfully"}`, w.Body.String())
}






	
