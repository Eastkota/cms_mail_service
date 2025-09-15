package services

import (
	"cms_mail_service/model"
	"cms_mail_service/repositories"
)

type EmailLogService struct{
	Repository repositories.Repository
}

func NewEmailLogService(repository repositories.Repository) *EmailLogService {
	return &EmailLogService{Repository: repository}
}

func (s *EmailLogService) FetchEmailLog() ([]model.EmailLog, error) {
    emails, err := s.Repository.FetchEmailLog()
    if err != nil {
        return nil, err
    }
    return emails, nil
}
