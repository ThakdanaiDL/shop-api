package entities

import "time"

type PlayerCoin struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID  string    `gorm:"type:varchar(64);not null;"` // ID of the player
	Amount    int64     `gorm:"not null;"`                  // Amount of coins
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`   // Created at
}

// type (
// 	PlayerCoin struct {
// 		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
// 		PlayerID  string    `gorm:"type:varchar(64);not null;"`
// 		Amount    int64     `gorm:"not null;"`
// 		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
// 	}
// )
