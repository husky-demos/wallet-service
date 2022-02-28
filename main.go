package main

import (
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
	walletService "wallet-service/pb/wallet-service"
	"wallet-service/service"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		glog.Fatalln(err)
	}
	server := grpc.NewServer()
	walletService.RegisterWalletServiceServer(server, new(service.WalletService))
	if err = server.Serve(lis); err != nil {
		glog.Fatalln(err)
	}
}
