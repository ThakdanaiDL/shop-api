package entities

import "time"

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

// type Item struct {
// 	ID          uint64    `gorm:"primaryKey;autoIncrement"`
// 	AdminID     *string   `gorm:"type:varchar(64);"`
// 	Name        string    `gorm:"type:varchar(64);unique;not null;"`
// 	Description string    `gorm:"type:varchar(128);not null;"`
// 	Picture     string    `gorm:"type:varchar(256);not null;"`
// 	Price       uint      `gorm:"not null;"`
// 	IsArchive   bool      `gorm:"not null;default:false;"`
// 	CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
// 	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
// }
