package controller

import (
	_itemManagingService "github.com/bigstth/isekai-shop-api/pkg/itemManaging/service"
)

type ItemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(itemManagingService _itemManagingService.ItemManagingService) ItemManagingController {
	return &ItemManagingControllerImpl{itemManagingService}
}
