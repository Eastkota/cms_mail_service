package model

type EmailLogResponse struct {
	Data  interface{}
	Error *MailError
}

type MultipleMailData struct{
	EmailLog []EmailLog `json:"email_log"`
	Error    *MailError `json:"error"`
}
