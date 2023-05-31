package model

import (
	"gorm.io/gorm"
)

type VirtualMachine struct {
	gorm.Model
	Name   string `gorm:"index"`
	CPU    uint32
	Memory uint32
	Disks  []Disk `gorm:"foreignKey:vm_id"`
}

type Disk struct {
	gorm.Model
	Size             uint64
	Name             string
	Type             uint8
	VirtualMachineID uint `gorm:"column:vm_id"`
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
