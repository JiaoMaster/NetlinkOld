package main

import (
	"context"
	"fmt"
	pb "github.com/NetLinkOld/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClientConn(c context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(c, target, opts...)
}

func main() {
	c := context.Background()
	clientConn, err := GetClientConn(c, "localhost:8000", nil)
	fmt.Println(err)
	defer clientConn.Close()
	queServiceClient := pb.NewQueListServiceClient(clientConn)
	resp, err := queServiceClient.GetQueList(
		c,
		&pb.GetQueListRequest{
			Page:     "1",
			Amount:   "10",
			Ch:       3,
			Location: "中国",
		},
	)

	fmt.Println(resp)
}
