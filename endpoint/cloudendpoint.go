package endpoint

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"go-microservice/model"
	"go-microservice/service"
	"net/http"
	"strconv"
)

func MakeVMPostEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		vm := request.(model.VirtualMachine)
		error := service.Create(&vm)
		return vm, error
	}
}

func MakeDiskEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		disk := request.(model.Disk)
		error := service.Create(&disk)
		return disk, error
	}
}

func MakeVMGetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		vm := model.VirtualMachine{}
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		vm.ID = uint(i)
		if err := service.Get(&vm); err == nil {
			c.JSON(http.StatusOK, vm)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func MakeDiskGetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var disk model.Disk
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		disk.ID = uint(i)
		if err := service.Get(&disk); err == nil {
			c.JSON(http.StatusOK, disk)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

	}
}

func MakeListVMEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		vms, err := service.List(&model.VirtualMachine{}, c.Request.URL.Query())
		if err == nil {
			c.JSON(http.StatusOK, vms)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

	}
}

func MakeListDiskEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		disks, err := service.List(&model.Disk{}, c.Request.URL.Query())
		if err == nil {
			c.JSON(http.StatusOK, disks)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

	}
}

func MakePatchDiskEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var disk, u model.Disk

		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		disk.ID = uint(i)

		if err := c.BindJSON(&u); err == nil {
			if err = service.Patch(&disk, u); err == nil {
				c.JSON(http.StatusOK, disk)

			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func MakePatchVMEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm, u model.VirtualMachine
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		vm.ID = uint(i)

		if err := c.BindJSON(&u); err == nil {
			if err = service.Patch(&vm, u); err == nil {
				c.JSON(http.StatusOK, vm)

			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func MakeDeleteDiskEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var disk model.Disk
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		disk.ID = uint(i)
		if err := service.Delete(&disk); err == nil {
			c.JSON(http.StatusOK, disk)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func MakeDeleteVMEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vm model.VirtualMachine
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 32)
		vm.ID = uint(i)
		if err := service.Delete(&vm); err == nil {
			c.JSON(http.StatusOK, vm)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

	}
}

func MakeTokenEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := service.NewToken(); err == nil {
			c.Header("X-Subject-Token", token)
			c.JSON(http.StatusOK, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}

	}
}

func MakeValidateTokenEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Cloud-Token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": "Need token in header"})
		} else {
			if isValid, err := service.ValidateToken(token); !isValid {
				c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}
		}

	}
}
