package repository

import (
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_itemCreatingException "github.com/bigstth/isekai-shop-api/pkg/itemManaging/exceptions"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		if custom.IsUniqueConstraintError(err, "item") {
			r.logger.Error("Item already exists:", err)
			return nil, &_itemCreatingException.ItemAlreadyExistsException{}

		}

		r.logger.Error("Failed to create item:", err)
		return nil, &_itemCreatingException.ItemCreatingException{}
	}

	return item, nil
}
func (r *itemManagingRepositoryImpl) Editing(itemId uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Model(&entities.Item{}).Where("id = ?", itemId).Updates(itemEditingReq).Error; err != nil {
		r.logger.Error("Failed to create item:", err)
		return 0, &_itemCreatingException.ItemEditingException{}
	}

	return itemId, nil
}
