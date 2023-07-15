package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpcpool/greeter/proto"
	"log"
	"net"
)

var (
	//命令行参数
	port = flag.Int("port", 5001, "")
)

func main() {
	//从命令行取到参数
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	//注册服务
	proto.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	fmt.Println("Server Recv " + in.Msg)
	return &proto.HelloReply{
		Msg: "Hello client",
	}, nil
}
