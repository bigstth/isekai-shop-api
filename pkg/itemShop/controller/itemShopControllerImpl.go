package controller

import "github.com/bigstth/isekai-shop-api/pkg/itemShop/service"

type itemShopControllerImpl struct {
	itemShopService service.ItemShopService
}

func NewItemShopControllerImpl(itemShopService service.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}
