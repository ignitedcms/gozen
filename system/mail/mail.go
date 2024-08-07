/*
|---------------------------------------------------------------
| Mail
|---------------------------------------------------------------
|
| Mail utility helper file for easily configuring your email
| settings
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package mail

import (
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
	smtpHost       string
	smtpPort       string
	senderEmail    string
	senderPassword string
	recipientEmail string
	templatePath   string
	anchor         string
	htmlContent    []byte
	message        []byte
	auth           smtp.Auth
}

func New() *Mail {
	m := &Mail{}
	m.setupFromEnv()
	return m
}

func (m *Mail) setupFromEnv() {
	m.smtpHost = os.Getenv("MAIL_HOST")
	m.smtpPort = os.Getenv("MAIL_PORT")
	m.senderEmail = os.Getenv("MAIL_USERNAME")
	m.senderPassword = os.Getenv("MAIL_PASSWORD")
	m.auth = smtp.PlainAuth("", m.senderEmail, m.senderPassword, m.smtpHost)
}

func (m *Mail) SetAnchor(anchor string) *Mail {
	m.anchor = anchor
	return m
}

func (m *Mail) SetRecipient(email string) *Mail {
	m.recipientEmail = email
	return m
}

func (m *Mail) SetTemplatePath(path string) *Mail {
	m.templatePath = path
	return m
}

func (m *Mail) LoadTemplate() *Mail {
	htmlContent, err := ioutil.ReadFile(m.templatePath)
	if err != nil {
		log.Fatal("Error reading HTML file:", err)
		return m
	}

	// Convert the byte slice to a string
	text := string(htmlContent)

	// Perform the string replacement
	text = strings.ReplaceAll(text, "{{title}}", "Password Reset")
	text = strings.ReplaceAll(text, "{{body}}", "Please click on the button to reset your password.")
	text = strings.ReplaceAll(text, "{{anchor}}", m.anchor)
	text = strings.ReplaceAll(text, "{{anchortext}}", "Click here")

	//we need to convert back to bytes for it to work
	replacedBytes := []byte(text)

	m.htmlContent = replacedBytes
	return m
}

func (m *Mail) BuildMessage() *Mail {
	m.message = []byte("From: " + m.senderEmail + "\r\n" +
		"To: " + m.recipientEmail + "\r\n" +
		"Subject: Test Email\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		string(m.htmlContent))
	return m
}

func (m *Mail) Send() []byte {
	err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, m.auth, m.senderEmail, []string{m.recipientEmail}, m.message)
	if err != nil {
		log.Println("Error sending email:", err)
		return []byte(err.Error())
	}
	log.Println("Email sent successfully!")
	return []byte("Email sent successfully!")
}
