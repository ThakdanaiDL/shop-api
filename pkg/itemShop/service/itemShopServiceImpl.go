package services

import (
	_itemShopRepository "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/repository"
)

type itemShopServiceImp struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopRepositoryImpl(itemShopRepository _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImp{itemShopRepository}
}
