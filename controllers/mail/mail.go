package mail

import (
	//"fibs/system/formutils"
	"gozen/system/mail"
	"gozen/system/rendering"
	//"fibs/system/validation"
	"net/http"
	//"path/filepath"
)

// index page for mail
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "mail/index", nil)
}

func MailView(w http.ResponseWriter, r *http.Request) {
	rendering.RenderTemplate(w, r, "mail", nil)
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	to := "test@mail.com"
	mail.Test(to)

}
