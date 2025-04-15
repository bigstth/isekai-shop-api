package repository

import "github.com/bigstth/isekai-shop-api/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
