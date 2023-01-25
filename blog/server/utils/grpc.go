package utils

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GetGrpcOptions(tlsEnabled bool) []grpc.ServerOption {
	opts := []grpc.ServerOption{}

	if tlsEnabled {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	return opts
}
