package repositories

import "cms_mail_service/model"

type Repository interface {
	FetchEmailLog() ([]model.EmailLog, error)
	SendMail(mailData model.MailInput) error
}