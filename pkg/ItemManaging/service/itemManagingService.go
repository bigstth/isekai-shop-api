package service

import (
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingRequest *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
}
