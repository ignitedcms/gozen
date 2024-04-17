package mail

import (
	"fmt"
	"gozen/models/users"
   "gozen/system/mail"
	"gozen/system/rendering"
	"gozen/system/formutils"
	"gozen/system/validation"
	"net/http"
)

// index page for mail
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "mail/index", nil)
}

func MailView(w http.ResponseWriter, r *http.Request) {
	data := formutils.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	rendering.RenderTemplate(w, r, "mail", data)
}

func SendMail(w http.ResponseWriter, r *http.Request) {

	validator := &validation.Validator{}

	validator.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email"))

	//email := r.FormValue("email")
	email := r.FormValue("email")

	postData := formutils.SetAndGetPostData(w, r)

	if validator.HasErrors() {
		formutils.HandleValidationErrors(w, r, validator, postData, "mail")
		return
	}

	query, err := users.GetHash(email)
   //no record found
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(query.Email)

      //to := query.Email
      mail.Test(email)
	}
}
