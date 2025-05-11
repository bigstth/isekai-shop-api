package service

import (
	"github.com/bigstth/isekai-shop-api/entities"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingRepository "github.com/bigstth/isekai-shop-api/pkg/itemManaging/repository"
	_itemShopModel "github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/bigstth/isekai-shop-api/pkg/itemShop/repository"
)

type ItemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository, itemShopRepository _itemShopRepository.ItemShopRepository) ItemManagingService {
	return &ItemManagingServiceImpl{itemManagingRepository, itemShopRepository}
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

func (s *ItemManagingServiceImpl) Editing(itemId uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {

	itemIdResult, err := s.itemManagingRepository.Editing(itemId, itemEditingReq)
	if err != nil {
		return nil, err
	}

	itemEntityResult, err := s.itemShopRepository.FindByID(itemIdResult)

	if err != nil {
		return nil, err

	}
	return itemEntityResult.ToItemModel(), nil
}
