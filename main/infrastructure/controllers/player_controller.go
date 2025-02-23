package controllers

import (
	"strconv"
	"strings"

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

func (pc *PlayerController) Create(c *gin.Context) {
	log.Infof("starting create player controller...")
	var request requests.PlayerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Errorf("error on bind request: %s", err)
		responses.BadRequest(c)
		return
	}
	log.Infof("--request received in create controller: %+v", request)

	player := mappers.MapToEntityPlayer(request)
	err := pc.command.CreateCommand(player)
	if err != nil {
		if strings.Contains(err.Error(), "player already exist") {
			log.Errorf("player %s %s - %s already exist", player.Name, player.LastName, player.Position)
			responses.Conflict(c)
			return
		}
		log.Errorf("player controller -> create: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Created(c)
}

func (pc *PlayerController) List(c *gin.Context) {
	log.Infof("starting list all player controller...")
	players, err := pc.command.ListCommand()
	if err != nil {
		log.Errorf("player controller -> list: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Success(c, players)
}

func (pc *PlayerController) FindByID(c *gin.Context) {
	log.Infof("starting find player by id controller...")

	idParse, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorf("error on parse id: %s", err)
		responses.BadRequest(c)
		return
	}
	if idParse == 0 {
		log.Error("id is required")
		responses.BadRequest(c)
		return
	}
	player, err := pc.command.FindByIDCommand(idParse)
	if err != nil {
		log.Errorf("player controller -> findByID: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Success(c, player)
}

func (pc *PlayerController) Update(c *gin.Context) {
	log.Infof("starting update player controller...")

	idParse, _ := strconv.Atoi(c.Param("id"))
	if idParse == 0 {
		log.Error("id is required")
		responses.BadRequest(c)
		return
	}

	var request requests.PlayerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Errorf("error on bind request: %s", err)
		responses.BadRequest(c)
		return
	}
	log.Infof("--id received in update controller: %d", idParse)
	log.Infof("--request received in update controller: %+v", request)

	player := mappers.MapToEntityPlayer(request)

	err := pc.command.UpdateCommand(idParse, player)
	if err != nil {
		log.Errorf("player controller -> update: %s", err)
		responses.InternalServerError(c)
		return
	}
	responses.Success(c, "")
}
