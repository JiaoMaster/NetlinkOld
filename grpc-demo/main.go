package main

import (
	pb "github.com/NetLinkOld/grpc-demo/proto"
	"github.com/NetLinkOld/grpc-demo/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterQueListServiceServer(s, server.NewQueListServer())
	reflection.Register(s)

	list, err := net.Listen("tcp", ":"+"8000")
	if err != nil {
		log.Fatal("net.Listen err: %v", err)
	}

	err = s.Serve(list)
	if err != nil {
		log.Fatal("server.Serve err: %v", err)
	}
}
