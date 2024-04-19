package welcome

import (
	"gozen/system/session"
	"gozen/system/templates"
	"net/http"
)

// welcome page
func Index(w http.ResponseWriter, r *http.Request) {

	session.Set(w, r, "foo", "A session test")
	// Render the template and write it to the response
	templates.RenderTemplate(w, r, "welcome", nil)
}

func Install(w http.ResponseWriter, r *http.Request) {

	templates.RenderTemplate(w, r, "install", nil)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	// Render the template and write it to the response
	templates.RenderTemplate(w, r, "dashboard", nil)
}

func Session(w http.ResponseWriter, r *http.Request) {

	b := session.Get(r, "foo")
	w.Write([]byte(b))
}

func Destroy(w http.ResponseWriter, r *http.Request) {

	session.Destroy(w, r)

	http.Redirect(w, r, "/login", http.StatusFound)
}
