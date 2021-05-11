module hello-world

go 1.16

replace hello-world/proto/google/api => ./proto/google/api

replace hello-world/server => ./server

replace hello-world/proto/google => ./proto/google

require (
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/spf13/cobra v1.0.1-0.20201006035406-b97b5ead31f7
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4
	google.golang.org/genproto v0.0.0-20210429181445-86c259c2b4ab // indirect
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)
