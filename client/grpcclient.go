package client

import (
	"context"
	"fmt"
	"go-microservice/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("127.0.0.1:8081", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewStringSvcClient(conn)
	capitalStr, err := client.Uppercase(context.Background(), &pb.UppercaseRequest{S: "thank you for interaction us"})
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Print(capitalStr.V)
	count, err := client.Count(context.Background(), &pb.CountRequest{S: "thank you for interaction us"})
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Print(count.V)
}
