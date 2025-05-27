package model

type (
	Item struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
		Price       uint   `json:"price"`

		// AdminID     *string   `gorm:"type:varchar(64);"`        // ID of the admin who created the item client ไม่จำเป็นต้องรู้
		// IsAchived   bool      `gorm:"default:false;not null;"`  // ID of the admin who created the item client ไม่จำเป็นต้องรู้
		// CreatedAt   time.Time `gorm:"autoCreateTime;not null;"` // ID of the admin who created the item client ไม่จำเป็นต้องรู้
		// UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;"` // ID of the admin who created the item client ไม่จำเป็นต้องรู้
	}

	ItemShopFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
	}
)
