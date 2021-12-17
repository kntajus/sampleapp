package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kntajus/sampleapp/api/server"
	"github.com/kntajus/sampleapp/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("ports:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	portClient := protos.NewPortDomainServiceClient(conn)
	s := server.New(portClient)

	exitChan := make(chan struct{})

	go func() {
		if err := s.Serve(); err != nil {
			if err != http.ErrServerClosed {
				log.Printf("Serve error: %v", err)
				exitChan <- struct{}{}
			}
		}
	}()

	go func() {
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
		stop := <-stopChan

		log.Printf("Signal '%v' received, shutting down", stop)
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		if err := s.Shutdown(ctx); err != nil {
			log.Printf("Server shutdown error: %v", err)
		}
		exitChan <- struct{}{}
	}()

	<-exitChan
}
