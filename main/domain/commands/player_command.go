package commands

import (
	"github.com/JoubertNatividade/FutClub/main/domain/entities"
	"github.com/JoubertNatividade/FutClub/main/domain/interfaces"
	log "github.com/sirupsen/logrus"
)

type IPlayerCommand interface {
	CreateCommand(user *entities.Player) error
	ListCommand() ([]entities.Player, error)
	FindByIDCommand(id int) (*entities.Player, error)
}

type PlayerCommand struct {
	repository interfaces.IPlayerRepository
}

func NewPlayerCommand(repository interfaces.IPlayerRepository) PlayerCommand {
	return PlayerCommand{repository}
}

func (c PlayerCommand) CreateCommand(user *entities.Player) error {
	log.Info("starting create command...")
	return c.repository.Create(user)
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
