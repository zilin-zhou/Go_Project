package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcpool/greeter/proto"
	"grpcpool/pool"
	"log"
)

var (
	addr = flag.String("addr", "localhost:5001", "")
)

func main() {
	flag.Parse()
	//建立连接
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()
	//sayHello(conn)

	//使用连接池
	pool, err := pool.GetPool(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	conn := pool.Get()
	sayHello(conn)
	pool.Put(conn)

}
func sayHello(conn *grpc.ClientConn) {
	c := proto.NewGreeterClient(conn)
	ctx := context.Background()
	in := &proto.HelloRequest{
		Msg: "Hello Serve",
	}
	r, err := c.SayHello(ctx, in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Msg)
}
