package repositories

import (
	"cms_mail_service/config"
	"cms_mail_service/helpers"
	"cms_mail_service/model"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-mail/mail"
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

func (r *EmailLogRepository) SendMail(mailData model.MailInput) error {
	err := helpers.ValidateEmail(mailData)
	if err != nil {
		return fmt.Errorf("validation failed: %v", err)
	}
	if mailData.To == "" {
		return fmt.Errorf("valid email adderss is required")
	}

	if mailData.Title == "" {
		return fmt.Errorf("title is required")
	}

	if mailData.Content == "" {
		return fmt.Errorf("content is required")
	}

	db, err := helpers.GetGormDB()
	if err != nil {
		return fmt.Errorf("failed to get DB connection: %v", err)
	}

	emailLog := model.EmailLog{
		Receiver: 		mailData.To,
		Subject:        mailData.Subject,
		Title: 			mailData.Title,
		Content:  		mailData.Content,
		Status:       	"queued",
		CreatedAt:    	time.Now(),
	}

	if err := db.Create(&emailLog).Error; err != nil {
		return fmt.Errorf("failed to create email log: %v", err)
	}

	m := mail.NewMessage()
	m.SetHeader("From", config.MailFromAddress())
	m.SetHeader("To", mailData.To)
	m.SetHeader("Subject", mailData.Subject)

	mailTemplateData := model.MailData{
		Title:      mailData.Title,
		Content:     mailData.Content,
	}

	m.SetBody("text/html", helpers.MailTemplate(mailTemplateData))

	d := mail.NewDialer(config.MailHost(), config.MailPort(), config.MailUsername(), config.MailPassword())
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		// Update email log with failed status
		db.Model(&emailLog).Updates(map[string]interface{}{
			"status":     "failed",
			"error_message":  err.Error(),
			"updated_at": time.Now(),
		})
		return fmt.Errorf("failed to send email: %v", err)
	}

	db.Model(&emailLog).Updates(map[string]interface{}{
		"status":     "sent",
		"updated_at": time.Now(),
	})

	return nil
}

