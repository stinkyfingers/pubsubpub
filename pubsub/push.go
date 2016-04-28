package pubsub

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/pubsub"

	"encoding/json"
	"io/ioutil"
)

var (
	ctx context.Context

	ProjectID    = "cp100-john"
	Subscription = "test"
)

// Pull gets a single message from the PUBSUB queue
func Push(topic string, data interface{}) error {
	var err error
	if ctx == nil {
		err = Context()
		if err != nil {
			return err
		}
	}
	exists, err := pubsub.TopicExists(ctx, topic)
	if !exists {
		err = pubsub.CreateTopic(ctx, topic)
	}
	if err != nil {
		return err
	}

	// build message
	var msg pubsub.Message
	msg.Data, err = json.Marshal(&data)
	if err != nil {
		return err
	}

	// publish
	_, err = pubsub.Publish(ctx, topic, &msg)
	return err
}

func Context() error {
	jsonKey, err := ioutil.ReadFile("keys/cp100-f39fd3c5c9f5.json")
	if err != nil {
		return err
	}
	conf, err := google.JWTConfigFromJSON(jsonKey, pubsub.ScopeCloudPlatform, pubsub.ScopePubSub)
	if err != nil {
		return err
	}
	ctx = cloud.NewContext(ProjectID, conf.Client(oauth2.NoContext))
	return nil
}
