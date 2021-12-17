package server

import (
	"context"
	"io"

	"github.com/kntajus/sampleapp/protos"
)

type PortStore interface {
	UpsertPort(context.Context, *protos.PortWithID) error
}

type Server struct {
	protos.UnimplementedPortDomainServiceServer
	store PortStore
}

func New(store PortStore) *Server {
	return &Server{store: store}
}

func (s *Server) UpdatePorts(stream protos.PortDomainService_UpdatePortsServer) error {
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err = s.store.UpsertPort(stream.Context(), port); err != nil {
			return err
		}
	}
}
