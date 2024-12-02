package main

import "github.com/JoubertNatividade/FutClub/main/infrastructure/controllers"

type App struct {
	controller controllers.PlayerController
}

func NewApp(controller controllers.PlayerController) *App {
	return &App{controller}
}
