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
	// query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // selct * from items
	query := r.db.Model(&entities.Item{}).Where("is_archive  = ?", false)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}
	//##################### การ สร้าง query #####################

	//########################## pageinate ##########################
	// offset := page-1 *limit
	//########################## pageinate ##########################
	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	size := int(itemFilter.Size)

	if err := query.Offset(offset).Limit(size).Find(&itemList).Order("id asc").Error; err != nil { //เปลี่ยนจาก r.db เป็น query ก็จะได้ filter ที่เราต้องการ
		// if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf(" Fail to Listing item :%s", err.Error())
		return nil, &_itemShopExceptions.Itemlisting{}
		// return nil, err

	}

	return itemList, nil

}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemshopModel.ItemShopFilter) (int64, error) { // Listing retrieves a list of items from the database

	//##################### การ สร้าง query #####################
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // selct * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}
	//##################### การ สร้าง query #####################
	var count int64
	if err := query.Count(&count).Error; err != nil {

		r.logger.Errorf(" Counting item fali :%s", err.Error())
		return -1, &_itemShopExceptions.ItemCounting{}

	}

	return count, nil

}
