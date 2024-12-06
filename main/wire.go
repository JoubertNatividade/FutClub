//go:build !test && wireinject

package main

import "github.com/google/wire"

func initApp() *App {
	return nil
}

func (a *App) initApp() {
	wire.Build(
		NewApp,
	)
}
