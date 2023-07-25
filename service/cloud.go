package service

import (
	"fmt"
	"go-microservice/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

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

func Create(i interface{}) error {
	res := db.Create(i)
	return res.Error
}

func Get(i interface{}) error {
	if vm, ok := i.(*model.VirtualMachine); ok {
		disks := make([]model.Disk, 0)
		res := db.First(i)
		if res.Error != nil {
			return res.Error
		}

		if err := db.Model(i).Association("Disks").Find(&disks); err != nil {
			return err
		}
		vm.Disks = disks
		return nil
	}
	res := db.First(i)
	return res.Error
}

func List(i interface{}, filters map[string][]string) (interface{}, error) {
	query := make(map[string]interface{})

	for k, v := range filters {
		query[k] = v[0]
	}

	if _, ok := i.(*model.VirtualMachine); ok {
		var vms []model.VirtualMachine
		res := db.Where(query).Find(&vms)
		if res.Error == nil {
			for index, value := range vms {
				disks := []model.Disk{}
				err := db.Model(&value).Association("Disks").Find(&disks)
				if err == nil {
					vms[index].Disks = disks
				}
			}
			return vms, nil
		} else {
			return nil, res.Error
		}
	}

	if _, ok := i.(*model.Disk); ok {
		var disks []model.Disk
		delete(query, "vmname")
		if val, ok := filters["vmname"]; ok {
			var vmids []int
			var vms []model.VirtualMachine
			res := db.Where(map[string]string{"name": val[0]}).Find(&vms)
			if res.Error != nil {
				return nil, res.Error
			}
			for _, v := range vms {
				vmids = append(vmids, int(v.ID))
			}
			db.Where(query).Where("virtual_machine_id IN ?", vmids).Find(&disks)
			return disks, res.Error

		} else {
			db.Where(query).Find(&disks)
		}
		return disks, nil

	}

	return nil, fmt.Errorf("No Object defined")
}

func Patch(i interface{}, u interface{}) error {
	res := db.Find(i)
	if res.Error != nil {
		return res.Error
	}
	res = db.Model(i).Updates(u)
	return res.Error
}

func Delete(i interface{}) error {
	res := db.Delete(i)
	return res.Error
}
