package mailer

import (
	"context"
	"fmt"

	mailerv1 "github.com/anuragkumar19/connect/events/gen/mailer/v1"
	"github.com/anuragkumar19/connect/infra/nats"
	"github.com/nats-io/nats.go/jetstream"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	js jetstream.JetStream
}

func NewClient(nc *nats.NATS) (Client, error) {
	js, err := jetstream.New(nc.Conn)
	if err != nil {
		return Client{}, fmt.Errorf("failed to create jetstream client from nats.Conn: %w", err)
	}

	return Client{
		js: js,
	}, nil
}

func (c *Client) SendEmail(ctx context.Context, msg *mailerv1.SendEmail) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal msg: %w", err)
	}
	if _, err := c.js.Publish(ctx, sendEmailSubject, b); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}
