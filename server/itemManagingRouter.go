package server

import (
	_itemManagingController "github.com/bigstth/isekai-shop-api/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/bigstth/isekai-shop-api/pkg/itemManaging/repository"
	_itemManagingService "github.com/bigstth/isekai-shop-api/pkg/itemManaging/service"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemShopRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemManagingService.NewItemManagingServiceImpl(itemShopRepository)
	itemShopController := _itemManagingController.NewItemManagingControllerImpl(itemShopService)

	_ = itemShopController
	_ = router
}
