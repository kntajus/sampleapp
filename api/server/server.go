package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/kntajus/sampleapp/api/port"
	"github.com/kntajus/sampleapp/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	updatePortsURL = "/update"
	getPortURLStem = "/ports/"
)

var getPortURLFormat = regexp.MustCompile("^/ports/[A-Z]{5}$")

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
	http.HandleFunc(getPortURLStem, s.getPort)
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

func (s *Server) getPort(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if !getPortURLFormat.MatchString(path) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	portID := path[len(getPortURLStem):]

	resp, err := s.portClient.GetPort(context.Background(), &protos.GetPortRequest{Id: portID})
	if err != nil {
		httpCode := http.StatusInternalServerError
		st := status.Convert(err)
		if st.Code() == codes.NotFound {
			httpCode = http.StatusNotFound
		}
		w.WriteHeader(httpCode)
		return
	}

	json.NewEncoder(w).Encode(resp.GetPort())
}
