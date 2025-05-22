package entities

import "time"

// type PurchaseHistory struct {
// 	ID              uint64 `gorm:"primaryKey;autoIncrement;"`   // Unique identifier for the purchase history entry
// 	PlayerID        string `gorm:"type:varchar(64);not null;"`  // ID of the player who made the purchase
// 	ItemID          uint64 `gorm:"type:bigint;not null;"`       // ID of the item purchased
// 	ItemName        string `gorm:"type:varchar(64);not null;"`  // Name of the item purchased
// 	ItemDescription string `gorm:"type:varchar(128);not null;"` // Description of the item purchased
// 	ItemPrice       int16  `gorm:"type:int;not null;"`          // Price of the item purchased
// 	ItemPicture     string    `gorm:"type:varchar(256);not null;"` // Picture URL of the item purchased
// 	Quantity  int16     `gorm:"type:int;not null;"`       // Quantity of the item purchased
// 	Isbuying  bool      `gorm:"boolean;not null;"`        // Indicates if the purchase is a buying transaction
// 	CreatedAt time.Time `gorm:"autoCreateTime;not null;"` // Timestamp of when the purchase was made
// }

type PurchaseHistory struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID        string    `gorm:"type:varchar(64);not null;"`
	ItemID          uint64    `gorm:"type:bigint;not null;"`
	ItemName        string    `gorm:"type:varchar(64);not null;"`
	ItemDescription string    `gorm:"type:varchar(128);not null;"`
	ItemPrice       uint      `gorm:"not null;"`
	ItemPicture     string    `gorm:"type:varchar(128);not null;"`
	Quantity        uint      `gorm:"not null;"`
	IsBuying        bool      `gorm:"type:boolean;not null;"`
	CreatedAt       time.Time `gorm:"not null;autoCreateTime;"`
}
