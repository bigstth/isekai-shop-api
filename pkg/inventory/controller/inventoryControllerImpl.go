package controller

import (
	_inventoryService "github.com/bigstth/isekai-shop-api/pkg/inventory/service"
)

type InventoryControllerImpl struct {
	inventoryService _inventoryService.InventoryService
}

func NewInventoryControllerImpl(inventoryService _inventoryService.InventoryService) *InventoryControllerImpl {
	return &InventoryControllerImpl{inventoryService}
}
