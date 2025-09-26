package repositories

import (
	"cms_mail_service/model"
	"gorm.io/gorm"
)

type EmailLogRepository struct {
	db *gorm.DB
}

func NewEmailLogRepository(db *gorm.DB) *EmailLogRepository {
	return &EmailLogRepository{db: db}
}

func (r *EmailLogRepository) FetchEmailLog() ([]model.EmailLog, error) {
    var emails []model.EmailLog
    err := r.db.Find(&emails).Error
    if err != nil {
        return nil, err
    }
    return emails, nil





    
}
