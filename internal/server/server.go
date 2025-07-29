package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	s := &Server{
		httpServer: &http.Server{
			Addr: addr,
			Handler: mux,
		},
	}

	mux.HandleFunc("/send-email", s.handleSendEmail)
	return s
}

func (s *Server) Start() {
	log.Println("Server started on :", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Println("Stopping HTTP server...")
	s.httpServer.Shutdown(ctx)
}

func (s *Server) handleSendEmail(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	type emailJob struct {
		To string `json:"to"`
		Subject string `json:"subject"`
		Body string `json:"body"`
	}
	var job emailJob

	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Print("Received email job: ", job)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status":"enqueued"}`))
}