package endpoint

import (
	"context"
	"go-microservice/service"
	"go-microservice/transport"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return transport.UppercaseResponse{v, err.Error()}, nil
		}
		return transport.UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CountRequest)
		v := svc.Count(req.S)
		return transport.CountResponse{v}, nil
	}
}
