package controllers

import (
	"github.com/JoubertNatividade/FutClub/global/infrastructure/responses"
	"github.com/JoubertNatividade/FutClub/main/domain/commands"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/controllers/mappers"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/controllers/requests"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type PlayerController struct {
	command commands.IPlayerCommand
}

func NewPlayerController(command commands.IPlayerCommand) *PlayerController {
	return &PlayerController{command: command}
}

func (self *PlayerController) Create(c *gin.Context) {
	var request requests.PlayerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		responses.BadRequest(c)
	}

	player := mappers.MapToEntityPlayer(request)
	err := self.command.CreateCommand(player)
	if err != nil {
		log.Errorf("Error when try create player: %s", err)
	}
	responses.Created(c)
}
