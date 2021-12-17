package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kntajus/sampleapp/protos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const knownPortID = "KNOWN"

func TestUpdatePorts(t *testing.T) {
	mockUpdate := &mockUpdatePortsClient{}
	s := New(&mockClient{updateClient: mockUpdate})
	s.dataFile = "../../data/ports.json"

	response := httptest.NewRecorder()
	s.updatePorts(response, httptest.NewRequest("POST", "/update", nil))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 1632, mockUpdate.sendCount)
}

func TestGetPort(t *testing.T) {
	knownName := "Somewhere"
	portResponse := &protos.GetPortResponse{
		Port: &protos.Port{
			Name: knownName,
		},
	}
	s := New(&mockClient{portResponse: portResponse})

	response := httptest.NewRecorder()
	s.getPort(response, httptest.NewRequest("GET", "/ports/"+knownPortID, nil))

	assert.Equal(t, http.StatusOK, response.Code)

	decoder := json.NewDecoder(response.Body)
	content := &protos.Port{}
	err := decoder.Decode(content)
	require.NoError(t, err)
	assert.Equal(t, knownName, content.GetName())
}

type mockClient struct {
	updateClient protos.PortDomainService_UpdatePortsClient
	portResponse *protos.GetPortResponse
}

func (c *mockClient) UpdatePorts(ctx context.Context, opts ...grpc.CallOption) (protos.PortDomainService_UpdatePortsClient, error) {
	return c.updateClient, nil
}
func (c *mockClient) GetPort(ctx context.Context, in *protos.GetPortRequest, opts ...grpc.CallOption) (*protos.GetPortResponse, error) {
	if in.GetId() != knownPortID {
		return nil, assert.AnError
	}
	return c.portResponse, nil
}

type mockUpdatePortsClient struct {
	grpc.ClientStream
	sendCount int
}

func (c *mockUpdatePortsClient) Send(*protos.PortWithID) error {
	c.sendCount++
	return nil
}
func (c *mockUpdatePortsClient) CloseAndRecv() (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}
