package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
	"time"

	//gw "grpc-gateway/myservice/myservice"
	//gw2 "grpc-gateway/myservice/service2"
	gw3 "grpc-gateway/myservice/service3"
	pb3 "grpc-gateway/myservice/service3"
	"log"
	"net/http"
)

var (
	grpcServerPortEnd = flag.String("grpc-Server-Port-End", "localhost:5001", "")
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	//自定义接口
	mux.HandlePath("GET", "/ping", pingHandler)
	//自定义接口上传文件通过GRPC客户端传递到服务端哦
	mux.HandlePath("POST", "/uploadfile", UploadFileHandler)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//service2
	//if err := gw2.RegisterService2HandlerFromEndpoint(ctx, mux, *grpcServerPortEnd, opts); err != nil {
	//	log.Fatalln(err)
	//}

	//service3
	if err := gw3.RegisterService3HandlerFromEndpoint(ctx, mux, *grpcServerPortEnd, opts); err != nil {
		log.Fatalln(err)
	}
	return http.ListenAndServe(":8081", mux)

}
func pingHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	fmt.Fprintf(w, "{\"msg\":\"pong\"}")
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//读取文件
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//attachment为字段名
	f, header, err := r.FormFile("attachment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	//创建grpc连接
	conn, err := grpc.Dial(*grpcServerPortEnd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	//创建grpc客户端
	c := pb3.NewService3Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second+30)
	defer cancel()
	//取出元数据
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"file_name": header.Filename}))
	stream, err := c.EchoUpload(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf := make([]byte, 100)
	for {
		//向buf填充数据
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if n == 0 {
			break
		}
		stream.Send(&pb3.UploadRequest{
			Size:    int64(n),
			Content: buf[:n],
		})

	}
	//客户端
	res, err := stream.CloseAndRecv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, res.Path)
}
