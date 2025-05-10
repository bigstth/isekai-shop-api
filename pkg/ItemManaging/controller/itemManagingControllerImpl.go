package controller

import (
	"net/http"

	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/bigstth/isekai-shop-api/pkg/itemManaging/service"
	"github.com/labstack/echo/v4"
)

type ItemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &ItemManagingControllerImpl{itemManagingService}
}

func (c *ItemManagingControllerImpl) Creating(ctx echo.Context) error {
	itemCreatingRequest := new(_itemManagingModel.ItemCreatingReq)

	customEchoRequest := custom.NewCustomEchoRequest(ctx)

	if err := customEchoRequest.Bind(itemCreatingRequest); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	item, err := c.itemManagingService.Creating(itemCreatingRequest)

	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, item)
}
