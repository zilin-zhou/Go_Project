package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc-gateway/myservice/service3"
	"io"
	"log"
	"net"
	"os"
)

// myservice/myservice的服务

/*
type Server1 struct {
	service1.UnimplementedMyServiceServer
}

// myservice/myservice的一个服务
func (s *Server1) Echo(ctx context.Context, in *service1.StringMessage) (*service1.StringMessage, error) {
	fmt.Printf("server RECV:%+v\n", in)
	return &service1.StringMessage{
		Value: "request ok",
	}, nil
}

*/
/***********************************************************************/

/*
// myservice/service2的服务

	type Server2 struct {
		service2.UnimplementedService2Server
	}

// myservice/service2的三个服务

	func (s *Server2) Echo(ctx context.Context, in *service2.SimpleMessage) (*service2.SimpleMessage, error) {
		fmt.Printf("server RECV:%+v\n", in)
		return in, nil
	}

	func (s *Server2) EchoBody(ctx context.Context, in *service2.SimpleMessage) (*service2.SimpleMessage, error) {
		fmt.Printf("server RECV:%+v\n", in)
		return in, nil
	}

	func (s *Server2) EchoDelete(ctx context.Context, in *service2.SimpleMessage) (*service2.SimpleMessage, error) {
		fmt.Printf("server RECV:%+v\n", in)
		return in, nil
	}
*/
var (
	port = flag.Int("port", 5001, "")
)

/***********************************************************************/

// myservice/service2的服务

type Server3 struct {
	service3.UnimplementedService3Server
}

func (s *Server3) Echo(ctx context.Context, in *service3.SimpleMessage) (*service3.SimpleMessage, error) {
	fmt.Printf("server RECV:%+v\n", in)
	return in, nil
}
func (s *Server3) EchoBody(ctx context.Context, in *service3.SimpleMessage) (*service3.SimpleMessage, error) {
	fmt.Printf("server RECV:%+v\n", in)
	return in, nil
}
func (s *Server3) EchoDelete(ctx context.Context, in *service3.SimpleMessage) (*service3.SimpleMessage, error) {
	fmt.Printf("server RECV:%+v\n", in)
	return in, nil
}

// 上传文件
func (s *Server3) EchoUpload(stream service3.Service3_EchoUploadServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Fatalln("metadata get failed!")
	}
	fileName := md["file_name"][0]
	fmt.Println("server RECV:%v\n", fileName)
	filePath := "myservice/server/upload/" + fileName
	dst, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer dst.Close()
	//接受文件内容
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		//写入
		dst.Write(req.Content[:req.Size])
	}
	// 服务器返回
	stream.SendAndClose(&service3.UploadResponse{
		Path: filePath,
	})
	return nil
}

/***********************************************************************/

func main() {
	flag.Parse()
	//创建一个监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln(err)
	}
	//创建NewServer
	s := grpc.NewServer()
	//注册服务
	//myservice.RegisterMyServiceServer(s, &Server1{})
	//service2.RegisterService2Server(s, &Server2{})
	service3.RegisterService3Server(s, &Server3{})
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
