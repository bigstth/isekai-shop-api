package service

import (
	"github.com/bigstth/isekai-shop-api/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository repository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepository repository.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}
