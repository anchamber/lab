package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/anchamber/lab/services/lines/server"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/anchamber/lab/proto/line"
)

func main() {
	srv := &server.Server{}
	newServer := grpc.NewServer()
	pb.RegisterLineServer(newServer, srv)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal(err)
	}
	// Start Server
	go func() {
		log.Println("Setting proxy newServer port", 8080)
		if err := newServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	connectEventStore()

	// Graceful Shutdown
	waitForShutdown(newServer)
}

func connectEventStore() {
	settings, err := esdb.ParseConnectionString("esdb://admin:changeit@eventstore:2113?tls=false")

	if err != nil {
		panic(err)
	}
	db, err := esdb.NewClient(settings)

	// endregion createClient
	if err != nil {
		panic(err)
	}

	// region createEvent
	testEvent := struct {
		Id            string
		ImportantData string
	}{
		Id:            uuid.NewUUID().String(),
		ImportantData: "Connected to eventstore",
	}

	data, err := json.Marshal(testEvent)

	if err != nil {
		panic(err)
	}

	eventData := esdb.EventData{
		ContentType: esdb.JsonContentType,
		EventType:   "TestEvent",
		Data:        data,
	}
	_, err = db.AppendToStream(context.Background(), "client-connections", esdb.AppendToStreamOptions{}, eventData)
	// endregion appendEvents

	if err != nil {
		panic(err)
	}
}

func waitForShutdown(server *grpc.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	server.GracefulStop()

	log.Println("Shutting down")
	os.Exit(0)
}
