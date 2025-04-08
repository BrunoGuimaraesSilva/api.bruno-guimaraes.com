package domain

import (
	"fmt"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/domain/templates"
)

// EmailAddress is a Value Object
type EmailAddress string

func NewEmailAddress(value string) (EmailAddress, error) {
	if value == "" {
		return "", fmt.Errorf("email address cannot be empty")
	}
	return EmailAddress(value), nil
}

// Message is an Entity and Aggregate Root
type Message struct {
	SenderName     string
	SenderEmail    EmailAddress
	Content        string
	AdminRecipient EmailAddress
}

func NewMessage(senderName string, senderEmail EmailAddress, content string) (*Message, error) {
	if senderName == "" {
		return nil, fmt.Errorf("sender name cannot be empty")
	}
	if content == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}
	return &Message{
		SenderName:     senderName,
		SenderEmail:    senderEmail,
		Content:        content,
		AdminRecipient: EmailAddress("bruno.sil16441@gmail.com"),
	}, nil
}

// PrepareAdminEmail constructs the admin notification using the template
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

// PrepareClientEmail constructs the client confirmation using the template
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
