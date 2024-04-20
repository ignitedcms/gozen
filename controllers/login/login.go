package login

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"gozen/models/users"
	"gozen/system/hash"
	"gozen/system/mail"
	"gozen/system/session"
	"gozen/system/templates"
	"gozen/system/validation"
	"net/http"
)

func ConfirmHash(w http.ResponseWriter, r *http.Request) {

	token := chi.URLParam(r, "token")

	//Now check if token exists in db
	//if it does allow user to update
	//password for given email address
	//if token does not exist throw error

	check := users.CheckToken(token)
	if check == "error" {
      w.Write([]byte("Invalid token"))
	} else {
		fmt.Print("change password")

	   session.Set(w, r, "loggedin", "1")
	   session.Set(w, r, "userid", check)

	   http.Redirect(w, r, "/profile", http.StatusFound)
	}

}

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
	t := users.SetToken(rand, email)
	fmt.Print(t)

	recipientEmail := email
	templatePath := "mail/email_template.html"

	//Warning get site path from .env!!!
	anch := "http://localhost:3000/hash/" + rand

	result := mail.New().
		SetRecipient(recipientEmail).
		SetTemplatePath(templatePath).
		SetAnchor(anch).
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

	//fmt.Print(foo.Password)

	t := hash.CheckPasswordHash(password, foo.Password)

	if !t {
		templates.Errors(w, r, v, postData, "login")
		return
	}

	//set session to loggedin
	//set userid
	session.Set(w, r, "loggedin", "1")
	session.Set(w, r, "name", foo.Name)
	session.Set(w, r, "userid", foo.Id)

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

//profile reset password
//uses session userid

func Password(w http.ResponseWriter, r *http.Request) {

	v := &validation.Validator{}

	v.Required("password", r.FormValue("password"))

	password := r.FormValue("password")

	postData := templates.PostData(w, r)

	if v.HasErrors() {
		templates.Errors(w, r, v, postData, "profile")
		return
	}

   userid := session.Get(r,"userid")
   fmt.Print(password)
   fmt.Print(userid)
}


