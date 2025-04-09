package application

import (
	"github.com/BrunoGuimaraesSilva/api.bruno-guimaraes.com/domain"
)

type SendMessageService struct {
	EmailRepo domain.EmailRepository
}

func NewSendMessageService(emailRepo domain.EmailRepository) *SendMessageService {
	return &SendMessageService{EmailRepo: emailRepo}
}

func (s *SendMessageService) SendMessageAdmin(msg domain.Message) (adminResp string, err error) {
	adminEmail, err := msg.PrepareAdminEmail()
	if err != nil {
		return "", err
	}

	adminResp, err = s.EmailRepo.Send(adminEmail)
	if err != nil {
		return "", err
	}

	return adminResp, nil
}

func (s *SendMessageService) SendMessageClient(msg domain.Message) (clientResp string, err error) {
	clientEmail, err := msg.PrepareClientEmail()
	if err != nil {
		return "", err
	}

	clientResp, err = s.EmailRepo.Send(clientEmail)
	if err != nil {
		return "", err
	}

	return clientResp, nil
}

func (s *SendMessageService) SendMessage(senderName, senderEmail, content string) (adminResp, clientResp string, err error) {
	msg, err := domain.NewMessage(senderName, senderEmail, content)
	if err != nil {
		return "", "", err
	}

	adminResp, err = s.SendMessageAdmin(*msg)
	if err != nil {
		return "", "", err
	}

	clientResp, err = s.SendMessageClient(*msg)
	if err != nil {
		return "", "", err
	}

	return adminResp, clientResp, nil
}
