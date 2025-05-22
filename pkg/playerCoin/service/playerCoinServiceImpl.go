package service

import (
	"github.com/bigstth/isekai-shop-api/entities"
	_playerCoinModel "github.com/bigstth/isekai-shop-api/pkg/playerCoin/model"
	"github.com/bigstth/isekai-shop-api/pkg/playerCoin/repository"
)

type playerCoinServiceImpl struct {
	playerCoinRepository repository.PlayerCoinRepository
}

func NewPlayerCoinServiceImpl(playerCoinRepository repository.PlayerCoinRepository) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *_playerCoinModel.CoinAddingReq) (*_playerCoinModel.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}

	playerCoinEntityResult, err := s.playerCoinRepository.CoinAdding(playerCoinEntity)
	if err != nil {
		return nil, err
	}

	playerCoinEntityResult.PlayerID = coinAddingReq.PlayerID

	return playerCoinEntityResult.ToPlayerCoinModel(), nil
}
