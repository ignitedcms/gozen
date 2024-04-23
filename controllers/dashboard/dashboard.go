package dashboard

import (
	"gozen/models/users"
	"gozen/system/session"
	"gozen/system/templates"
	"net/http"
)

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	userid := session.Get(r, "userid")

	users.Delete(userid)
	http.Redirect(w, r, "/", http.StatusFound)

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

func Dashboard(w http.ResponseWriter, r *http.Request) {

	if session.Get(r, "loggedin") != "1" {

		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	data := map[string]interface{}{
		"name": session.Get(r, "name"),
	}

	templates.Render(w, r, "dashboard", data)
}

func Session(w http.ResponseWriter, r *http.Request) {

	b := session.Get(r, "foo")
	w.Write([]byte(b))
}

func Destroy(w http.ResponseWriter, r *http.Request) {

	session.Destroy(w, r)

	http.Redirect(w, r, "/login", http.StatusFound)
}
