package commands

type PlayerInterface interface {
	Create(user *PlayerCommand) error
}
