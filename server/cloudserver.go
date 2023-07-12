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
		router.POST("/vm", gin.WrapH(vmHandler))
		router.POST("/disk", gin.WrapH(diskHandler))
		router.GET("/disk/:id", endpoint.MakeDiskGetEndpoint())
		router.GET("/vm/:id", endpoint.MakeVMGetEndpoint())
		router.GET("/vm", endpoint.MakeListVMEndpoint())
		router.GET("/disk", endpoint.MakeListDiskEndpoint())
		router.PATCH("/disk/:id", endpoint.MakePatchDiskEndpoint())
		router.PATCH("/vm/:id", endpoint.MakePatchVMEndpoint())
		router.DELETE("/disk/:id", endpoint.MakeDeleteDiskEndpoint())
		router.DELETE("/vm/:id", endpoint.MakeDeleteVMEndpoint())
		listen := fmt.Sprintf(":%d", port)
		router.Run(listen)
	}()
}
