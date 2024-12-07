package main

import "github.com/JoubertNatividade/FutClub/main/infrastructure/controllers"

type App struct {
	playerController *controllers.PlayerController
}

func NewApp(playerController *controllers.PlayerController) *App {
	return &App{playerController: playerController}
}
