package server

import (
	_playerCoinController "github.com/bigstth/isekai-shop-api/pkg/playerCoin/controller"
	_playerCoinRepository "github.com/bigstth/isekai-shop-api/pkg/playerCoin/repository"
	_playerCoinService "github.com/bigstth/isekai-shop-api/pkg/playerCoin/service"
)

func (s *echoServer) initPlayerCoinRouter(m *authorizingMiddleware) {
	router := s.app.Group("/v1/player-coin")

	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db, s.app.Logger)
	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(playerCoinRepository)
	playerCoinController := _playerCoinController.NewPlayerCoinControllerImpl(playerCoinService)

	router.POST("/add", playerCoinController.CoinAdding, m.PlayerAuthorizing)
}
