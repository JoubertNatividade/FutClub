package repositories

import (
	"github.com/JoubertNatividade/FutClub/main/domain/interfaces"
	"github.com/google/wire"
)

var Container = wire.NewSet(
	NewPlayerRepository,
	wire.Bind(new(interfaces.IPlayerRepository), new(*PlayerRepository)),
)
