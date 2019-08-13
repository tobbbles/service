package index

import "net/http"

var (
	Path    = "/"
	Methods = []string{"GET"}
)

type Endpoint struct{}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := []byte(`{"status": "ok"}`)

	w.Write(response)
}

func (e *Endpoint) Path() string {
	return Path
}

func (e *Endpoint) Methods() []string {
	return Methods
}
