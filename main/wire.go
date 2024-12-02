//go:build !test && wireinject

package main

func initApp() *App {
	return nil
}

func (a *App) Run() {
	a.controller.CreatePlayer()
}
