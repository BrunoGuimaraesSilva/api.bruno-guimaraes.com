package domain

type Email struct {
	From    string
	To      []string
	Subject string
	HTML    string
}
