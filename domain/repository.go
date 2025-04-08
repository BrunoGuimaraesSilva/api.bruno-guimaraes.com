package domain

type EmailRepository interface {
	Send(email Email) (interface{}, error)
}
