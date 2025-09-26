package helpers

import (
	"cms_mail_service/model"
	"fmt"
)

func ValidateEmail(inputData model.MailInput) error {
	if inputData.Subject == "" {
		return fmt.Errorf("subject type is required")
	}
	if inputData.Content == "" {
		return fmt.Errorf("content is required")
	}
	if inputData.To == "" {
		return fmt.Errorf("email is required")
	}
	if inputData.Title == "" {
		return fmt.Errorf("title is required")
	}

	return nil
}

