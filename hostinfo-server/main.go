package main

import (
	"context"
	"fmt"
	pb "github.com/kazshinohara/pb/grpc-client-lb-demo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
)

const (
	port = ":8080"
)

type TestServer struct {
	pb.UnimplementedTesterServer
}

func (s *TestServer) GetHostInfo(ctx context.Context, empty *emptypb.Empty) (*pb.HostInfo, error) {
	hostname, err := os.Hostname()
	if err != nil{
		fmt.Printf("%v", err)
	}
	log.Printf("request received")
	return &pb.HostInfo{Hostname: string(hostname)}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTesterServer(s, &TestServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
