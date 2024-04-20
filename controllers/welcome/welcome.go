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
	templates.Render(w, r, "welcome", nil)
}

func ProfileView(w http.ResponseWriter, r *http.Request) {

	//restrict access if not logged in
	if session.Get(r, "loggedin") != "1" {

		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	data := map[string]interface{}{
		"name": session.Get(r, "name"),
	}

	templates.Render(w, r, "profile", data)
}

// Validate
func Profile(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("validating"))
}

func Install(w http.ResponseWriter, r *http.Request) {

	templates.Render(w, r, "install", nil)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	// Render the template and write it to the response
	templates.Render(w, r, "dashboard", nil)
}

func Session(w http.ResponseWriter, r *http.Request) {

	b := session.Get(r, "foo")
	w.Write([]byte(b))
}

func Destroy(w http.ResponseWriter, r *http.Request) {

	session.Destroy(w, r)

	http.Redirect(w, r, "/login", http.StatusFound)
}
