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
