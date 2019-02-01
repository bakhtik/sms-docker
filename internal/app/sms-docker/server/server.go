package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bakhtik/sms-docker/internal/app/sms-docker/model"
	"github.com/bakhtik/sms-docker/internal/app/sms-docker/views"
)

// ServerConfig stores the hostname and port number
type ServerConfig struct {
	Hostname string `json:"Hostname"` // Server name
	HTTPPort int    `json:"HTTPPort"` // HTTP port
}

func Run(s ServerConfig, cache model.Cache) {

	mux := http.NewServeMux()

	handlers(mux, cache)

	server := &http.Server{
		Addr:    httpAddress(s),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

// httpAddress returns the HTTP address
func httpAddress(s ServerConfig) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}

func handlers(mux *http.ServeMux, cache model.Cache) {
	files := http.FileServer(http.Dir("web/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/", views.IndexHandler(cache))
}
