package port

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/kntajus/sampleapp/protos"
)

func Update(data io.Reader, portClient protos.PortDomainServiceClient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := portClient.UpdatePorts(ctx)
	if err != nil {
		return err
	}

	portChan := make(chan *protos.PortWithID)
	go func() {
		parse(data, portChan)
	}()

	for port := range portChan {
		if err := stream.Send(port); err != nil {
			return err
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func parse(reader io.Reader, out chan<- *protos.PortWithID) {
	defer close(out)

	dec := json.NewDecoder(reader)
	// Skip over opening { in file
	_, err := dec.Token()
	if err != nil {
		log.Print(err)
		return
	}

	for dec.More() {
		port, err := readPort(dec)
		if err != nil {
			log.Print(err)
			return
		}
		out <- port
	}
}

func readPort(dec *json.Decoder) (*protos.PortWithID, error) {
	// Read "key", which is the ID
	t, err := dec.Token()
	if err != nil {
		return nil, err
	}
	id, ok := t.(string)
	if !ok {
		return nil, err
	}

	// Read "value", which is the port data
	var p protos.Port
	err = dec.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &protos.PortWithID{Id: id, Port: &p}, nil
}
