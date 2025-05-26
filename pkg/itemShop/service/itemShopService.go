package services

import (
	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing() ([]*_itemshopModel.Item, error) // layer service จะใช้ model ของ itemShopเเทน entities
}
