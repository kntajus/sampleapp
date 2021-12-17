package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kntajus/sampleapp/ports/server"
	"github.com/kntajus/sampleapp/ports/store"
	"github.com/kntajus/sampleapp/protos"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal("Could not listen on port 10000")
	}

	grpcServer := grpc.NewServer()
	s := server.New(store.NewMap())
	protos.RegisterPortDomainServiceServer(grpcServer, s)

	exitChan := make(chan struct{})

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("Serve error: %v", err)
			exitChan <- struct{}{}
		}
	}()

	go func() {
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
		stop := <-stopChan

		log.Printf("Signal '%v' received, shutting down", stop)
		grpcServer.GracefulStop()
		exitChan <- struct{}{}
	}()

	<-exitChan
}
