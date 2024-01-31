package presentation

import (
	"io"
	"net/http"
	_ "net/http/pprof"
)

// ProbeAPI is a health check API for the application
type ProbeAPI struct {
}

// NewProbeAPI creates a new instance of ProbeAPI
func NewProbeAPI() *ProbeAPI {
	api := &ProbeAPI{}
	http.HandleFunc("/health", api.Health)
	return api
}

// Serve starts the API server
func (api *ProbeAPI) Serve(port string) {
	http.ListenAndServe(":"+port, nil)
}

// Health is the health check handler
func (api *ProbeAPI) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}
