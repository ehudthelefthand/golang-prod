package app

import "golang-prod/app/controllers"

type App struct {
	Server *controllers.Server
}

func New() *App {
	return &App{
		Server: controllers.NewServer(),
	}
}

func (a *App) Run() {
	a.Server.Run()
}
