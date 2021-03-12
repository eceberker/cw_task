package gcpclient

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/eceberker/cw_task/helpers"
	"google.golang.org/api/option"
)

// Publisher struct
type Publisher struct {
	topics map[string]*pubsub.Topic
	// TODO
	closed bool
	client *pubsub.Client
	config PublisherConfig
}

// PublisherConfig configs TODO for future improvements, publisher uses default settings for now
type PublisherConfig struct {
	// ProjectID is the Google Cloud Engine project ID.
	ProjectID string
	// TODO If false (default), `Publisher` tries to create a topic if there is none with the requested name.
	DoNotCreateTopicIfMissing bool
	// ConnectTimeout defines the timeout for connecting to Pub/Sub
	ConnectTimeout time.Duration
	// PublishTimeout defines the timeout for publishing messages.
	PublishTimeout time.Duration
	// Settings for client library.
	PublishSettings *pubsub.PublishSettings
	ClientOptions   []option.ClientOption
}

func (c *PublisherConfig) setDefaults() {

	env := helpers.GetEnv()
	cred := env["GOOGLE_CREDETENTIALS_FILE_NAME"]
	ProjectID := env["PROJECT_ID"]

	if c.ProjectID == "" {
		c.ProjectID = ProjectID
	}
	if c.ConnectTimeout == 0 {
		c.ConnectTimeout = time.Second * 10
	}
	if c.PublishTimeout == 0 {
		c.PublishTimeout = time.Second * 5
	}

	if c.ClientOptions == nil {
		c.ClientOptions = append(c.ClientOptions, option.WithCredentialsFile(cred))
	}
	c.DoNotCreateTopicIfMissing = true

}

// NewPublisher returns new publisher struct with publisher client
func NewPublisher() (*Publisher, error) {

	var config PublisherConfig

	//TO DO : import settings, otherwise use defaults
	config.setDefaults()

	pub := &Publisher{
		topics: map[string]*pubsub.Topic{},
		config: config,
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout)
	defer cancel()

	var err error
	pub.client, err = pubsub.NewClient(ctx, config.ProjectID, config.ClientOptions...)
	if err != nil {
		return nil, err
	}

	return pub, nil
}
