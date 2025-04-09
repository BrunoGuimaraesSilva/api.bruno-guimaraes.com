package domain

import (
	"fmt"
	"regexp"
	"strings"
)

type EmailAddress string

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

const (
	maxLocalPartLength = 64
	maxDomainLength    = 255
	maxTotalLength     = 254
)

func NewEmailAddress(value string) (EmailAddress, error) {
	if value == "" {
		return "", fmt.Errorf("email address cannot be empty")
	}

	email := strings.TrimSpace(value)

	if len(email) > maxTotalLength {
		return "", fmt.Errorf("email address is too long (maximum %d characters)", maxTotalLength)
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email format: must contain exactly one @ symbol")
	}

	localPart := parts[0]
	domain := parts[1]

	if len(localPart) > maxLocalPartLength {
		return "", fmt.Errorf("local part of email is too long (maximum %d characters)", maxLocalPartLength)
	}
	if len(domain) > maxDomainLength {
		return "", fmt.Errorf("domain part of email is too long (maximum %d characters)", maxDomainLength)
	}

	if localPart == "" || domain == "" {
		return "", fmt.Errorf("email address cannot have empty local part or domain")
	}

	if !emailRegex.MatchString(email) {
		return "", fmt.Errorf("invalid email format")
	}

	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
		return "", fmt.Errorf("domain cannot start or end with a dot")
	}
	if strings.Contains(domain, "..") {
		return "", fmt.Errorf("domain cannot contain consecutive dots")
	}

	return EmailAddress(email), nil
}
