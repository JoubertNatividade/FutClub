package controllers

import "github.com/google/wire"

var Container = wire.Bind(
	NewPlayerController,
)
