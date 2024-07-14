package routes

import (
	"github.com/ratheeshkumar25/pkg/server"
	"github.com/ratheeshkumar25/pkg/user/delivery"
)



type AdminRoutes struct{
	Server *server.Server
	Admin delivery.AdminUseCases
}

func (a AdminRoutes)AdminRoutes(){
	a.Server.R.POST("/adminsignup",a.Admin.RegisterAdminHandler)
	a.Server.R.POST("/adminlogin",a.Admin.LoginAdminHandler)
	a.Server.R.GET("/userlist",a.Admin.GetUserListHandler)
	a.Server.R.POST("/addproduct",a.Admin.AddProductHandler)
	a.Server.R.GET("/getproduct",a.Admin.GetProductHandler)
	a.Server.R.PUT("/productupdate",a.Admin.UpdateProductHandler)
	a.Server.R.DELETE("/productdelet/:id",a.Admin.DeletProductHandler)
}

// NewAdminInit creates a new AdminRoutes instance
func NewAdminInit(server *server.Server, admin delivery.AdminUseCases) *AdminRoutes {
    return &AdminRoutes{
        Server: server,
        Admin:  admin,
    }
}