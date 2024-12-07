package repositories

import (
	"github.com/JoubertNatividade/FutClub/main/domain/entities"
	log "github.com/sirupsen/logrus"
)

type PlayerRepository struct{}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

func (r *PlayerRepository) Create(player *entities.Player) error {
	log.Info("starting create repository...")
	log.Info("player criado com sucesso!")

	return nil
}
