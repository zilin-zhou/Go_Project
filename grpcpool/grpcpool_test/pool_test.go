package grpcpool_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcpool/greeter/proto"
	pool2 "grpcpool/pool"
	"testing"
)

func BenchmarkGrpcWithOutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			b.Error(err)
		}
		in := &proto.HelloRequest{
			Msg: "hello Server",
		}
		c := proto.NewGreeterClient(conn)
		c.SayHello(context.Background(), in)
	}
}

func BenchmarkGrpcWithPool(b *testing.B) {

	pool, err := pool2.GetPool("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := pool.Get()
		in := &proto.HelloRequest{
			Msg: "Hello Server",
		}
		c := proto.NewGreeterClient(conn)
		c.SayHello(context.Background(), in)
		pool.Put(conn)
	}
}
