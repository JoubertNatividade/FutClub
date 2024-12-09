//go:build !test && wireinject

package main

import (
	"github.com/JoubertNatividade/FutClub/main/domain/commands"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/controllers"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/repositories"
	"github.com/google/wire"
)

func initializeApp() *App {
	wire.Build(
		ConnectionMySQL,
		repositories.Container,
		commands.Container,
		controllers.Container,
		NewApp,
	)
	return nil
}
