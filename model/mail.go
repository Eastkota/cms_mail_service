package model

import (
	"time"
	"github.com/google/uuid"
)

type EmailLog struct {
    ID          	uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Receiver    	string    	`gorm:"type:varchar(255); not null" json:"receiver"`
    Subject         string     	`gorm:"type:varchar(10); not null" json:"subject"`
    Title  	        string     	`gorm:"type:varchar(20); not null" json:"title"` 
    Content         string     	`gorm:"type:text; not null" json:"content"`
    Status       	string     	`gorm:"type:status_enum;default:'queued'" json:"status"` 
    ErrorMsg     	*string    	`gorm:"column:error_message" json:"error_message"`
    CreatedAt   	time.Time  	`gorm:"type:timestamptz;not null" json:"created_at"`
    UpdatedAt   	time.Time  	`gorm:"type:timestamptz;not null" json:"updated_at"`
}

func (EmailLog) TableName() string {
    return "email_log.email_logs"
}

type MailData struct {
	Title       string `json:"title"`
	Content     string `jason:"content"`
}
