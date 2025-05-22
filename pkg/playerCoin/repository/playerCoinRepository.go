package repository

import "github.com/bigstth/isekai-shop-api/entities"

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
}
