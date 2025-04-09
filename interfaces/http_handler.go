package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/application"
)

type EmailRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type EmailResponse struct {
	AdminEmail  string `json:"adminEmail,omitempty"`
	ClientEmail string `json:"clientEmail,omitempty"`
	Message     string `json:"message,omitempty"`
}

type MessageHandler struct {
	Service *application.SendMessageService
}

func NewMessageHandler(service *application.SendMessageService) *MessageHandler {
	return &MessageHandler{Service: service}
}

func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"message": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var reqBody EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"message": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	adminResp, clientResp, err := h.Service.SendMessage(reqBody.Name, reqBody.Email, reqBody.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"message": "%v"}`, err), http.StatusBadRequest)
		return
	}

	resp := EmailResponse{
		AdminEmail:  adminResp,
		ClientEmail: clientResp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
