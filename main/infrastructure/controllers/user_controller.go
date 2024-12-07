package controllers

import (
	"github.com/JoubertNatividade/FutClub/main/domain/commands"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	command commands.IPlayerCommand
}

func NewPlayerController(command commands.IPlayerCommand) *PlayerController {
	return &PlayerController{command: command}
}

func (self *PlayerController) Create(c *gin.Context) {
	//var request requests.PlayerRequest
	// if err := request.Validate(); err != nil {
	// 	return err
	// }
	//
	//player := mappers.MapToEntityPlayer(request)
	//err := self.command.Create(player)
	//if err != nil {
	//	log.Error(err)
	//}
	c.JSON(200, gin.H{"message": "ok"})
}
