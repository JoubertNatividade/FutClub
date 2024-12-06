package commands

import (
	"github.com/JoubertNatividade/FutClub/main/domain/entities"
	"github.com/JoubertNatividade/FutClub/main/domain/interfaces"
)

type PlayerCommand struct {
	repository interfaces.IPlayerRepository
}

func NewPlayerCommand() *PlayerCommand {
	return &PlayerCommand{}
}

func (c *PlayerCommand) Create(user entities.Player) error {

	return nil
}
