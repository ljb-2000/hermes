package main

import (
	"fmt"
	"net/http"
)

// RunServer starts a web server on port 8080, which serves the
// Hermes web interface
func RunWebServer(supervisor *AgentSupervisor) {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		NewRulesPage().Render(w)
	})

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
