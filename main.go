package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/application"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/infrastructure"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/interfaces"
	"github.com/rs/cors" // Add this import
)

func main() {
	// Load environment variables
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

	emailRepo := infrastructure.NewResendEmailRepository(apiKey)
	service := application.NewSendMessageService(emailRepo)
	handler := interfaces.NewMessageHandler(service)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://bruno-guimaraes.com"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	http.Handle("/api/send-message", c.Handler(interfaces.AuthMiddleware(authToken, handler.SendMessage)))

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
