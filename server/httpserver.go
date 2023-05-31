package server

import (
	"fmt"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-microservice/endpoint"
	"go-microservice/service"
	"go-microservice/transport"
	"net/http"
	"os"
)

func LaunchHttpSever(port int, proxy string, consulport int) {
	logger := log.NewLogfmtLogger(os.Stderr)
	listen := fmt.Sprintf(":%d", port)
	var svc service.StringService
	svc = service.NewStringService()

	svc = transport.ProxyingMiddleware(proxy, logger)(svc)

	svc = service.NewLogging(logger, svc, "http")

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})
	svc = service.NewInstrument(requestCount, requestLatency, countResult, svc)

	uppercaseHandler := httptransport.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)
	countHandler := httptransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)
	go func() {

		http.Handle("/uppercase", uppercaseHandler)
		http.Handle("/count", countHandler)
		http.Handle("/metrics", promhttp.Handler())
		logger.Log("msg", "HTTP", "addr", listen)
		logger.Log("err", http.ListenAndServe(listen, nil))
	}()

}
