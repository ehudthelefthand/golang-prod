package app

import (
	"golang-prod/app/controllers"
	"golang-prod/config"
)

type App struct {
	Server *controllers.Server
}

func New() *App {
	return &App{
		Server: controllers.NewServer(
			config.LoadConfig(),
		),
	}
}

func (a *App) Run() {
	a.Server.Run()
}
