package server

import (
	"context"
	"io"

	"github.com/kntajus/sampleapp/ports/store"
	"github.com/kntajus/sampleapp/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PortStore interface {
	UpsertPort(context.Context, *protos.PortWithID) error
	GetPort(context.Context, string) (*protos.Port, error)
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

func (s *Server) GetPort(ctx context.Context, req *protos.GetPortRequest) (*protos.GetPortResponse, error) {
	p, err := s.store.GetPort(ctx, req.GetId())
	if err != nil {
		if err == store.ErrNotFound {
			return nil, status.Error(codes.NotFound, "unknown port identifier")
		}
		return nil, err
	}
	return &protos.GetPortResponse{Port: p}, nil
}
