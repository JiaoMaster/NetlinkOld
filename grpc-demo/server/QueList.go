package server

import (
	"context"
	"encoding/json"
	pb "grpc-demo"
	"grpc-demo/bapi"
)

type QueListServer struct {
}

func NewQueListServer() *QueListServer {
	return &QueListServer{}
}

func (q *QueListServer) GetQueList(c context.Context, r *pb.GetQueListRequest) (*pb.GetQueListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8080")
	body, err := api.GetQueList(c, r.GetPage(), r.GetAmount())
	if err != nil {
		return nil, err
	}

	queList := pb.GetQueListReply{}
	err = json.Unmarshal(body, &queList)
	if err != nil {
		return nil, err
	}
	return &queList, nil
}
