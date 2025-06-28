package controllers

import (
	"golang-prod/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	serv := &Server{
		Router: gin.Default(),
		Config: config,
	}

	serv.SetupRoutes()

	return serv
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}
