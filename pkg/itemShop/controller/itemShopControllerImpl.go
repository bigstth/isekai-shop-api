package controller

import (
	"net/http"

	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_itemShopModel "github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService service.ItemShopService
}

func NewItemShopControllerImpl(itemShopService service.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(ctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)
	customEchoRequest := custom.NewCustomEchoRequest(ctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(ctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)

	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, itemModelList)
}
