package server

import (
	"fmt"
	"go-microservice/client"
	"go-microservice/endpoint"
	"go-microservice/transport"
	"net/http"

	"os"
	"os/signal"

	"syscall"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func LaunchGW(httpport int, consulport int) {
	httpAddr := fmt.Sprintf(":%d", httpport)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	ins := client.NewConsulClient("stringservice", consulport, []string{"stringservice"}, logger)

	uppercase := endpoint.MakeAPIGatewayEndpoint(ins, "/uppercase", logger)
	count := endpoint.MakeAPIGatewayEndpoint(ins, "/count", logger)

	uppercaseHandler := httptransport.NewServer(uppercase, transport.DecodeUppercaseRequest, transport.EncodeResponse)
	countHandler := httptransport.NewServer(count, transport.DecodeCountRequest, transport.EncodeResponse)

	http.Handle("/stringsvc/uppercase", uppercaseHandler)
	http.Handle("/stringsvc/count", countHandler)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport.
	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errc <- http.ListenAndServe(httpAddr, nil)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
