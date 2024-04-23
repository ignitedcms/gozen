/*
|---------------------------------------------------------------
| Routes
|---------------------------------------------------------------
|
| We define all routes here that the main app uses
| We must import the controllers we need
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package routes

import (
	"github.com/go-chi/chi/v5"
	"gozen/controllers/dashboard"
	"gozen/controllers/examples"
	"gozen/controllers/login"
	"gozen/controllers/mail"
	"gozen/controllers/upload"
	"gozen/controllers/welcome"
)

func LoadRoutes(r *chi.Mux) {

	r.Get("/", welcome.Index)

	r.Get("/examples", examples.Index)
	r.Get("/forms", examples.Form)
	r.Post("/forms", examples.FormValidate)

	r.Get("/session", dashboard.Session)
	r.Get("/destroy", dashboard.Destroy) // Should be POST request
	r.Get("/dashboard", dashboard.Dashboard)
	r.Get("/profile", dashboard.ProfileView)
	r.Post("/profile", dashboard.Profile)
	r.Post("/delete-account", dashboard.DeleteAccount)

	r.Get("/login", login.Index)
	r.Post("/login", login.Login)
	r.Post("/password", login.Password)
	r.Get("/hash/{token}", login.ConfirmHash)
	r.Get("/signup", login.SignupView)
	r.Post("/signup", login.Signup)
	r.Get("/forgot", login.ForgotView)
	r.Post("/forgot", login.Forgot)

	r.Get("/socket", examples.Socket)

	r.Get("/upload", upload.Index)
	r.Post("/upload", upload.UploadFile)

	r.Get("/mail", mail.MailView)
	r.Post("/mail", mail.SendMail)

	//r.Get("/users", users.Index)
	//r.Get("/users/create", users.CreateView)
	//r.Post("/users/create", users.Create)
	//r.Get("/users/update/{id}", users.UpdateView)
	//r.Post("/users/update/{id}", users.Update)
	//r.Get("/users/delete/{id}", users.Destroy)

} //end
