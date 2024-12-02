package commands

import "github.com/JoubertNatividade/FutClub/main/domain/entities"

type PlayerCommand struct{}

func NewPlayerCommand() *PlayerCommand {
	return &PlayerCommand{}
}

func (c *PlayerCommand) Create(user entities.Player) error {
	return nil
}
