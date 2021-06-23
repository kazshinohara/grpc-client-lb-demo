package main

import (
	"context"
	pb "github.com/kazshinohara/pb/grpc-client-lb-demo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"os"
	"time"
)

var domain = os.Getenv("DOMAIN")
var port = os.Getenv("PORT")

func main() {
	conn, err := grpc.Dial("dns:///" + domain + ":" + port, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTesterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()
	for i := 0; i < 30; i++ {
		r, err := c.GetHostInfo(ctx, &emptypb.Empty{})
		if err != nil {
			log.Fatalf("could not get hostinfo: %v", err)
		}
		log.Printf(r.HostnameAndIp)
		time.Sleep(time.Second * 1)
	}
}

