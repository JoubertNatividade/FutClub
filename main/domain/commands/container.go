package commands

import (
	"github.com/google/wire"
)

var Container = wire.NewSet(
	NewPlayerCommand,
	wire.Bind(new(IPlayerCommand), new(PlayerCommand)),
)
