package client

import (
	"go-microservice/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBMigrationAll() {
	if db == nil {
		panic("Connect database before migration")
	}
	//migrate DB schema as need
	db.AutoMigrate(&model.Disk{}, &model.VirtualMachine{})
}

func InitSQLit(path string) {
	sqlite, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = sqlite
}
