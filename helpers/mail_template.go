package helpers

import (
	"cms_mail_service/model"
	"html/template"

	"bytes"
	"log"
	"os"
	"path/filepath"
	"time"
)

func MailTemplate(mailDetail model.MailData) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	tmplPath := filepath.Join(cwd, "helpers", "mail_template.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	var buf bytes.Buffer

	data := struct {
		Title   string
		Content template.HTML
		Year    int
	}{
		Title:   mailDetail.Title,
		Content: template.HTML(mailDetail.Content),
		Year:    time.Now().Year(),
	}
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	return buf.String()
}
