package repository

import (
	"github.com/ThakdanaiDL.git/shop-api/entities"
	_itemShopExceptions "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/exception"
	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
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

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemshopModel.ItemShopFilter) ([]*entities.Item, error) { // Listing retrieves a list of items from the database
	itemList := make([]*entities.Item, 0) // Create a slice to hold the items

	//##################### การ สร้าง query #####################
	query := r.db.Model(&entities.Item{}) // selct * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}
	//##################### การ สร้าง query #####################

	if err := query.Find(&itemList).Error; err != nil { //เปลี่ยนจาก r.db เป็น query ก็จะได้ filter ที่เราต้องการ
		// if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf(" Fail to Listing item :%s", err.Error())
		return nil, &_itemShopExceptions.Itemlisting{}
		// return nil, err

	}

	return itemList, nil

}
