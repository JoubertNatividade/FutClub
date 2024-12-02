package main

import (
	"github.com/JoubertNatividade/FutClub/global/helpers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting aplication FutClub...")

	engine := helpers.GetRouters()

	engine.Run(":3000")

}
