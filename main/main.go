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
		players.GET("", apps.playerController.List)
		players.GET("/:id", apps.playerController.FindByID)
		players.PATCH("/:id", apps.playerController.Update)
		players.DELETE("/:id", apps.playerController.Delete)
	}
	engine.Run(":3000")
}
