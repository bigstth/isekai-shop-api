package service

import (
	_playerCoinModel "github.com/bigstth/isekai-shop-api/pkg/playerCoin/model"
)

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error)
	Showing(playerId string) *_playerCoinModel.PlayerCoinShowing
}
