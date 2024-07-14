package routes

import (
	"github.com/ratheeshkumar25/pkg/server"
	"github.com/ratheeshkumar25/pkg/user/delivery"
)

type UserRoutes struct {
	Server *server.Server
	User   delivery.UserUseCases
}

func (u *UserRoutes) UsersRoutes() {
	u.Server.R.POST("/signup", u.User.RegisterUserHandler)
	u.Server.R.POST("/login", u.User.LoginUserHandler)
	u.Server.R.PUT("/usersupdate", u.User.UpdateUserHandler)
	u.Server.R.DELETE("/userdelete/:id", u.User.DeleteUserHandler)
}

func NewUserInit(server *server.Server, user *delivery.UserHandler) *UserRoutes {
	return &UserRoutes{
		Server: server,
		User:   user,
	}
}
