package initialize

import (
	"basic/global"
	models "basic/internal/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func InitMySQL() {
	MySQL_Connect_String := "root:rootMYSQL@tcp(127.0.0.1:3307)/User"
	db, err := gorm.Open(mysql.Open(MySQL_Connect_String), &gorm.Config{})
	if err != nil {
	    log.Fatalln(err)
	}
	// Tạo các bảng
	db.AutoMigrate(&models.Auth{})
	global.SetDB(db)
}