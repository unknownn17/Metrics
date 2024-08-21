package router

import (
	"conn/internal/connections"
	metrics17 "conn/internal/metrics"
	"expvar"
	"fmt"
	"log"
	"net/http"
)

func Router() {
	r := http.NewServeMux()
	handler := connections.NewHandler()
	r.HandleFunc("POST /books/add", handler.Create)
	r.HandleFunc("GET /books/{id}", handler.GetbyId)
	r.HandleFunc("GET /books", handler.Get)
	r.HandleFunc("PUT /books/{id}", handler.Update)
	r.HandleFunc("DELETE /books/{id}", handler.Delete)
	r.Handle("/debug/vars", expvar.Handler())
	r.Handle("/", metrics17.Metrics(http.HandlerFunc(handler.Check)))

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
