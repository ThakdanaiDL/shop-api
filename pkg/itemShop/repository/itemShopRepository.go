package repository

import (
	"github.com/ThakdanaiDL.git/shop-api/entities"
	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemshopModel.ItemShopFilter) ([]*entities.Item, error)
}
