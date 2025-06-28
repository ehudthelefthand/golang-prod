package controllers

import (
	"golang-prod/app/models"
	"golang-prod/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
	Store  models.Store
}

func NewServer(config *config.Config, store models.Store) *Server {
	serv := &Server{
		Router: gin.Default(),
		Config: config,
		Store:  store,
	}

	serv.SetupRoutes()

	return serv
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}
