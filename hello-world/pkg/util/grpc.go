package util

import (
	"google.golang.org/grpc"
	"net/http"
	"strings"
)

func GrpcHandleFunc(grpcServer *grpc.Server,otherhandler http.Handler) http.Handler {
	if otherhandler == nil {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			grpcServer.ServeHTTP(writer,request)
		})
	}
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//http 2.0
		if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"),"application/grpc") {
			grpcServer.ServeHTTP(writer,request)
		}else {
			otherhandler.ServeHTTP(writer,request)
		}
	})
}
