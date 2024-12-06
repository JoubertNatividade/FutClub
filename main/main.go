package main

import (
	"github.com/JoubertNatividade/FutClub/global/helpers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting aplication FutClub...")

	engine := helpers.GetRouters()

	apps := initApp()

	players := engine.Group("/player")
	{
		players.POST("", apps.controller.Create)
	}

	engine.Run(":3000")

}
