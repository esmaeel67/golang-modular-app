package jetstream

import (
	"context"

	"github.com/esmaeel67/golang-modular-app/internal/am"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

const maxRetries = 3

type Stream struct {
	streamName string
	js         nats.JetStreamContext
}

var _ am.MessageStream[am.RawMessage, am.RawMessage] = (*Stream)(nil)

func NewStream(streamName string, js nats.JetStreamContext) *Stream {
	return &Stream{
		streamName: streamName,
		js:         js,
	}
}

func (s *Stream) Publish(ctx context.Context, topicName string, rawMsg am.RawMessage) (err error) {
	var data []byte

	data, err = proto.Marshal(&StreamMessage{
		Id:   rawMsg.ID(),
		Name: rawMsg.MessageName(),
		Data: rawMsg.Data(),
	})

	if err != nil {
		return err
	}

	var p nats.PubAckFuture
	p, err = s.js.PublishMsgAsync(&nats.Msg{
		Subject: topicName,
		Data:    data,
	}, nats.MsgId(rawMsg.ID()))
	if err != nil {
		return err
	}

	// retry a handful of times to publish the messages
	go func(future nats.PubAckFuture, tries int) {
		var err error

		for {
			select {
			case <-future.Ok():
				return
			case <-future.Err(): // error ignored; try again
				//TODO: add some variable delay between tries
				tries = tries - 1
				if tries <= 0 {
					//TODO: do more then give up
					return
				}
				future, err = s.js.PublishMsgAsync(future.Msg())
				if err != nil {
					//TODO: do more than give up
					return
				}
			}
		}

	}(p, maxRetries)

	return nil

}
