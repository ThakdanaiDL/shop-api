package services

import (
	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemshopModel.ItemShopFilter) (*_itemshopModel.ItemResult, error) // layer service จะใช้ model ของ itemShopเเทน entities
}
