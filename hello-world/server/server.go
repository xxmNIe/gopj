package server

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"hello-world/pkg/util"
	pb "hello-world/proto"
	"log"
	"net"
	"net/http"
)

var (
	ServerPort string
	CertName string
	CertPemPath string
	CertKeyPath string
	EndPoint string
)

func Server() (err error) {
	//1.一个conn连接
	EndPoint  = ":"+ServerPort
	conn,err := net.Listen("tcp",EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	tlsConfig := util.GetTLSConfig(CertPemPath,CertKeyPath)
	//2.创建
	srv :=createInternalServer(conn,tlsConfig)


	log.Printf("gRPC and https listen on: %s\n", ServerPort)
	if err = srv.Serve(conn);err !=nil {
			log.Printf("ListenAndServe: %v\n", err)
	}
	return err
}
func createInternalServer(conn net.Listener, tlsConfig *tls.Config) (*http.Server){

	/**
	credentials.NewClientTLSFromFile：从客户机的输入证书文件构造TLS凭证
	grpc.WithTransportCredentials：配置一个连接级别的安全凭据(例：TLS、SSL)，返回值为type DialOption
	grpc.DialOption：DialOption选项配置我们如何设置连接（其内部具体由多个的DialOption组成，决定其设置连接的内容）
	 */
	var opts []grpc.ServerOption

	//grpc server
	creds,err := credentials.NewServerTLSFromFile(CertPemPath,CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	opts = append(opts,grpc.Creds(creds))

	//此处创建了一个没有注册服务的grpc服务端
	//grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	//register grpc pb
	//注册grpc服务
	pb.RegisterHelloWorldServer(grpcServer,NewHelloService())


	//gw server
	//创建grpc-gateway关联组件
	ctx := context.Background()
	//dcreds,err := credentials.NewClientTLSFromFile(CertPemPath,CertName)
	//if err != nil {
	//	log.Printf("Failed to create client TLS credentials %v", err)
	//}
	//dopts :=[]grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	dopts :=[]grpc.DialOption{grpc.WithInsecure()}
	//gw 的runtime
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb
	//注册hellow-world到grpc-gateway
	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx,gwmux,EndPoint,dopts);err !=nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	mux := http.NewServeMux()

	//http服务
	mux.Handle("/",gwmux)

	return &http.Server{
		Addr: EndPoint,
		Handler: util.GrpcHandleFunc(grpcServer,mux),
		//TLSConfig: tlsConfig,
	}
}