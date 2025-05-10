package service

import (
	"github.com/bigstth/isekai-shop-api/entities"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingRepository "github.com/bigstth/isekai-shop-api/pkg/itemManaging/repository"
	_itemShopModel "github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
)

type ItemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository) ItemManagingService {
	return &ItemManagingServiceImpl{itemManagingRepository: itemManagingRepository}
}

func (s *ItemManagingServiceImpl) Creating(itemCreatingRequest *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingRequest.Name,
		Description: itemCreatingRequest.Description,
		Price:       itemCreatingRequest.Price,
		Picture:     itemCreatingRequest.Picture,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil

}
