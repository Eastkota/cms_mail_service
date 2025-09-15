package helpers

import "cms_mail_service/model"

func FormatError(err error) *model.EmailLogResponse {
	return &model.EmailLogResponse{
		Data: nil,
		Error: &model.MailError{
			Message: err.Error(),
		},
	}
}
