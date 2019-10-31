package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/client"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	pscontext "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub/context"
	ce_ps "github.com/lopezator/crdb-cdc/cc/ce-ps"
)

func receive(ctx context.Context, event cloudevents.Event, resp *cloudevents.EventResponse) {
	fmt.Printf("Event Context: %+v\n", event.Context)
	fmt.Printf("Transport Context: %+v\n", pscontext.TransportContextFrom(ctx))

	data := &ce_ps.CdcChange{}
	if err := event.DataAs(data); err != nil {
		log.Fatalf("Got Data Error: %v\n", err)
	}
	fmt.Printf("Data: %+v\n", data)
	fmt.Printf("----------------------------\n")
}

func main() {
	ctx := context.Background()

	// Get PubSub Project ID
	projectId := os.Getenv("PUBSUB_PROJECT_ID")
	if projectId == "" {
		log.Fatalln("You must set the PUBSUB_PROJECT_ID env")
	}

	// Get PubSub Topic
	topic := os.Getenv("PUBSUB_TOPIC")
	if topic == "" {
		log.Fatalln("You must set the PUBSUB_TOPIC env")
	}

	// Get PubSub Subscription
	subscription := os.Getenv("PUBSUB_SUBSCRIPTION")
	if subscription == "" {
		log.Fatalln("You must set the PUBSUB_SUBSCRIPTION env")
	}

	// create pubsub client
	ps, err := pubsub.New(context.Background(),
		pubsub.WithProjectID(projectId),
		pubsub.WithTopicID(topic),
		pubsub.WithSubscriptionID(subscription))
	if err != nil {
		log.Fatalf("failed to create pubsub transport, %s", err.Error())
	}
	c, err := client.New(ps)
	if err != nil {
		log.Fatalf("failed to create client, %s", err.Error())
	}

	log.Println("Created client, listening...")

	if err := c.StartReceiver(ctx, receive); err != nil {
		log.Fatalf("failed to start pubsub receiver, %s", err.Error())
	}
}