package service

import (
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
)

type ItemManagingService interface {
	Creating(itemCreatingRequest *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error)
	Editing(itemId uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error)
}
