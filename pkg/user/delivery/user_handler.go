package delivery

import (
	"strconv"

	"github.com/gin-gonic/gin"
	user "github.com/ratheeshkumar25/pkg/user/entity"
	"github.com/ratheeshkumar25/pkg/user/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

type UserUseCases interface {
	RegisterUserHandler(c *gin.Context)
	LoginUserHandler(c *gin.Context)
	UpdateUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
}

func (u *UserHandler) RegisterUserHandler(c *gin.Context) {
	var user user.UserRegister
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"Error": "binding error"})
		return
	}

	err := u.userUseCase.RegisterUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"Error": "user already exists"})
		return
	}
	c.JSON(200, gin.H{"Status": "User registration done successfully"})
}

func (u *UserHandler) LoginUserHandler(c *gin.Context) {
	var userLogin user.UserLogin
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"Error": "binding error"})
		return
	}

	user, err := u.userUseCase.Login(&userLogin)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Status": "Success", "user": gin.H{
		"username": user.UserName,
		"name":     user.Name,
		"email":    user.Email,
		"phone": user.Phone,
	}})
}

func (u *UserHandler) UpdateUserHandler(c *gin.Context) {
	var existinguser user.UserRegister
	if err := c.BindJSON(&existinguser); err != nil {
		c.JSON(400, gin.H{"Error": "binding error"})
		return
	}

	err := u.userUseCase.UpdateUser(&existinguser)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	// Fetch the user details
	user, err := u.userUseCase.GetUserDetail(existinguser.ID)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Status": "User details updated successfully", "user": gin.H{
		"username": user.UserName,
		"name":     user.Name,
		"email":    user.Email,
		"phone":    user.Phone,
	}})
}

func (u *UserHandler) DeleteUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"Error": "Invalid UserID"})
		return
	}

	err = u.userUseCase.RemoveUser(uint(id))
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Status": "User details deleted successfully"})
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}
