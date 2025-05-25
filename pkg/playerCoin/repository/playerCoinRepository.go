package repository

import (
	"github.com/bigstth/isekai-shop-api/entities"
	_playerCoinModel "github.com/bigstth/isekai-shop-api/pkg/playerCoin/model"
)

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerId string) (*_playerCoinModel.PlayerCoinShowing, error)
}
