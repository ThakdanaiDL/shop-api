package main

import (
	"github.com/ThakdanaiDL.git/shop-api/config"
	database "github.com/ThakdanaiDL.git/shop-api/databases"
	"github.com/ThakdanaiDL.git/shop-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()

	db := database.NewPosgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoriesMigration(tx)
	purchaseHistoryMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

}

func playerMigration(tx *gorm.DB) {

	tx.Migrator().CreateTable(entities.Player{})
}

func adminMigration(tx *gorm.DB) {

	tx.Migrator().CreateTable(entities.Admin{})

}

func inventoriesMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(entities.Inventory{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(entities.Item{})

}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(entities.PlayerCoin{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable((entities.PurchaseHistory{}))
}
