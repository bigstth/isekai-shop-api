package repository

import (
	"github.com/bigstth/isekai-shop-api/databases"
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/bigstth/isekai-shop-api/pkg/custom"
	_itemManagingException "github.com/bigstth/isekai-shop-api/pkg/itemManaging/exceptions"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)

	if err := r.db.Connect().Create(itemEntity).Scan(item).Error; err != nil {
		if custom.IsUniqueConstraintError(err, "item") {
			r.logger.Error("Item already exists:", err)
			return nil, &_itemManagingException.ItemAlreadyExistsException{}

		}

		r.logger.Error("Failed to create item:", err)
		return nil, &_itemManagingException.ItemCreatingException{}
	}

	return item, nil
}
func (r *itemManagingRepositoryImpl) Editing(itemId uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemId).Updates(itemEditingReq).Error; err != nil {
		r.logger.Error("Failed to create item:", err)
		return 0, &_itemManagingException.ItemEditingException{}
	}

	return itemId, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemId uint64) error {

	item := new(entities.Item)
	if err := r.db.Connect().Table("items").Where("id = ?", itemId).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Warn("Item not found, ID:", itemId)
			return &_itemManagingException.ItemNotFoundException{}
		}
		r.logger.Error("Failed to query item:", err)
		return &_itemManagingException.ItemNotFoundException{}
	}

	// Check if the item is already archived
	if item.IsArchive {
		r.logger.Warn("Item already archived, ID:", itemId)
		return &_itemManagingException.ItemNotFoundException{}
	}

	result := r.db.Connect().Table("items").Where("id = ?", itemId).Update("is_archive", true)

	if result.RowsAffected == 0 {
		r.logger.Error("Item not found or already archived")
		return &_itemManagingException.ItemNotFoundException{}
	}

	if result.Error != nil {
		r.logger.Error("Failed to archive item:", result.Error)
		return &_itemManagingException.ItemArchivingException{}
	}

	return nil
}
