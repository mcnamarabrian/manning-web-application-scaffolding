package handlers

import "net/http"

func SetupHandlers(m *http.ServeMux) {
	m.HandleFunc("/", indexHandler)
	m.HandleFunc("/api", apiHandler)
	m.HandleFunc("/healthcheck", healthcheckHandler)

	// Original solution
	m.Handle("/static", http.FileServer(http.Dir("./static")))

	// Provided solution
	// fileServer := http.FileServer(http.Dir("./static"))
	// m.Handle("/static/", http.StripPrefix("/static/", fileServer))
}
