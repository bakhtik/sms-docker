package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	var visitsErr error
	visits, err := getVisitsCount()
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
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// handling transient errors
func getVisitsCount() (visits int64, err error) {
	for retry := 0; retry < 5; retry++ {
		visits, err = redisClient.Incr("counter").Result()
		if err == nil {
			return
		}
	}
	return 0, err
}

func Foo() string {
	return "hola!"
}
