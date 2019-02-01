package server

import (
	"log"
	"net/http"

	"github.com/bakhtik/sms-docker/internal/app/sms-docker/model"
	"github.com/bakhtik/sms-docker/internal/app/sms-docker/views"
	"github.com/go-redis/redis"
)

// Server stores the hostname and port number
type Server struct {
	Hostname string `json:"Hostname"` // Server name
	HTTPPort int    `json:"HTTPPort"` // HTTP port
}

func Run() {

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: "redis:6379", // "redis:6379" when in container
	// })

	redisCache := &model.RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: "redis:6379", // "redis:6379" when in container
		}),
	}

	mux := http.NewServeMux()

	handlers(mux, redisCache)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func handlers(mux *http.ServeMux, cache model.Cache) {
	files := http.FileServer(http.Dir("web/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/", views.IndexHandler(cache))
}
