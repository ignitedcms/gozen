package login

import (
	"fmt"
	"gozen/models/users"
	"gozen/system/formutils"
	"gozen/system/hash"
	"gozen/system/mail"
	"gozen/system/rendering"
	"gozen/system/validation"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := formutils.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	rendering.RenderTemplate(w, r, "login", data)
}

func ForgotView(w http.ResponseWriter, r *http.Request) {

	data := formutils.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	rendering.RenderTemplate(w, r, "forgot", data)
}

//we need to send an email reset if password is found

func Forgot(w http.ResponseWriter, r *http.Request) {

	validator := &validation.Validator{}

	validator.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email"))

	email := r.FormValue("email")

	postData := formutils.SetAndGetPostData(w, r)

	if validator.HasErrors() {
		formutils.HandleValidationErrors(w, r, validator, postData, "forgot")
		return
	}

	query, err := users.GetHash(email)
	//no record found
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(query.Email)

		to := query.Email
		mail.Test(to)
	}

	//http.Redirect(w, r, "/dashboard", http.StatusFound)
}

// post
func Login(w http.ResponseWriter, r *http.Request) {

	validator := &validation.Validator{}

	validator.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email")).
		Required("password", r.FormValue("password"))

		//email := r.FormValue("email")
	password := r.FormValue("password")
	email := r.FormValue("email")

	postData := formutils.SetAndGetPostData(w, r)

	if validator.HasErrors() {
		formutils.HandleValidationErrors(w, r, validator, postData, "login")
		return
	}

	//check password hash
	foo, err := users.GetHash(email)

	if err != nil {
		fmt.Print(err)
		formutils.HandleValidationErrors(w, r, validator, postData, "login")
		return
	}

	fmt.Print(foo.Password)

	t := hash.CheckPasswordHash(password, foo.Password)

	if !t {
		formutils.HandleValidationErrors(w, r, validator, postData, "login")
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func SignupView(w http.ResponseWriter, r *http.Request) {

	data := formutils.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	rendering.RenderTemplate(w, r, "sign-up", data)

}

// post
func Signup(w http.ResponseWriter, r *http.Request) {
	validator := &validation.Validator{}

	validator.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email")).
		Required("password", r.FormValue("password")).
		Required("name", r.FormValue("name"))

	email := r.FormValue("email")
	password := r.FormValue("password")
	name := r.FormValue("name")

	postData := formutils.SetAndGetPostData(w, r)

	if validator.HasErrors() {
		formutils.HandleValidationErrors(w, r, validator, postData, "sign-up")
		return
	}
	//else success
	//check if already in db

	p, _ := hash.HashPassword(password)

	t, _ := users.Create(name, email, p)
	fmt.Print(t)

	http.Redirect(w, r, "/dashboard", http.StatusFound)

}
