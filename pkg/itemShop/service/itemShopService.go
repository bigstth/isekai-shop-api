package service

import "github.com/bigstth/isekai-shop-api/pkg/itemShop/model"

type ItemShopService interface {
	Listing() ([]*model.Item, error)
}
