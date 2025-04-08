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

func (s *SendMessageService) SendMessage(senderName, senderEmail, content string) (adminResp, clientResp interface{}, err error) {
	emailAddr, err := domain.NewEmailAddress(senderEmail)
	if err != nil {
		return nil, nil, err
	}

	msg, err := domain.NewMessage(senderName, emailAddr, content)
	if err != nil {
		return nil, nil, err
	}

	// Prepare and send admin email
	adminEmail, err := msg.PrepareAdminEmail()
	if err != nil {
		return nil, nil, err
	}
	adminResp, err = s.EmailRepo.Send(adminEmail)
	if err != nil {
		return nil, nil, err
	}

	// Prepare and send client email
	clientEmail, err := msg.PrepareClientEmail()
	if err != nil {
		return nil, nil, err
	}
	clientResp, err = s.EmailRepo.Send(clientEmail)
	if err != nil {
		return nil, nil, err
	}

	return adminResp, clientResp, nil
}
