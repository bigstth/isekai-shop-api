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

	//pagination is (page-1)* size
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(limit).Order("id DESC").Find(&itemList).Error; err != nil {
		r.logger.Error("Error while getting item list: %s", err.Error())
		return nil, &_itemShopExceptions.ItemListing{}
	}

	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *model.ItemFilter) (int64, error) {

	//pagination is (page-1)* size
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilter.Description+"%")
	}

	count := new(int64)

	if err := query.Count(count).Error; err != nil {
		r.logger.Error("Failed to get counting: %s", err.Error())
		return -1, &_itemShopExceptions.ItemCounting{}
	}

	return *count, nil
}
