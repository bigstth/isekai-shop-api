package controller

import "github.com/bigstth/isekai-shop-api/pkg/playerCoin/service"

type playerCoinControllerImpl struct {
	playerCoinService service.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService service.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}
