package errors

import (
	"fmt"
	"os"
)

var environment string

func init() {
	environment = os.Getenv("GO_ENV")
	if environment == "" {
		environment = "development"
	}
}

type ProductionError struct {
	Message string
}

type DevelopmentError struct {
	UserMessage   string
	TechnicalInfo error
	StackTrace    string
}

func (e *ProductionError) Error() string {
	return e.Message
}

func (e *DevelopmentError) Error() string {
	return fmt.Sprintf("User Message: %s\nTechnical Details: %v\nStack Trace: %s",
		e.UserMessage,
		e.TechnicalInfo,
		e.StackTrace)
}

func FormatError(userMessage string, technicalError error, stackTrace string) error {
	if environment == "production" {
		return &ProductionError{
			Message: userMessage,
		}
	}

	return &DevelopmentError{
		UserMessage:   userMessage,
		TechnicalInfo: technicalError,
		StackTrace:    stackTrace,
	}
}

func IsProduction() bool {
	return environment == "production"
}

func GetEnvironment() string {
	return environment
}
