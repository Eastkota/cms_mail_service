package model

type MultipleMailData struct{
	EmailLog []EmailLog `json:"email_log"`
	Error    *MailError `json:"error"`
}

type GenericMailResponse struct {
	Data  interface{}
	Error *MailError
}

type GenericMailSuccessData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
