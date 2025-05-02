package repository

import (
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopExceptions "github.com/bigstth/isekai-shop-api/pkg/itemShop/exceptions"
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {

	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *model.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{})

	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	if err := query.Find(&itemList).Error; err != nil {
		r.logger.Error("Error while getting item list: %s", err.Error())
		return nil, &_itemShopExceptions.ItemListing{}
	}

	return itemList, nil
}
