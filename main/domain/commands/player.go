package commands

import "github.com/JoubertNatividade/FutClub/main/domain/entities"

type IPlayerCommand interface {
	CreateCommand(user *entities.Player) error
	ListCommand() ([]entities.Player, error)
}
