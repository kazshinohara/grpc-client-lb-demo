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
	"strings"
)

const (
	port = ":8080"
)

type TestServer struct {
	pb.UnimplementedTesterServer
}

func chop(s string) string {
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}
	return s
}

func (s *TestServer) GetHostInfo(ctx context.Context, empty *emptypb.Empty) (*pb.HostInfo, error) {
	//hostname, err := exec.Command("hostname").Output()
	//if err != nil{
	//	fmt.Printf("%v", err)
	//}
	//ip, err := exec.Command("hostname", "-i").Output()
	//if err != nil{
	//	fmt.Printf("%v", err)
	//}
	hostname, err := os.Hostname()
	if err != nil{
		fmt.Printf("%v", err)
	}
	//return &pb.HostInfo{HostnameAndIp: chop(string(hostname)) + " : " + chop(string(ip)) }, nil
	log.Printf("request received")
	return &pb.HostInfo{HostnameAndIp: string(hostname)}, nil
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
