/*                                                                          
|---------------------------------------------------------------            
| Mail test
|---------------------------------------------------------------            
|
| 
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/       
package mail

import (
	"gozen/system/mail"
	"gozen/system/templates"
	"gozen/system/validation"
	"net/http"
)

// index page for mail
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	templates.Render(w, r, "mail/index", nil)
}

func MailView(w http.ResponseWriter, r *http.Request) {
	data := templates.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	templates.Render(w, r, "mail", data)
}

func SendMail(w http.ResponseWriter, r *http.Request) {

	v := &validation.Validator{}

	v.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email"))

	//email := r.FormValue("email")
	email := r.FormValue("email")

	postData := templates.PostData(w, r)

	if v.HasErrors() {
		templates.Errors(w, r, v, postData, "mail")
		return
	}

	//to := query.Email
	mail.Test(email)
}
