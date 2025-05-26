package repository

import "github.com/ThakdanaiDL.git/shop-api/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
