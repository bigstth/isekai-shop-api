package repository

import (
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopExceptions "github.com/bigstth/isekai-shop-api/pkg/itemShop/exceptions"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {

	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Error("Error while getting item list: %s", err.Error())
		return nil, &_itemShopExceptions.ItemListing{}
	}

	return itemList, nil
}
