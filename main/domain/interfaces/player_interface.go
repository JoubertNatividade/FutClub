package interfaces

import "github.com/JoubertNatividade/FutClub/main/domain/entities"

type IPlayerRepository interface {
	Create(player *entities.Player) error
	List() ([]entities.Player, error)
	FindByID(id int) (*entities.Player, error)
	FindByPlayer(player *entities.Player) (*entities.Player, error)
	Update(id int, player *entities.Player) error
	Delete(id int) error
}
