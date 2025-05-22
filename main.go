package main

import (
	"github.com/ThakdanaiDL.git/shop-api/config"
	database "github.com/ThakdanaiDL.git/shop-api/databases"
	"github.com/ThakdanaiDL.git/shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPosgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.Connect())

	server.Start()

	// server.Start() // เรียกใช้ method Start ของ echoServer

}
