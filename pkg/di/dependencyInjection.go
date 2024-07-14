// package di

// import (
// 	"github.com/ratheeshkumar25/pkg/database"
// 	"github.com/ratheeshkumar25/pkg/routes"
// 	"github.com/ratheeshkumar25/pkg/server"
// 	"github.com/ratheeshkumar25/pkg/user/delivery"
// 	"github.com/ratheeshkumar25/pkg/user/repository"
// 	"github.com/ratheeshkumar25/pkg/user/usecase"
// )

// func Init() *server.Server{
// 	server := server.NewHTTPServer()
// 	db := database.ConnectDatabase()
// 	repo :=repository.NewUserRepository(db)
// 	usecase := usecase.NewUserUsecase(repo)
// 	delivery := delivery.NewUserHandler(usecase)
// 	userRoutes := routes.NewUserInit(server,delivery)
// 	userRoutes.UsersRoutes()
// 	return server
// }

package di

import (
    "github.com/ratheeshkumar25/pkg/server"
    "github.com/ratheeshkumar25/pkg/user/delivery"
    "github.com/ratheeshkumar25/pkg/user/repository"
    "github.com/ratheeshkumar25/pkg/user/usecase"
    "github.com/ratheeshkumar25/pkg/routes"
    "github.com/ratheeshkumar25/pkg/database"
)

func Init() *server.Server {
    // Initialize the HTTP server
    server := server.NewHTTPServer()

    // Connect to the database
    db := database.ConnectDatabase()

    // Create a new repository instance for Admin
    adminRepo := repository.NewAdminUserRepository(db)

    // Create a new use case instance for Admin
    adminUseCase := usecase.NewAdminUseCase(adminRepo)

    // Create a new handler instance for Admin
    adminHandler := delivery.NewAdminHandler(adminUseCase)

    // Create new routes for Admin and pass in the handler
    adminRoutes := routes.NewAdminInit(server, adminHandler)

    // Setup Admin routes
    adminRoutes.AdminRoutes()

    // Create a new repository instance for User
    userRepo := repository.NewUserRepository(db)

    // Create a new use case instance for User
    userUseCase := usecase.NewUserUsecase(userRepo)

    // Create a new handler instance for User
    userHandler := delivery.NewUserHandler(userUseCase)

    // Create new routes for User and pass in the handler
    userRoutes := routes.NewUserInit(server, userHandler)

    // Setup User routes
    userRoutes.UsersRoutes()

    // Return the initialized server
    return server
}
