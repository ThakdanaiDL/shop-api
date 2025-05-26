package repository

import (
	"github.com/ThakdanaiDL.git/shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	logger echo.Logger
	db     *gorm.DB
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db: db, logger: logger}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) { // Listing retrieves a list of items from the database
	itemList := make([]*entities.Item, 0) // Create a slice to hold the items

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf(" Fail to Listing item :%s", err.Error())
		return nil, err

	}

	return itemList, nil

}
