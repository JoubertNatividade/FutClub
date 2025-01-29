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
		log.Errorf("player controller -> create: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Created(c)
}

func (self *PlayerController) List(c *gin.Context) {
	log.Infof("starting list all player controller...")
	players, err := self.command.ListCommand()
	if err != nil {
		log.Errorf("player controller -> list: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Success(c, players)
}
