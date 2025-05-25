package repository

import (
	"github.com/bigstth/isekai-shop-api/databases"
	"github.com/bigstth/isekai-shop-api/entities"
	_playerCoinException "github.com/bigstth/isekai-shop-api/pkg/playerCoin/exception"
	_playerCoinModel "github.com/bigstth/isekai-shop-api/pkg/playerCoin/model"
	"github.com/labstack/echo/v4"
)

type playerCoinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerCoinRepositoryImpl(db databases.Database, logger echo.Logger) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{db, logger}
}

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	playerCoin := new(entities.PlayerCoin)
	if err := r.db.Connect().Create(playerCoinEntity).Scan(playerCoin).Error; err != nil {
		r.logger.Error("Failed to add coin: %s", err.Error())
		return nil, &_playerCoinException.CoinAddingException{}
	}
	return playerCoin, nil
}

func (r *playerCoinRepositoryImpl) Showing(playerId string) (*_playerCoinModel.PlayerCoinShowing, error) {
	playerCoin := new(_playerCoinModel.PlayerCoinShowing)

	if err := r.db.Connect().Table("player_coins").Select("player_id, SUM(amount) as coin").
		Where("player_id = ?", playerId).Group("player_id").Scan(playerCoin).Error; err != nil {
		r.logger.Error("Failed to show player coin: %s", err.Error())
		return nil, &_playerCoinException.PlayerCoinShowingException{}
	}
	return playerCoin, nil
}
