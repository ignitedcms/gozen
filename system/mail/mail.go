package mail

import (
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

func Test(to string) {

	// SMTP server configuration
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")

	// Sender email authentication
	senderEmail := os.Getenv("MAIL_USERNAME")
	senderPassword := os.Getenv("MAIL_PASSWORD")

	// Recipient email address
	// Obtained from form
	recipientEmail := "test@mail.com"

	// Path to HTML template file
	templatePath := "mail/email_template.html"

	// Read HTML content from file
	htmlContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatal("Error reading HTML file:", err)
		return
	}

	// Message content
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
		//w.Write([]byte("Error sending email, check credentials:"))
		log.Println("Error sending email:", err)
	} else {
		//w.Write([]byte("Email sent successfully!"))
		log.Println("Email sent successfully!")
	}
}
