package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/index", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	visits, err := redisClient.Incr("counter").Result()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Visits int64
		Error  error
	}{
		visits,
		err,
	}
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
