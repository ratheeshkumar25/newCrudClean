package di

import (
	"github.com/ratheeshkumar25/pkg/database"
	"github.com/ratheeshkumar25/pkg/routes"
	"github.com/ratheeshkumar25/pkg/server"
	"github.com/ratheeshkumar25/pkg/user/delivery"
	"github.com/ratheeshkumar25/pkg/user/repository"
	"github.com/ratheeshkumar25/pkg/user/usecase"
)

func Init() *server.Server{
	server := server.NewHTTPServer()
	db := database.ConnectDatabase()
	repo :=repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	delivery := delivery.NewUserHandler(usecase)
	userRoutes := routes.NewUserInit(server,delivery)
	userRoutes.UsersRoutes()
	return server
}