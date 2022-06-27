package fabric

import (
	"math/rand"
	"net/url"
)

type Header map[string]string

type Request struct {
	Exchange string

	Topic string

	RoutingKey string

	// URL specifies either the URI being requested (for server
	// requests) or the URL to access (for client requests).
	URL *url.URL

	// Header contains the request header fields either received
	// by the server or to be sent by the client.
	Header Header

	CorrelationID string

	ContentType string

	// Body is the request's body.
	Body []byte
}

func NewRequest(exchange, uri, topic, contentType string, body []byte) (*Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	req := &Request{
		Exchange:      exchange,
		Topic:         topic,
		URL:           u,
		CorrelationID: randomString(32),
		ContentType:   contentType,
		Body:          body,
	}
	return req, nil
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
