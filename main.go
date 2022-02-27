package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	v1 "wallet-service/pb/wallet-service"
	"wallet-service/service"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	v1.RegisterWalletServiceServer(server, new(service.WalletService))
	if err = server.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
