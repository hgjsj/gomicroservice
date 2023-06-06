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
		transport.EncodeCloudResponse,
	)
	diskHandler := httptransport.NewServer(
		endpoint.MakeDiskEndpoint(),
		transport.DecodeDiskRequest,
		transport.EncodeCloudResponse,
	)

	go func() {
		router := gin.Default()

		//Proceed VM router
		router.POST("/vm", gin.WrapH(vmHandler))
		router.POST("/disk", gin.WrapH(diskHandler))
		// router.GET("/vm")
		// router.GET("/vm/[id]")
		listen := fmt.Sprintf(":%d", port)
		router.Run(listen)
	}()
}
