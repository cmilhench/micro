package fabric

import (
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type Client struct {
	// Timeout specifies a time limit for requests made by this Client.
	Timeout    time.Duration
	Name       string
	Properties map[string]string
}

var DefaultClient = &Client{
	Timeout: time.Second * 30,
	Name:    "https://github.com/cmilhench/micro/client",
}

func Publish(exchange, url, topic, contentType string, body []byte) (resp *Response, err error) {
	return DefaultClient.Publish(exchange, url, topic, contentType, body)
}

func (c *Client) Publish(exchange, url, topic, contentType string, body []byte) (resp *Response, err error) {
	req, err := NewRequest(exchange, url, topic, contentType, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Do(req *Request) (*Response, error) {
	timeout := time.After(c.Timeout)
	// Connect to the server.
	cfg := amqp.Config{
		Dial:       amqp.DefaultDial(time.Second * 10),
		Heartbeat:  60 * time.Second,
		Locale:     "en_US",
		Properties: amqp.Table{},
	}
	cfg.Properties["connection_name"] = c.Name
	for k, v := range c.Properties {
		cfg.Properties[k] = v
	}

	conn, err := amqp.DialConfig(req.URL.String(), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to dial amqp, %s", err)
	}
	notifyClose := make(chan *amqp.Error)
	conn.NotifyClose(notifyClose)
	defer conn.Close()

	// Create a channel.
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel, %s", err)
	}
	defer ch.Close()

	// Declare the exchange.
	err = ch.ExchangeDeclare(
		req.Exchange, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted when usused
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange, %s", err)
	}

	// Declare the reply queue.
	replyQ, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // auto-delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare reply queue, %s", err)
	}

	// Listen for the reply.
	replies, err := ch.Consume(
		replyQ.Name, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register reply consumer, %s", err)
	}

	// Publish the message.
	msg := amqp.Publishing{
		ContentType:   req.ContentType,
		CorrelationId: req.CorrelationID,
		ReplyTo:       replyQ.Name,
		Body:          req.Body,
	}
	for k, v := range req.Header {
		msg.Headers[k] = v
	}
	err = ch.Publish(
		req.Exchange, // exchange
		req.Topic,    // routing key
		false,        // mandatory
		false,        // immediate
		msg)
	if err != nil {
		return nil, fmt.Errorf("failed to publish a message, %s", err)
	}

	// Wait for the reply.
	for {
		select {
		case err = <-notifyClose:
			return nil, fmt.Errorf("connection closed, %s", err)
		case <-timeout:
			return nil, errors.New("timeout") // i.e. timeout 504
		case msg := <-replies:
			if req.CorrelationID == msg.CorrelationId {
				header := make(map[string]string)
				for k, v := range msg.Headers {
					header[k], _ = v.(string)
				}
				return &Response{Header: header, Body: msg.Body, Request: req}, nil
			}
		}
	}
}
