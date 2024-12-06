package controllers

import (
	"github.com/JoubertNatividade/FutClub/main/domain/commands"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/controllers/mappers"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/controllers/requests"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type PlayerController struct {
	command commands.PlayerCommand
}

func NewPlayerController(
	command commands.PlayerCommand,
) *PlayerController {
	return &PlayerController{command}
}

func (self *PlayerController) Create(c *gin.Context) error {
	var request requests.PlayerRequest
	// if err := request.Validate(); err != nil {
	// 	return err
	// }

	player := mappers.MapToEntityPlayer(request)
	err := self.command.Create(player)
	if err != nil {
		log.Error(err)
	}
	return nil
}
