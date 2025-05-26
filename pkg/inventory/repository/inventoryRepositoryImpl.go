package repository

import (
	"github.com/bigstth/isekai-shop-api/databases"
	"github.com/bigstth/isekai-shop-api/entities"
	_inventoryExceptions "github.com/bigstth/isekai-shop-api/pkg/inventory/exception"
	"github.com/labstack/echo/v4"
)

type InventoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewInventoryRepositoryImpl(db databases.Database, logger echo.Logger) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{db, logger}
}

func (r *InventoryRepositoryImpl) Filling(inventoryEntities []*entities.Inventory) ([]*entities.Inventory, error) {

	inventoryEntitiesResult := make([]*entities.Inventory, 0)

	if err := r.db.Connect().CreateInBatches(inventoryEntities, len(inventoryEntities)).Scan(&inventoryEntitiesResult).Error; err != nil {
		r.logger.Error("Failed to fill inventory: %s", err.Error())
		return nil, &_inventoryExceptions.InventoryFillingException{
			PlayerID: inventoryEntities[0].PlayerID,
			ItemID:   inventoryEntities[0].ItemID,
		}
	}

	return inventoryEntitiesResult, nil
}

func (r *InventoryRepositoryImpl) Removing(playerID string, itemID uint64, limit int) error {
	inventoryEntities, err := r.findItemInInventoryByPlayerID(playerID, itemID, limit)

	if err != nil {
		r.logger.Error("Failed to find item in inventory: %s", err.Error())
		return err
	}

	tx := r.db.Connect().Begin()

	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true

		if err := tx.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			tx.Rollback()
			r.logger.Error("Failed to remove item from inventory: %s", err.Error())
			return &_inventoryExceptions.PlayerItemRemoving{
				ItemID: itemID,
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		r.logger.Error("Failed to commit transaction: %s", err.Error())
		return &_inventoryExceptions.PlayerItemRemoving{
			ItemID: itemID,
		}
	}

	return nil

}

func (r *InventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) int64 {
	var count int64
	if err := r.db.Connect().Model(&entities.Inventory{}).Where("player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false).
		Count(&count).Error; err != nil {
		r.logger.Error("Failed to count player items: %s", err.Error())
		return 0
	}
	return count
}

func (r *InventoryRepositoryImpl) findItemInInventoryByPlayerID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)

	if err := r.db.Connect().Model(&entities.Inventory{}).Where("player_id = ? AND item_id =? and is_deleted = ?", playerID, itemID, false).
		Limit(limit).Find(inventoryEntities).Error; err != nil {
		r.logger.Error("Failed to find item in inventory: %s", err.Error())
		return nil, &_inventoryExceptions.PlayerItemsFinding{
			PlayerID: playerID,
		}
	}

	return inventoryEntities, nil
}
