package server

import (
	"fmt"
	"go-microservice/endpoint"
	"go-microservice/transport"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func LauchCloudServer(port int) {
	vmHandler := httptransport.NewServer(
		endpoint.MakeVMPostEndpoint(),
		transport.DecodeVMRequest,
		httptransport.EncodeJSONResponse,
	)
	diskHandler := httptransport.NewServer(
		endpoint.MakeDiskEndpoint(),
		transport.DecodeDiskRequest,
		httptransport.EncodeJSONResponse,
	)

	go func() {
		router := gin.Default()
		//Proceed VM router
		router.POST("/token", endpoint.MakeTokenEndpoint())
		router.POST("/vm", endpoint.MakeValidateTokenEndpoint(), gin.WrapH(vmHandler))
		router.POST("/disk", endpoint.MakeValidateTokenEndpoint(), gin.WrapH(diskHandler))
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
