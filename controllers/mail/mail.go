package mail

import (
	//"fibs/system/formutils"
	"gozen/system/rendering"
	"io/ioutil"
	//"fibs/system/validation"
	"log"
	"net/http"
	"net/smtp"
	//"path/filepath"
)

// index page for mail
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "mail/index", nil)
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	// SMTP server configuration
	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := "25" // or 465 for SSL/TLS

	// Sender email authentication
	senderEmail := ""
	senderPassword := ""

	// Recipient email address
	recipientEmail := "foo@mail.com"

	// Path to HTML template file
	templatePath := "mail/email_template.html"

	// Read HTML content from file
	htmlContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatal("Error reading HTML file:", err)
		return
	}

	// Message contenti
	message := []byte("From: " + senderEmail + "\r\n" +
		"To: " + recipientEmail + "\r\n" +
		"Subject: Test Email\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		string(htmlContent))

	// SMTP authentication
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// Send email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		log.Println("Error sending email:", err)
	} else {
		log.Println("Email sent successfully!")
	}
}
