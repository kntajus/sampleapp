package store

import (
	"context"

	"github.com/kntajus/sampleapp/protos"
)

type Map struct {
	ports map[string]*protos.Port
}

func NewMap() *Map {
	return &Map{ports: make(map[string]*protos.Port)}
}

func (m *Map) UpsertPort(_ context.Context, port *protos.PortWithID) error {
	m.ports[port.GetId()] = port.GetPort()
	return nil
}

func (m *Map) GetPort(id string) *protos.Port {
	return m.ports[id]
}
