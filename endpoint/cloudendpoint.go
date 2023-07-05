package endpoint

import (
	"context"
	"go-microservice/client"
	"go-microservice/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
)

func MakeVMPostEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		vm := request.(model.VirtualMachine)
		re_vm := client.CreateItem(vm)
		return re_vm, nil
	}
}

func MakeDiskEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		disk := request.(model.Disk)
		re_disk := client.CreateItem(disk)
		return re_disk, nil
	}
}

func MakeVMGetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		vm := model.VirtualMachine{}
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		vm.ID = uint(i)
		re_vm := client.GetItem(vm)
		c.JSON(200, re_vm)
	}
}

func MakeDiskGetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var disk model.Disk
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		disk.ID = uint(i)
		re_disk := client.GetItem(disk)
		c.JSON(200, re_disk)
	}
}

func MakeListVMEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		re_vm := client.ListItems(model.VirtualMachine{})
		return re_vm, nil
	}
}

func MakeListDiskEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		re_disk := client.ListItems(model.Disk{})
		return re_disk, nil
	}
}

func MakePatchDiskEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var disk, u model.Disk

		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		disk.ID = uint(i)

		if c.BindJSON(&u) == nil {
			c.JSON(200, client.PatchItem(disk, u))
		}

		c.JSON(304, nil)
	}
}

func MakePatchVMEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm, u model.VirtualMachine
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		vm.ID = uint(i)

		if c.BindJSON(&u) == nil {
			c.JSON(200, client.PatchItem(vm, u))
		}

		c.JSON(304, nil)
	}
}

func MakeDeleteDiskEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var disk model.Disk
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		disk.ID = uint(i)
		c.JSON(200, client.DeleteItem(disk))
	}
}

func MakeDeleteVMEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm model.VirtualMachine
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		vm.ID = uint(i)
		c.JSON(200, client.DeleteItem(vm))
	}
}
