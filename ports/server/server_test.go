package server_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/kntajus/sampleapp/ports/server"
	"github.com/kntajus/sampleapp/ports/store"
	"github.com/kntajus/sampleapp/protos"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestUpdatePorts(t *testing.T) {
	tests := map[string]struct {
		store          server.PortStore
		msgs           []interface{}
		expectedStored map[string]*protos.Port
		expectedErr    string
	}{
		"Success": {
			store: store.NewMap(),
			msgs: []interface{}{
				&protos.PortWithID{Id: "ABCDE", Port: &protos.Port{Name: "Alphabet"}},
				&protos.PortWithID{Id: "ZYXWV", Port: &protos.Port{Name: "Tebahpla"}},
			},
			expectedStored: map[string]*protos.Port{
				"ABCDE": {Name: "Alphabet"},
				"ZYXWV": {Name: "Tebahpla"},
			},
		},
		"Receive Error": {
			msgs: []interface{}{
				errors.New("broken stream"),
			},
			expectedErr: "broken stream",
		},
		"Store Error": {
			store: &mockStore{upsertError: errors.New("broken store")},
			msgs: []interface{}{
				&protos.PortWithID{Id: "ABCDE", Port: &protos.Port{Name: "Alphabet"}},
			},
			expectedErr: "broken store",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := server.New(test.store)
			dataStream := &mockUpdatePortsServer{msgs: test.msgs}

			err := s.UpdatePorts(dataStream)
			if test.expectedErr != "" {
				assert.Contains(t, err.Error(), test.expectedErr)
				return
			}

			assert.NoError(t, err)
			for id, port := range test.expectedStored {
				storedPort := test.store.(*store.Map).GetPort(id)
				assert.Equal(t, port, storedPort)
			}
		})
	}
}

type mockUpdatePortsServer struct {
	grpc.ServerStream
	msgs []interface{}
}

func (m *mockUpdatePortsServer) Recv() (*protos.PortWithID, error) {
	if len(m.msgs) == 0 {
		return nil, io.EOF
	}
	msg := m.msgs[0]
	m.msgs = m.msgs[1:]
	if err, ok := msg.(error); ok {
		return nil, err
	}
	return msg.(*protos.PortWithID), nil
}

func (m *mockUpdatePortsServer) Context() context.Context          { return context.Background() }
func (m *mockUpdatePortsServer) SendAndClose(*emptypb.Empty) error { return nil }

type mockStore struct {
	upsertError error
}

func (b *mockStore) UpsertPort(context.Context, *protos.PortWithID) error { return b.upsertError }
