package domain

type EmailRepository interface {
	Send(email Email) (string, error)
}
