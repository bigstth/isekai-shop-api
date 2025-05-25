package controller

import (
	"net/http"

	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_playerCoinModel "github.com/bigstth/isekai-shop-api/pkg/playerCoin/model"
	"github.com/bigstth/isekai-shop-api/pkg/playerCoin/service"
	"github.com/bigstth/isekai-shop-api/pkg/validation"
	"github.com/labstack/echo/v4"
)

type playerCoinControllerImpl struct {
	playerCoinService service.PlayerCoinService
}

func NewPlayerCoinControllerImpl(playerCoinService service.PlayerCoinService) PlayerCoinController {
	return &playerCoinControllerImpl{playerCoinService}
}

func (c *playerCoinControllerImpl) CoinAdding(ctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	coinAddingReq := new(_playerCoinModel.CoinAddingReq)

	customEchoRequest := custom.NewCustomEchoRequest(ctx)
	if err := customEchoRequest.Bind(coinAddingReq); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	coinAddingReq.PlayerID = playerID
	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)

	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, playerCoin)
}

func (c *playerCoinControllerImpl) Showing(ctx echo.Context) error {
	playerID, err := validation.PlayerIDGetting(ctx)
	if err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	playerCoinShowing := c.playerCoinService.Showing(playerID)

	return ctx.JSON(http.StatusOK, playerCoinShowing)
}
