package main

import (
	"github.com/JoubertNatividade/FutClub/global/helpers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("starting application Fut'Art...")

	engine := helpers.GetRouters()
	apps := initializeApp()
	players := engine.Group("/player")
	{
		players.POST("", apps.playerController.Create)
	}
	engine.Run(":3000")
}
