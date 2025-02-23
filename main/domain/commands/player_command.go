package commands

import (
	"fmt"

	"github.com/JoubertNatividade/FutClub/main/domain/entities"
	"github.com/JoubertNatividade/FutClub/main/domain/interfaces"
	log "github.com/sirupsen/logrus"
)

type IPlayerCommand interface {
	CreateCommand(player *entities.Player) error
	ListCommand() ([]entities.Player, error)
	FindByIDCommand(id int) (*entities.Player, error)
	UpdateCommand(id int, player *entities.Player) error
}

type PlayerCommand struct {
	repository interfaces.IPlayerRepository
}

func NewPlayerCommand(repository interfaces.IPlayerRepository) PlayerCommand {
	return PlayerCommand{repository}
}

func (c PlayerCommand) CreateCommand(player *entities.Player) error {
	log.Info("starting create command...")
	isPlayer, err := c.repository.FindByPlayer(player)
	if err != nil {
		return fmt.Errorf("could not find playe %s %s. Here is the reason: %s", player.Name, player.LastName, err)
	}
	if isPlayer.PlayerID != 0 {
		return fmt.Errorf("player already exist")
	}
	return c.repository.Create(player)
}

func (c PlayerCommand) ListCommand() ([]entities.Player, error) {
	log.Info("starting list command...")
	return c.repository.List()
}

func (c PlayerCommand) FindByIDCommand(id int) (*entities.Player, error) {
	log.Info("starting find by id command...")
	log.Infof("--id received:  %d in FindByIDCommand", id)
	return c.repository.FindByID(id)
}

func (c PlayerCommand) UpdateCommand(id int, player *entities.Player) error {
	log.Info("starting update command...")
	isPlayer, err := c.repository.FindByID(id)
	if err != nil {
		return fmt.Errorf("could not find playe %s %s. Here is the reason: %s", player.Name, player.LastName, err)
	}
	if isPlayer.PlayerID != 0 {
		return fmt.Errorf("player already exist")
	}
	return c.repository.Update(id, player)
}
