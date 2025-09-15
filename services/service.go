package services

import "cms_mail_service/model"

type Services interface {
	FetchEmailLog() ([]model.EmailLog, error)
}
