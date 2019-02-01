package views

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func IndexHandler(redisClient *redis.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var visitsErr error
		visits, err := getVisitsCount(redisClient)
		if err != nil {
			log.Println(err)
			visitsErr = err
		}

		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}

		data := struct {
			Hostname string
			Visits   int64
			Error    error
		}{
			hostname,
			visits,
			visitsErr,
		}

		log.Println(r.RemoteAddr, r.RequestURI, r.Referer())

		generateHTML(w, data, "layout", "navbar", "index")
	})
}

// handling transient errors
func getVisitsCount(redisClient *redis.Client) (visits int64, err error) {
	for retry := 0; retry < 5; retry++ {
		visits, err = redisClient.Incr("counter").Result()
		if err == nil {
			return
		}
	}
	return 0, err
}
