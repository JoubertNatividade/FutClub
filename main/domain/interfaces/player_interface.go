package interfaces

import "github.com/JoubertNatividade/FutClub/main/domain/entities"

type PlayerInterface interface {
	Create(player *entities.Player) error
}
