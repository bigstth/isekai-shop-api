package repository

import "github.com/bigstth/isekai-shop-api/entities"

type InventoryRepository interface {
	Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error)
	Removing(playerID string, itemID uint, limit int) (*entities.Inventory, error)
	PlayerItemCounting(playerID string, itemID uint64) int64
}
