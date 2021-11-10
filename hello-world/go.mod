module hello-world

go 1.16

replace hello-world/proto/google/api => ./proto/google/api

replace hello-world/server => ./server

replace hello-world/proto/google => ./proto/google

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/spf13/cobra v1.2.1
	golang.org/x/net v0.0.0-20210913180222-943fd674d43e
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)
