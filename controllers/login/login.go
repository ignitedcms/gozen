package login

import (
	"fmt"
	"gozen/models/users"
	"gozen/system/hash"
	"gozen/system/mail"
	"gozen/system/templates"
	"gozen/system/validation"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := templates.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	templates.Render(w, r, "login", data)
}

func ForgotView(w http.ResponseWriter, r *http.Request) {

	data := templates.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	templates.Render(w, r, "forgot", data)
}

//we need to send an email reset if password is found

func Forgot(w http.ResponseWriter, r *http.Request) {

	v := &validation.Validator{}

	v.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email")).
		Exists("email", r.FormValue("email"), "users", "email")

	email := r.FormValue("email")

	postData := templates.PostData(w, r)

	if v.HasErrors() {
		templates.Errors(w, r, v, postData, "forgot")
		return
	}

	//Let's set the token and update the db

	rand := hash.RandomString()
	fmt.Print(rand)
	t := users.SetToken(rand, email)
	fmt.Print(t)

	recipientEmail := email
	templatePath := "mail/email_template.html"

	result := mail.New().
		SetRecipient(recipientEmail).
		SetTemplatePath(templatePath).
		LoadTemplate().
		BuildMessage().
		Send()

	w.Write(result)

	//http.Redirect(w, r, "/dashboard", http.StatusFound)
}

// post
func Login(w http.ResponseWriter, r *http.Request) {

	v := &validation.Validator{}

	v.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email")).
		Exists("email", r.FormValue("email"), "users", "email").
		Required("password", r.FormValue("password"))

		//email := r.FormValue("email")
	password := r.FormValue("password")
	email := r.FormValue("email")

	postData := templates.PostData(w, r)

	if v.HasErrors() {
		templates.Errors(w, r, v, postData, "login")
		return
	}

	//check password hash
	foo, err := users.GetHash(email)

	if err != nil {
		fmt.Print(err)
		templates.Errors(w, r, v, postData, "login")
		return
	}

	fmt.Print(foo.Password)

	t := hash.CheckPasswordHash(password, foo.Password)

	if !t {
		templates.Errors(w, r, v, postData, "login")
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func SignupView(w http.ResponseWriter, r *http.Request) {

	data := templates.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	templates.Render(w, r, "sign-up", data)

}

// post
func Signup(w http.ResponseWriter, r *http.Request) {
	v := &validation.Validator{}

	v.Required("email", r.FormValue("email")).
		Email("email", r.FormValue("email")).
		Unique("email", r.FormValue("email"), "users", "email").
		Required("password", r.FormValue("password")).
		Required("name", r.FormValue("name"))

	email := r.FormValue("email")
	password := r.FormValue("password")
	name := r.FormValue("name")

	postData := templates.PostData(w, r)

	if v.HasErrors() {
		templates.Errors(w, r, v, postData, "sign-up")
		return
	}
	//else success

	p, _ := hash.HashPassword(password)

	t, _ := users.Create(name, email, p)
	fmt.Print(t)

	http.Redirect(w, r, "/dashboard", http.StatusFound)

}
