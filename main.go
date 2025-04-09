package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/application"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/infrastructure"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/interfaces"
	"github.com/rs/cors"
)

func main() {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		fmt.Println("RESEND_API_KEY not set")
		return
	}

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		fmt.Println("AUTH_TOKEN not set")
		return
	}

	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		fmt.Println("CORS_ORIGIN not set")
		return
	}

	emailRepo := infrastructure.NewResendEmailRepository(apiKey)
	service := application.NewSendMessageService(emailRepo)
	handler := interfaces.NewMessageHandler(service)

	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(corsOrigin, ";"),
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	http.Handle("/api/send-message", c.Handler(interfaces.AuthMiddleware(authToken, handler.SendMessage)))
	http.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Welcome to api.bruno-guimaraes.com"}`))
	}))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/api" || path == "/api/send-message" || strings.HasPrefix(path, "/api/") {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/api", http.StatusPermanentRedirect)
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
