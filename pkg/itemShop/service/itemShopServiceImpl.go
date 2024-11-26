package service

import "github.com/bigstth/isekai-shop-api/pkg/itemShop/repository"

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}
