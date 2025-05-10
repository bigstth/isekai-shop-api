package repository

import "github.com/bigstth/isekai-shop-api/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
