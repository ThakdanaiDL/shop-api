package entities

import "time"

type Inventory struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID  string    `gorm:"type:varchar(64);not null;"`
	ItemID    uint64    `gorm:"type:bigint;not null;"`
	IsDeleted bool      `gorm:"default:false;not null;"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`
}

// type Inventory struct {
// 	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
// 	PlayerID  string    `gorm:"type:varchar(64);not null;"`
// 	ItemID    uint64    `gorm:"type:bigint;not null;"`
// 	IsDeleted bool      `gorm:"not null;default:false;"`
// 	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
// }
