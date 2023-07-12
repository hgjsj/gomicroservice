package client

import (
	"go-microservice/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	InitSQLit("cloud.db")
}

func DBMigrationAll() {
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

func CreateItem(c model.CRUD) interface{} {
	return c.Create(db)
}

func GetItem(c model.CRUD) interface{} {
	return c.Get(db)
}

func ListItems(c model.CRUD, filters map[string][]string) interface{} {
	return c.List(db, filters)
}

func PatchItem(c model.CRUD, u model.CRUD) interface{} {
	return c.Patch(db, u)
}

func DeleteItem(c model.CRUD) interface{} {
	return c.Delete(db)
}
