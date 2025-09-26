package helpers

import "cms_mail_service/model"

func FormatError(err error) *model.GenericMailResponse {
	return &model.GenericMailResponse{
		Data: nil,
		Error: &model.MailError{
			Message: err.Error(),
		},
	}
}
