package endpoint

import (
	"context"
	"go-microservice/pb"
	"go-microservice/service"
	"go-microservice/transport"

	"github.com/go-kit/kit/endpoint"
	gt "github.com/go-kit/kit/transport/grpc"
)

type Endpoints struct {
	Uppercase endpoint.Endpoint
	Count     endpoint.Endpoint
}

func makeGrpcUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return transport.UppercaseResponse{V: v, Err: err.Error()}, err
		}
		return transport.UppercaseResponse{V: v, Err: ""}, nil
	}
}

func makeGrpcCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.CountRequest)
		v := svc.Count(req.S)

		return transport.CountResponse{V: v}, nil
	}
}

func NewGrpcServer(endpoints Endpoints) pb.StringSvcServer {
	return &transport.GrpcServer{
		GrpcUppercase: gt.NewServer(
			endpoints.Uppercase,
			transport.DecodeGrpcUppercaseRequest,
			transport.EncodeGrpcUppercaseResponse,
		),
		GrpcCount: gt.NewServer(
			endpoints.Count,
			transport.DecodeGrpcCountRequest,
			transport.EncodeGrpcCountResponse,
		),
	}
}

func MakeEndpoints(s service.StringService) Endpoints {
	return Endpoints{
		Uppercase: makeGrpcUppercaseEndpoint(s),
		Count:     makeGrpcCountEndpoint(s),
	}
}
