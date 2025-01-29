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
	log.Infof("starting create player controller...")
	var request requests.PlayerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Errorf("error on bind request: %s", err)
		responses.BadRequest(c)
		return
	}

	player := mappers.MapToEntityPlayer(request)
	err := self.command.CreateCommand(player)
	if err != nil {
		log.Errorf("player controller -> create ->error on create player: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Created(c)
}
