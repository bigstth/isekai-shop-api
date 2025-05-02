package repository

import (
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *model.ItemFilter) ([]*entities.Item, error)
}
