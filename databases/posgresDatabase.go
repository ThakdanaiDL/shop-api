package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/ThakdanaiDL.git/shop-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type posgresDatabase struct {
	// DB *gorm.DB // embedded → Go ตั้งชื่อให้ว่า DB  เเต่ไม่เเสดง
	*gorm.DB // embedded → Go ตั้งชื่อให้ว่า DB  เเต่ไม่เเสดง
}

var (
	posgresDatabaseInstance *posgresDatabase
	once                    sync.Once
)

func NewPosgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			// "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			"host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s search_path=%s",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.SSLMode,
			conf.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		log.Printf("Connect to database %s", conf.DBName)

		// posgresDatabaseInstance = &posgresDatabase{DB: conn} // DB คือเเม่เเบบ struct  คล้ายๆ key : value คือ ค่าที่ user ใส่มาที่นี้คือ conn จึงเป็นการ ซ่อน struct เวลาเรียกใช้ เรียกเเค่ New...เพราะสนเเค่ value เนื่องจากเราสร้าง key รอไว้อยู่เเล้ว
		posgresDatabaseInstance = &posgresDatabase{conn} // DB คือเเม่เเบบ struct  คล้ายๆ key : value คือ ค่าที่ user ใส่มาที่นี้คือ conn จึงเป็นการ ซ่อน struct เวลาเรียกใช้ เรียกเเค่ New...เพราะสนเเค่ value เนื่องจากเราสร้าง key รอไว้อยู่เเล้ว

	})
	return posgresDatabaseInstance // ที่ต้อง เอาตัวเเปรมารับเพราะ เก็บค่าไว้เเค่ครั้งเดียวเเล้วรีเทินตลอด
	// return &posgresDatabase{DB: conn} ใช้เเบบนี้ได้ ถ้าไม่ใช้ singleton
}

func (db *posgresDatabase) Connect() *gorm.DB { // ประกาศเพื่อให้ ใช้ method ของ interface ได้ return ตามtyep ของ method ที่ประกาศไว้

	return db.DB // เรียก ใช้  db คือชื่อตัวเเปร ส่วน DB ใหญ่ คือ ชื่อของ สมาชิกใน struct ที่ซ่อนอยู่

}
