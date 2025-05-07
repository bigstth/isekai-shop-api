package service

import (
	_itemManagingRepository "github.com/bigstth/isekai-shop-api/pkg/ItemManaging/repository"
)

type ItemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(itemManagingRepository _itemManagingRepository.ItemManagingRepository) ItemManagingService {
	return &ItemManagingServiceImpl{itemManagingRepository: itemManagingRepository}
}
