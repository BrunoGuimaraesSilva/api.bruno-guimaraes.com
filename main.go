package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/application"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/infrastructure"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/interfaces"
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

	http.HandleFunc("/api/send-message", interfaces.AuthMiddleware(authToken, handler.SendMessage))
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
