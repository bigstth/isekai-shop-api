package repository

import "github.com/bigstth/isekai-shop-api/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindById(playerId string) (*entities.Player, error)
}
