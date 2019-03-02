package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"os"
)

type PubsubConsumer struct {
	PubsubClient *pubsub.Client
	PubsubTopic *pubsub.Topic
	PubsubSubscription *pubsub.Subscription
	Context context.Context
	ProjectID string
	TopicName string
	SubscriptionID string
}

func (pubsubConsumer *PubsubConsumer) PreparePubsubComponents() {
	// Prepare PubSub client
	pubsubConsumer.Context = context.Background()

	client, err := pubsub.NewClient(pubsubConsumer.Context, pubsubConsumer.ProjectID)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	} else {
		pubsubConsumer.PubsubClient = client
	}

	// Get existing topic
	pubsubConsumer.PubsubTopic = pubsubConsumer.PubsubClient.Topic(pubsubConsumer.TopicName)

	// Create or get existing subscription
	subscription, err := pubsubConsumer.PubsubClient.CreateSubscription(pubsubConsumer.Context,
																	    pubsubConsumer.SubscriptionID,
																	    pubsub.SubscriptionConfig{
																			Topic: pubsubConsumer.PubsubTopic,
																	    })

	if err != nil {
		pubsubConsumer.PubsubSubscription = pubsubConsumer.PubsubClient.Subscription(pubsubConsumer.SubscriptionID)
	} else {
		pubsubConsumer.PubsubSubscription = subscription
	}
}

func (pubsubConsumer *PubsubConsumer) Start(projectID string, topicName string, subscriptionID string) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "auth.json")

	pubsubConsumer.ProjectID = projectID
	pubsubConsumer.TopicName = topicName
	pubsubConsumer.SubscriptionID = subscriptionID
	pubsubConsumer.PreparePubsubComponents()

	pubsubConsumer.PubsubSubscription.Receive(pubsubConsumer.Context, func(i context.Context, message *pubsub.Message) {
		fmt.Printf("Message arrived: %s", message.Data)
		message.Ack()
	})
}

func main() {
	var pubsubConsumer = &PubsubConsumer{}
	pubsubConsumer.Start(os.Getenv("GCP_PROJECT_ID"), os.Getenv("PUBSUB_TOPIC_NAME"), os.Getenv("PUBSUB_SUBSCRIPTION_ID"))
}



