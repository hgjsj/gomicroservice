package model

import (
	"gorm.io/gorm"
)

type state int

const (
	Requested    state = 0
	Active       state = 1
	Building     state = 2
	Cancelled    state = 3
	Inconsistent state = 4
)

type Base struct {
	Name   string `gorm:"index" json:"name"`
	Status state  `gorm:"default:0" json:"state"`
}

type VirtualMachine struct {
	gorm.Model
	Base
	CPU    uint32 `json:"cpu"`
	Memory uint32 `json:"memory"`
	Disks  []Disk `json:"disk,omitempty"`
}

type Disk struct {
	gorm.Model
	Base
	Size             uint64 `json:"size"`
	VirtualMachineID uint   `json:"vm,omitempty"`
}

type DiskList struct {
	Disks []Disk `json:"disk"`
}

type Tabler interface {
	TableName() string
}

func (VirtualMachine) TableName() string {
	return "vm"
}

func (Disk) TableName() string {
	return "disk"
}

type CRUD interface {
	Create(*gorm.DB) CRUD
	Get(*gorm.DB) CRUD
	List(*gorm.DB, map[string][]string) interface{}
	Patch(*gorm.DB, CRUD) CRUD
	Delete(*gorm.DB) CRUD
}

func (vm VirtualMachine) Create(db *gorm.DB) CRUD {
	res := db.Create(&vm)
	if res.Error != nil {

	}
	return vm
}

func (d Disk) Create(db *gorm.DB) CRUD {
	res := db.Create(&d)
	if res.Error != nil {

	}
	return d
}

func (vm VirtualMachine) Get(db *gorm.DB) CRUD {
	disks := []Disk{}
	res := db.First(&vm)
	if res.Error != nil {
		return nil
	}

	if err := db.Model(&vm).Association("Disks").Find(&disks); err == nil {
		vm.Disks = disks
	}

	return vm
}

func (d Disk) Get(db *gorm.DB) CRUD {
	res := db.First(&d)
	if res.Error != nil {
		return nil
	}
	return d
}

func (vm VirtualMachine) List(db *gorm.DB, filters map[string][]string) interface{} {
	var vms []VirtualMachine
	query := make(map[string]interface{})

	for k, v := range filters {
		query[k] = v[0]
	}

	res := db.Where(query).Find(&vms)
	if res.Error == nil {
		for index, value := range vms {
			disks := []Disk{}
			res := db.Model(&value).Association("Disks").Find(&disks)
			if res == nil {
				vms[index].Disks = disks
			}
		}
		return vms
	}
	return nil
}

func (d Disk) List(db *gorm.DB, filters map[string][]string) interface{} {
	var disks []Disk
	query := make(map[string]interface{})

	for k, v := range filters {
		if k != "vmname" {
			query[k] = v[0]
		}

	}

	if val, ok := filters["vmname"]; ok {
		var vmids []int
		var vms []VirtualMachine
		db.Where(map[string]string{"name": val[0]}).Find(&vms)
		for _, v := range vms {
			vmids = append(vmids, int(v.ID))
		}
		db.Where(query).Where("virtual_machine_id IN ?", vmids).Find(&disks)
	} else {
		db.Where(query).Find(&disks)
	}

	return disks
}

func (d Disk) Patch(db *gorm.DB, u CRUD) CRUD {
	res := db.Find(&d)
	if res.Error != nil {
		return nil
	}
	db.Model(&d).Updates(u)
	return d
}

func (vm VirtualMachine) Patch(db *gorm.DB, u CRUD) CRUD {
	res := db.Find(&vm)
	if res.Error != nil {
		return nil
	}
	db.Model(&vm).Updates(u)
	return vm
}

func (vm VirtualMachine) Delete(db *gorm.DB) CRUD {
	res := db.Delete(&vm)
	if res.Error != nil {

	}
	return vm
}

func (d Disk) Delete(db *gorm.DB) CRUD {
	res := db.Delete(&d)
	if res.Error != nil {

	}
	return d
}
