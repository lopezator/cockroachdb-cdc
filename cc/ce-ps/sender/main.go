package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	ce_ps "github.com/lopezator/crdb-cdc/cc/ce-ps"
)

func cdcHandler(w http.ResponseWriter, r *http.Request) {
	// read request body line by line
	scanner := bufio.NewScanner(r.Body)
	for scanner.Scan() {
		// Unmarshal cdc event
		var cdcChange *ce_ps.CdcChange
		if err := json.Unmarshal(scanner.Bytes(), &cdcChange); err != nil {
			log.Fatalf("unable to unmarshal body %s\n%v\n", scanner.Text(), err)
		}

		fmt.Println("CDC event received on webhook!")

		// send out to cloudevents - pubsub
		cloudEventsPubsubHandler(cdcChange)
	}
}

func cloudEventsPubsubHandler(cdcChange *ce_ps.CdcChange) {
	// Get PubSub project ID
	projectId := os.Getenv("PUBSUB_PROJECT_ID")
	if projectId == "" {
		log.Fatalln("You must set the PUBSUB_PROJECT_ID env")
	}

	// Get PubSub topic
	topic := os.Getenv("PUBSUB_TOPIC")
	if topic == "" {
		log.Fatalln("You must set the PUBSUB_TOPIC env")
	}

	// Create pubsub transport
	ps, err := pubsub.New(context.Background(), pubsub.WithProjectID(projectId), pubsub.WithTopicID(topic))
	if err != nil {
		log.Fatalf("failed to crearte pubsub transport, %v\n", err)
	}

	// Create cloudevents client
	cloudEvents, err := cloudevents.NewClient(ps, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v\n", err)
	}

	// Create a cloud event notification
	event := cloudevents.NewEvent(cloudevents.VersionV03)
	event.SetType("com.cockroachlabs.createchangefeed.update")
	event.SetSource("mysubdomain.cockroachcloud.com")
	if err := event.SetData(cdcChange); err != nil {
		log.Fatalf("failed to set event data, %v\n", err)
	}

	// Send over the client's configured transport (pubsub)
	_, _, err = cloudEvents.Send(context.Background(), event)
	if err != nil {
		log.Fatalf("failed to send: %v\n", err)
	}

	fmt.Println("Cloud event notification sent over pubsub!")
}

func main() {
	http.HandleFunc("/cdc/", cdcHandler)
	fmt.Println("started server on localhost:8080/cdc")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}