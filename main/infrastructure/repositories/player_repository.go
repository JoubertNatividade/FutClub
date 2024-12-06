package repositories

import "github.com/JoubertNatividade/FutClub/main/domain/entities"

type PlayerRepository struct{}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

func (r *PlayerRepository) Create(player *entities.Player) error {

	return nil
}
