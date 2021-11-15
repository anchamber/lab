package main

import (
	"context"
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/pborman/uuid"
)

func main() {
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
