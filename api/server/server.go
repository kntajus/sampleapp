package server

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/kntajus/sampleapp/api/port"
	"github.com/kntajus/sampleapp/protos"
)

const (
	updatePortsURL = "/update"
)

type Server struct {
	portClient protos.PortDomainServiceClient
	server     *http.Server
}

func New(portClient protos.PortDomainServiceClient) *Server {
	return &Server{
		portClient: portClient,
		server:     &http.Server{Addr: ":8080"},
	}
}

func (s *Server) Serve() error {
	http.HandleFunc(updatePortsURL, s.updatePorts)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) updatePorts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != updatePortsURL {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	file, err := os.Open("/data/ports.json")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = port.Update(file, s.portClient)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
