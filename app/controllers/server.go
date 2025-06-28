package controllers

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	serv := &Server{
		Router: gin.Default(),
	}

	serv.SetupRoutes()

	return serv
}

func (s *Server) Run() {
	s.Router.Run()
}
