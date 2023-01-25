package main

import (
	"context"
	"log"
	"net"

	"github.com/alessandro-p/go-grpc-mongo-sample/blog/proto"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/server/api"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/server/repository"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/server/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	client := repository.Init()
	defer client.Disconnect(context.Background())

	address := "0.0.0.0:50051"
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Unable to listen on address %s, err: %v\n", address, err)
	}

	log.Printf("Listening on address %s\n", address)

	tlsEnabled := false
	opts := utils.GetGrpcOptions(tlsEnabled)
	grpcServer := grpc.NewServer(opts...)

	// NOTE: enable reflection through evans (https://github.com/ktr0731/evans)
	reflection.Register(grpcServer)

	postServiceServer := api.PostController{
		PostRepository: &repository.PostRepository{
			Client: client,
		},
	}

	proto.RegisterPostServiceServer(grpcServer, &postServiceServer)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Unable to serve grpc on listener: %v\n", err)
	}
}
