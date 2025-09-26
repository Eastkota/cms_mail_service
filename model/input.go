package model

type MailInput struct {
	To          string `json:"to"`
	Subject       string `json:"subject"`
	Title 		string `json:"title"`
	Content 	string `json:"content"`
}