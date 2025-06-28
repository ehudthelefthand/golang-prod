package app

import (
	"context"
	"golang-prod/app/controllers"
	"golang-prod/app/models"
	"golang-prod/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Server *controllers.Server
}

func New() *App {
	config := config.LoadConfig()
	return &App{
		Server: controllers.NewServer(
			config,
			connectDB(config),
		),
	}
}

func (a *App) Run() {
	a.Server.Run()
}

func connectDB(config *config.Config) models.Store {
	dbConnect, err := pgxpool.ParseConfig(config.DBURI)
	if err != nil {
		panic("Can not connect to database: " + err.Error())
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConnect)
	if err != nil {
		panic("Can not connect to database: " + err.Error())
	}

	store := models.NewStore(connPool)

	return store
}
