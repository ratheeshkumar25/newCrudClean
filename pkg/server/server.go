package server

import "github.com/gin-gonic/gin"

type Server struct {
	R *gin.Engine
}

func (s *Server) StartServer() {
	s.R.Run("localhost:3000")
}

func NewHTTPServer() *Server {
	router := gin.Default()
	return &Server{
		R: router,
	}
}
