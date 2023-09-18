package server

import (
	"context"
	"fmt"
	"go-microservice/endpoint"
	"go-microservice/service"
	"log"
	"net/http"
	"os"
	"time"
	"strings"
	"github.com/gin-gonic/gin"
)

const socketPath = "/run/spire/sockets/agent.sock"
const serverSpiffeID = "spiffe://example.org/gomicroservice/cloud"

func logFormat(param gin.LogFormatterParams) string{
	return fmt.Sprintf("[%s] %s %s %s %d %s %s\n", param.TimeStamp.Format(time.RFC3339),
	param.ClientIP, param.Method, param.Path, param.StatusCode, param.Request.UserAgent(), strings.TrimSuffix(param.ErrorMessage, "\n"), )
}

func LauchCloudServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(logFormat))
	newTokenFunc := endpoint.MakeTokenEndpoint()
	validateTokenFunc := endpoint.MakeValidateTokenEndpoint()
	if _, err := os.Stat(socketPath); err == nil{
		ctx := context.Background()
		jwts, err := service.NewSpiffeJWTSource(ctx, fmt.Sprintf("unix://%s",socketPath))
		if err != nil {
			fmt.Printf("unable to create JWTSource: %s", err.Error())
			return
		} 
		validateTokenFunc = endpoint.MakeValidateSpiffeJWTEndpoint(ctx, jwts, []string{serverSpiffeID})
		newTokenFunc = endpoint.MakeSpiffeJWTEndpoint(ctx, jwts, serverSpiffeID)
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
	fmt.Printf("[%s] Listening and serving HTTP on %d\n", time.Now().Format(time.RFC3339), port)
	go func() {
		// service connections
		if err := router.Run(listen); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

}
