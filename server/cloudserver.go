package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-microservice/endpoint"
)

func LauchCloudServer(port int) {
	router := gin.Default()
	go func() {

		//Proceed VM router
		router.POST("/token", endpoint.MakeTokenEndpoint())
		router.POST("/vm", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeVMPostEndpoint())
		router.POST("/disk", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeDiskPostEndpoint())
		router.GET("/disk/:id", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeDiskGetEndpoint())
		router.GET("/vm/:id", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeVMGetEndpoint())
		router.GET("/vm", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeListVMEndpoint())
		router.GET("/disk", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeListDiskEndpoint())
		router.PATCH("/disk/:id", endpoint.MakeValidateTokenEndpoint(), endpoint.MakePatchDiskEndpoint())
		router.PATCH("/vm/:id", endpoint.MakeValidateTokenEndpoint(), endpoint.MakePatchVMEndpoint())
		router.DELETE("/disk/:id", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeDeleteDiskEndpoint())
		router.DELETE("/vm/:id", endpoint.MakeValidateTokenEndpoint(), endpoint.MakeDeleteVMEndpoint())
		listen := fmt.Sprintf(":%d", port)
		router.Run(listen)
	}()
}
