package invoicer

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"
	"net/smtp"
	"net/textproto"
	"reflect"
	"testing"
)

type emailRecorder struct {
	addr string
	auth smtp.Auth
	from string
	to   []string
	msg  []byte
}
func TestCreateAttachment(t *testing.T) {
	fileName := "invoice.pdf"
	content := []byte{}
	contentType := "application/pdf"
	expected := &email.Attachment{ Filename: fileName, Header: textproto.MIMEHeader{}, Content: content }
	expected.Header.Set("Content-Type", contentType)
	expected.Header.Set("Content-Disposition", fmt.Sprintf("attachment;\r\n filename=\"%s\"", fileName))
	expected.Header.Set("Content-ID", fmt.Sprintf("<%s>", fileName))
	expected.Header.Set("Content-Transfer-Encoding", "base64")
	result := CreateAttachment(&content, &fileName, contentType)
	assert.True(t, reflect.DeepEqual(result, expected), "Should create a new email attachment, based on an array of bytes.")

	contentType = ""
	result = CreateAttachment(&content, &fileName, contentType)
	assert.True(t, reflect.DeepEqual(result, expected), "Should create a new email attachment, with empty contentType.")

	contentType = "application/pdf"
	fileName = ""
	result = CreateAttachment(&content, &fileName, contentType)
	assert.True(t, reflect.DeepEqual(result, expected), "Should create a new email attachment, with empty fileName.")
	assert.True(t, fileName == "invoice.pdf", "Should rename file to 'invoice.pdf' when fileName is empty.")
}

func TestEmailSender(t *testing.T) {
	f, r := mockSend(nil)
	sender := &emailSender{send: f}
	body := "Hello World"
	auth := smtp.PlainAuth("","sender@in-voicer.com", "password123", "smtp.in-voicer.com")
	err := sender.send("info@in-voicer.com", auth, "I am Invoicer", []string{"juanca@in-voicer.com"}, []byte(body))

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	assert.True(t, string(r.msg) == body, "Should send an email.")
}

func mockSend(errToReturn error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	r := new(emailRecorder)
	return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		*r = emailRecorder{addr, a, from, to, msg}
		return errToReturn
	}, r
}
