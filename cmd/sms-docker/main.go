package main

import (
	"log"
	"net/http"

	"github.com/bakhtik/sms-docker/internal/app/sms-docker/views"
	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // "redis:6379" when in container
	})

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("web/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	mux.Handle("/", views.IndexHandler(redisClient))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
