package transport

import (
	"context"
	gt "github.com/go-kit/kit/transport/grpc"
	"go-microservice/pb"
)

type GrpcServer struct {
	GrpcUppercase gt.Handler
	GrpcCount     gt.Handler
	pb.UnimplementedStringSvcServer
}

type GrpcUppercaseReq struct {
}

func (s *GrpcServer) Uppercase(ctx context.Context, in *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	_, str, err := s.GrpcUppercase.ServeGRPC(ctx, in.S)
	if err != nil {
		return nil, err
	}
	return str.(*pb.UppercaseResponse), nil
}

func (s *GrpcServer) Count(ctx context.Context, in *pb.CountRequest) (*pb.CountResponse, error) {
	_, str, err := s.GrpcCount.ServeGRPC(ctx, in.S)
	if err != nil {
		return nil, err
	}
	return str.(*pb.CountResponse), nil
}

func DecodeGrpcUppercaseRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(string)

	return UppercaseRequest{S: req}, nil
}

func EncodeGrpcUppercaseResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UppercaseResponse)
	return &pb.UppercaseResponse{V: res.V, Err: res.Err}, nil
}

func DecodeGrpcCountRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(string)
	return CountRequest{S: req}, nil
}

func EncodeGrpcCountResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(CountResponse)
	return &pb.CountResponse{V: int32(res.V)}, nil
}
