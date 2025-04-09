package domain

import (
	"fmt"
	"os"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/domain/templates"
)

type Message struct {
	SenderName     string
	SenderEmail    EmailAddress
	Content        string
	AdminRecipient EmailAddress
}

func NewMessage(senderName string, senderEmail string, content string) (*Message, error) {
	AdminEmail := os.Getenv("ADMIN_EMAIL")
	if AdminEmail == "" {
		AdminEmail = "no-reply@bruno-guimaraes.com"
	}

	if senderName == "" {
		return nil, fmt.Errorf("sender name cannot be empty")
	}
	if content == "" {
		return nil, fmt.Errorf("message cannot be empty")
	}

	emailAddr, err := NewEmailAddress(senderEmail)
	if err != nil {
		return nil, err
	}

	return &Message{
		SenderName:     senderName,
		SenderEmail:    emailAddr,
		Content:        content,
		AdminRecipient: EmailAddress(AdminEmail),
	}, nil
}

func (m *Message) PrepareAdminEmail() (Email, error) {
	html, err := templates.RenderAdminEmail(templates.AdminEmailData{
		Name:    m.SenderName,
		Email:   string(m.SenderEmail),
		Message: m.Content,
	})
	if err != nil {
		return Email{}, err
	}

	return Email{
		From:    "bruno-guimaraes.com <no-reply@bruno-guimaraes.com>",
		To:      []string{string(m.AdminRecipient)},
		Subject: fmt.Sprintf("Message from %s", m.SenderName),
		HTML:    html,
	}, nil
}

func (m *Message) PrepareClientEmail() (Email, error) {
	html, err := templates.RenderClientEmail(templates.ClientEmailData{
		Name:    m.SenderName,
		Message: m.Content,
	})
	if err != nil {
		return Email{}, err
	}

	return Email{
		From:    "bruno-guimaraes.com <no-reply@bruno-guimaraes.com>",
		To:      []string{string(m.SenderEmail)},
		Subject: fmt.Sprintf("We have received your message, %s", m.SenderName),
		HTML:    html,
	}, nil
}
