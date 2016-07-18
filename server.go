package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	cacheSince   = time.Now().Format(http.TimeFormat)
	cacheExpires = time.Now().AddDate(60, 0, 0).Format(http.TimeFormat)
)

func loadFavicon() ([]byte, error) {
	faviconReader, err := os.Open("static/images/favicon.ico")
	if err != nil {
		return nil, err
	}
	favicon, err := ioutil.ReadAll(faviconReader)
	if err != nil {
		return nil, err
	}
	return favicon, nil
}

func setFarFutureExpiresHeader(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "max-age:290304000, public")
	w.Header().Set("Last-Modified", cacheSince)
	w.Header().Set("Expires", cacheExpires)
}

// RunWebServer starts a web server on port 8080, which serves the
// Hermes web interface
func RunWebServer(supervisor *AgentSupervisor) {
	fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		setFarFutureExpiresHeader(w)
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
	})

	favicon, err := loadFavicon()
	if err != nil {
		log.Fatal("could not load favicon:", err)
	}
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		setFarFutureExpiresHeader(w)
		w.Write(favicon)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		NewAgentsPage(supervisor).Render(w)
	})

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
