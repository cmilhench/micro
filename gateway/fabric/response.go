package fabric

type Response struct {
	// Header maps header keys to values.
	Header Header

	// Body represents the response body.
	Body []byte

	// Request is the request that was sent to obtain this Response.
	Request *Request
}
