package service

import (
	_inventoryRepository "github.com/bigstth/isekai-shop-api/pkg/inventory/repository"
	_itemShopRepository "github.com/bigstth/isekai-shop-api/pkg/itemShop/repository"
)

type InventoryRepositoryImpl struct {
	inventoryRepository _inventoryRepository.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
}

func NewInventoryRepositoryImpl(inventoryRepository _inventoryRepository.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{inventoryRepository, itemShopRepository}
}
