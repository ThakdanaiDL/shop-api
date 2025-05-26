package entities

import (
	"time"

	_itemshopModel "github.com/ThakdanaiDL.git/shop-api/pkg/itemShop/model"
)

type Item struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;"`         // ID of the item
	AdminID     *string   `gorm:"type:varchar(64);"`                 // ID of the admin who created the item
	Name        string    `gorm:"type:varchar(64);unique;not null;"` // ID of the admin who created the item
	Description string    `gorm:"type:varchar(128);not null;"`       // ID of the admin who created the item
	Picture     string    `gorm:"type:varchar(256);not null;"`       // ID of the admin who created the item
	Price       uint      `gorm:"not null;"`
	IsAchived   bool      `gorm:"default:false;not null;"`  // ID of the admin who created the item
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;"` // ID of the admin who created the item
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;"` // ID of the admin who created the item
}

func (i *Item) ToItemShopModel() *_itemshopModel.Item {
	return &_itemshopModel.Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Picture:     i.Picture,
	}

}
