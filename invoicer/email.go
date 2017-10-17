package invoicer

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"net/textproto"
	"os"
	"strings"
)

type EmailConfig struct {
	Username   	string
	Password   	string
	ServerHost 	string
	ServerPort 	string
	FromEmail		string
	FromName		string
}

type EmailSender interface {
	Send(to []string, subject string, body []byte, attachment *email.Attachment) error
}

func NewEmailSender(conf EmailConfig) EmailSender {
	return &emailSender{conf, smtp.SendMail}
}

type emailSender struct {
	conf EmailConfig
	send func(string, smtp.Auth, string, []string, []byte) error
}

func (e *emailSender) Send(to []string, subject string, body []byte, attachment *email.Attachment) error {
	email := email.NewEmail()
	email.From = strings.Join([]string{e.conf.FromName, " <", e.conf.FromEmail, ">"}, "")
	email.To = to
	email.Subject = subject
	email.HTML = body
	if (attachment != nil) {
		email.Attachments = append(email.Attachments, attachment)
	}
	return email.SendWithTLS(
		e.conf.ServerHost + ":" + e.conf.ServerPort,
		smtp.PlainAuth("", e.conf.Username, e.conf.Password, e.conf.ServerHost),
		&tls.Config{InsecureSkipVerify: true},
	)
}

func CreateAttachment(content *[]byte, fileName *string, contentType string) (a *email.Attachment) {
	if (*fileName == "") {
		*fileName = "invoice.pdf"
	}
	at := &email.Attachment{
		Filename: *fileName,
		Header:   textproto.MIMEHeader{},
		Content:  *content,
	}

	// Get the Content-Type to be used in the MIMEHeader
	if contentType != "" {
		at.Header.Set("Content-Type", contentType)
	} else {
		// If the Content-Type is blank, set the Content-Type to "application/pdf"
		at.Header.Set("Content-Type", "application/pdf")
	}
	at.Header.Set("Content-Disposition", fmt.Sprintf("attachment;\r\n filename=\"%s\"", *fileName))
	at.Header.Set("Content-ID", fmt.Sprintf("<%s>", *fileName))
	at.Header.Set("Content-Transfer-Encoding", "base64")
	return at
}

func SendEmail(emailInfo *EmailInfo, attachment *email.Attachment) error {
	// STEP 1 - Parse email body
	data := struct{
    FromName	string
		Message		template.HTML
	}{
		emailInfo.FromName,
		template.HTML(emailInfo.Message),
	}
	parsedHtmlEmailBody, err := parseTemplateToString("/invoicer/templates/template_email.html", data)

	// STEP 2 - Send email
	if (err == nil) {
		emailConfig := &EmailConfig{
			FromEmail: emailInfo.FromEmail,
			FromName: emailInfo.FromName,
			Username: "in.voicer.robot@gmail.com",
			Password: os.Getenv("INVOICER_EMAIL_PASSWORD"),
			ServerHost: "smtp.gmail.com",
			ServerPort: "587",
		}
		sender := NewEmailSender(*emailConfig)
		err := sender.Send(emailInfo.ToEmails, emailInfo.Subject, []byte(parsedHtmlEmailBody), attachment)
		if (err != nil) {
			return err
		}
		return nil
	}
	return err
}