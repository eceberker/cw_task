package gcpclient

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/eceberker/cw_task/helpers"
	"github.com/eceberker/cw_task/models"
	"google.golang.org/api/iterator"
)

// Publish publishes message in default topic in GCP Pub/Sub
func (p *Publisher) Publish(msg string) models.PublishResponseModel {

	ctx, cancel := context.WithTimeout(context.Background(), p.config.PublishTimeout)
	defer cancel()

	var response models.PublishResponseModel

	env := helpers.GetEnv()
	TopicID := env["TOPIC_ID"]

	result := p.client.Topic(TopicID).Publish(ctx, &pubsub.Message{Data: []byte(msg)})
	<-result.Ready()

	id, err := result.Get(ctx)

	if err != nil {
		response.MessageText = fmt.Sprintf("An error occured %v.", err)
		response.Status = 500
		return response
	}

	response.Status = 200
	response.MessageID = id
	response.MessageText = "Log is sent succesfully."

	return response
}

// List returns topic list from GCP PubSub
func (p *Publisher) List() ([]*pubsub.Topic, error) {

	ctx := context.Background()

	var topics []*pubsub.Topic

	it := p.client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Next: %v", err)
		}
		fmt.Printf("topic: %v", topic.ID())
		topics = append(topics, topic)
	}

	return topics, nil
}
