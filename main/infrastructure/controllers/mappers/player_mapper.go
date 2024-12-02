package mappers

import (
	"github.com/JoubertNatividade/FutClub/main/domain/entities"
	"github.com/JoubertNatividade/FutClub/main/infrastructure/controllers/requests"
)

func MapToEntityPlayer(request requests.PlayerRequest) entities.Player {
	return entities.Player{
		Name:     request.Name,
		Position: request.Position,
		Goals:    request.Goals,
	}
}
