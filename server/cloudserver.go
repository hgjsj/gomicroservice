package server

import (
	"context"
	"fmt"
	"go-microservice/endpoint"
	"go-microservice/service"
	"os"

	"github.com/gin-gonic/gin"
)

const socketPath = "/run/spire/sockets/agent.sock"
const serverSpiffeID = "spiffe://example.org/gomicroservice/cloud"

func LauchCloudServer(port int) {
	router := gin.Default()
	go func() {
		newTokenFunc := endpoint.MakeTokenEndpoint()
		validateTokenFunc := endpoint.MakeValidateTokenEndpoint()
		if _, err := os.Stat(socketPath); err == nil{
			jwts, err := service.NewSpiffeJWTSource(context.Background(), socketPath)
			if err == nil {
				fmt.Printf("unable to create JWTSource: %s", err.Error())
				return
			} 
			validateTokenFunc = endpoint.MakeValidateSpiffeJWTEndpoint(context.Background(), jwts, []string{serverSpiffeID})
			newTokenFunc = endpoint.MakeSpiffeJWTEndpoint(context.Background(), jwts, serverSpiffeID)
			defer jwts.Close()
		}
		//Proceed VM router
		router.POST("/token", newTokenFunc)
		router.POST("/vm", validateTokenFunc, endpoint.MakeVMPostEndpoint())
		router.POST("/disk", validateTokenFunc, endpoint.MakeDiskPostEndpoint())
		router.GET("/disk/:id", validateTokenFunc, endpoint.MakeDiskGetEndpoint())
		router.GET("/vm/:id", validateTokenFunc, endpoint.MakeVMGetEndpoint())
		router.GET("/vm", validateTokenFunc, endpoint.MakeListVMEndpoint())
		router.GET("/disk", validateTokenFunc, endpoint.MakeListDiskEndpoint())
		router.PATCH("/disk/:id", validateTokenFunc, endpoint.MakePatchDiskEndpoint())
		router.PATCH("/vm/:id", validateTokenFunc, endpoint.MakePatchVMEndpoint())
		router.DELETE("/disk/:id", validateTokenFunc, endpoint.MakeDeleteDiskEndpoint())
		router.DELETE("/vm/:id", validateTokenFunc, endpoint.MakeDeleteVMEndpoint())
		listen := fmt.Sprintf(":%d", port)
		router.Run(listen)
	}()
}
