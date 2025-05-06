package service

import (
	"github.com/bigstth/isekai-shop-api/entities"
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/model"
	"github.com/bigstth/isekai-shop-api/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing(itemFilter *model.ItemFilter) (model.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return model.ItemResult{}, err
	}
	itemCounting, err := s.itemShopRepository.Counting(itemFilter)
	if err != nil {
		return model.ItemResult{}, err
	}

	totalPage := s.totalPageCalculation(itemCounting, itemFilter.Size)

	return *s.toItemResultResponse(itemList, itemFilter.Page, totalPage), nil
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItem int64, size int64) int64 {
	totalPage := totalItem / size
	if totalItem%size != 0 {
		totalPage++
	}
	if totalPage == 0 {
		totalPage = 1
	}
	return totalPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *model.ItemResult {
	itemModelList := make([]*model.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &model.ItemResult{
		Items: itemModelList,
		Paginate: model.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}
