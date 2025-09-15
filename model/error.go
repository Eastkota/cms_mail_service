package model

type MailError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field,omitempty"`
}
