package controller

import (
	"net/http"

	"github.com/bigstth/isekai-shop-api/pkg/custom"
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
	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return custom.Error(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, itemModelList)
}
