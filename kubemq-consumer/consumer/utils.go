package consumer

import (
	"encoding/json"

	"github.com/kubemq-io/kubemq-go"
	// "github.com/opentracing/opentracing-go"

	"github.com/bygui86/go-traces/kubemq-consumer/logging"
)

func (c *KubemqConsumer) subscribeToEventStore() (chan error, <-chan *kubemq.EventStoreReceive, error) {
	// subscribe to event-store channel
	errChan := make(chan error, 1)
	msgChan, err := c.client.SubscribeToEventsStore(
		c.ctx,
		c.config.kubemqChannel,
		c.config.kubemqGroup,
		errChan,
		kubemq.StartFromNewEvents())
	if err != nil {
		return nil, nil, err
	}

	return errChan, msgChan, nil
}

func (c *KubemqConsumer) startConsumer(msgChan <-chan *kubemq.EventStoreReceive, errChan chan error) {
	for {
		select {
		case streamMsg := <-msgChan:

			// TODO extract span from tags?
			// streamMsg.Tags

			msg := &Message{}
			jsonErr := json.Unmarshal(streamMsg.Body, msg)
			if jsonErr != nil {
				logging.SugaredLog.Errorf("JSON unmarshal of message %+v failed: %s", streamMsg, jsonErr.Error())
			}

			logging.SugaredLog.Infof("Message received from channel %s: %+v", c.config.kubemqChannel, msg)

		case err := <-errChan:
			logging.SugaredLog.Errorf("Error received from channel %s: %s", c.config.kubemqChannel, err.Error())

		case <-c.ctx.Done():
			return
		}
	}
}
