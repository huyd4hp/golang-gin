package global

import (
	"sync"

	"gorm.io/gorm"
)


var (
    dbInstance *gorm.DB
    once       sync.Once
)

func GetDB() *gorm.DB {
    return dbInstance
}

func SetDB(db *gorm.DB) {
    once.Do(func() {
		println("Set DB")
        dbInstance = db
    })
}