/* Written by Mutlu Polatcan
   01.03.2019
   Cloud Day 2019 - Custom Metric Autoscaling with GKE
   API Server that writes simple uuid to Pubsub when request.
   (Demo purpose only)
*/
package main

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"context"
	"github.com/google/uuid"
)

type ApiServer struct {
	PubsubClient *pubsub.Client
	PubsubTopic *pubsub.Topic
	Context context.Context
	TopicName string
	ProjectId string
}

func (apiServer *ApiServer) PreparePubsubComponents() {
	// Prepare PubSub client
	apiServer.Context = context.Background()

	client, err := pubsub.NewClient(apiServer.Context, apiServer.ProjectId)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	} else {
		apiServer.PubsubClient = client
	}

	// Create or get existing topic
	apiServer.PubsubTopic, err = apiServer.PubsubClient.CreateTopic(apiServer.Context, apiServer.TopicName)

	if err != nil {
		apiServer.PubsubTopic = apiServer.PubsubClient.Topic(apiServer.TopicName)
	}
}

func (apiServer *ApiServer) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	response, _ := json.Marshal(map[string]string{"status": "true"})
	w.Write(response)
}

func (apiServer *ApiServer) Stress(w http.ResponseWriter, r *http.Request) {
	uuid_val, _ := uuid.NewUUID()

	result := apiServer.PubsubTopic.Publish(apiServer.Context, &pubsub.Message{
		Data: []byte(uuid_val.String()),
	})

	id, err := result.Get(apiServer.Context)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Published Message %s; Message ID: %v\n", uuid_val, id)
}

func (apiServer *ApiServer) Start(projectID string, topicName string, port string) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", os.Getenv("GCP_AUTH_FILE_LOCATION"))

	apiServer.ProjectId = projectID
	apiServer.TopicName = topicName
	apiServer.PreparePubsubComponents()

	http.HandleFunc("/", apiServer.Healthcheck)
	http.HandleFunc("/stress", apiServer.Stress)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}


func main() {
	var apiServer = &ApiServer{}
	apiServer.Start(os.Getenv("GCP_PROJECT_ID"), os.Getenv("PUBSUB_TOPIC_NAME"), os.Getenv("API_SERVER_PORT"))
}
