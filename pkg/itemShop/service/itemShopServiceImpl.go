package service

import (
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing() ([]*model.Item, error) {
	itemList, err := s.itemShopRepository.Listing()
	if err != nil {
		return nil, err
	}

	itemModelList := make([]*model.Item, 0)
	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}
	return itemModelList, nil
}
