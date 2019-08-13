package server

import "net/http"

// Endpoint is an interface that can be fulfilled by structs providing request/response handling.
// This interface embeds the http.Handler interface, allowing us to pass an Endpoint as a http handler
// to functions such as mux handlers.
type Endpoint interface {
	Path() string
	Methods() []string

	http.Handler
}
