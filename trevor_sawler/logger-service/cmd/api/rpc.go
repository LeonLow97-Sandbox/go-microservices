package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

// RPCServer is the type for our RPC Server. Methods that take this as a receiver
// are available over RPC, as long as they are exported
type RPCServer struct {
}

// RPCPayload is the type for data we receive from RPC
type RPCPayload struct {
	Name string
	Data string
}

// LogInfo writes our payload to mongo (LogInfo must be in capital letter to export the method)
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	// writing to mongo
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC: " + payload.Name
	return nil
}
