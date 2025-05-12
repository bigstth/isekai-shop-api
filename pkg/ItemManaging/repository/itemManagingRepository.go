package repository

import (
	"github.com/bigstth/isekai-shop-api/entities"
	_itemManagingModel "github.com/bigstth/isekai-shop-api/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemId uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
	Archiving(itemId uint64) error
}
