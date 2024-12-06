package interfaces

import "github.com/JoubertNatividade/FutClub/main/domain/entities"

type IPlayerRepository interface {
	Create(player *entities.Player) error
}
