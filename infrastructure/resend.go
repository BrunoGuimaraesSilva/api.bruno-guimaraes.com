package infrastructure

import (
	"fmt"

	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/domain"

	"github.com/resend/resend-go/v2"
)

type ResendEmailRepository struct {
	Client *resend.Client
}

func NewResendEmailRepository(apiKey string) *ResendEmailRepository {
	client := resend.NewClient(apiKey)
	return &ResendEmailRepository{
		Client: client,
	}
}

func (r *ResendEmailRepository) Send(email domain.Email) (interface{}, error) {
	params := &resend.SendEmailRequest{
		From:    email.From,
		To:      email.To,
		Subject: email.Subject,
		Html:    email.HTML,
	}

	response, err := r.Client.Emails.Send(params)
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}
	return response.Id, nil
}
