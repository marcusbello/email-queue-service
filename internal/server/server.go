package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/marcusbello/email-queue-service/internal/email"
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


	var job email.EmailJob

	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !validateEmailJob(job) {
		http.Error(w, "Invalid input", http.StatusUnprocessableEntity)
		return
	}

	log.Print("Received email job: ", job)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status":"enqueued"}`))
}

func validateEmailJob(j email.EmailJob) bool {
	if j.To == "" || j.Subject == "" || j.Body == "" {
		return false
	}
	match, _ := regexp.MatchString(`^[^@\s]+@[^@\s]+\.[^@\s]+$`, j.To)
	return match
}