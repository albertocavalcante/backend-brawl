package api

import (
	"backend-brawl/src/db"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port  int
	Store *db.Store
}

func (s *Server) AddHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running...")
	})

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers := s.Store.GetCustomers()
		fmt.Fprintf(w, "Customers: %+v", customers)
	})
}

func (s *Server) ListenAndServe() {
	s.AddHandlers()

	addr := fmt.Sprintf(":%d", s.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
