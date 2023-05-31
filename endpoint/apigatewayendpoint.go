package endpoint

import (
	"go-microservice/transport"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeAPIGatewayEndpoint(ins *consul.Instancer, path string, logger log.Logger) endpoint.Endpoint {
	var apigatewayFactory sd.Factory
	apigatewayFactory = stringsvcFactory("POST", path, logger)
	apigatewayEndpoiter := sd.NewEndpointer(ins, apigatewayFactory, logger)
	apigatewayBalancer := lb.NewRoundRobin(apigatewayEndpoiter)
	apigatewayEndpoint := lb.Retry(3, 250*time.Millisecond, apigatewayBalancer)
	return apigatewayEndpoint

}

func stringsvcFactory(method string, path string, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path

		var (
			enc httptransport.EncodeRequestFunc
			dec httptransport.DecodeResponseFunc
		)
		switch path {
		case "/uppercase":
			enc, dec = transport.EncodeRequest, transport.DecodeUppercaseResponse
		case "/count":
			enc, dec = transport.EncodeRequest, transport.DeccodeCountResponse
		default:
			logger.Log("err", "unknown stringsvc path %q", path)
			os.Exit(1)
		}

		return httptransport.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}
