package infrastructure

import (
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/domain"
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/errors"
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

func (r *ResendEmailRepository) Send(email domain.Email) (string, error) {
	params := &resend.SendEmailRequest{
		From:    email.From,
		To:      email.To,
		Subject: email.Subject,
		Html:    email.HTML,
	}

	response, err := r.Client.Emails.Send(params)
	if err != nil {
		return "", errors.FormatError(
			"Unable to send email at this time, please try again later",
			err,
			"Email Service > Send > Resend API",
		)
	}
	return response.Id, nil
}
