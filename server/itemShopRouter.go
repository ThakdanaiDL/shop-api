package server

import (
	_itemShopController "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/controller"
	_itemShoprepository "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/repository"
	_itemShopService "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShoprepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopRepositoryImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)
	router.GET("", itemShopController.Listing)

}
