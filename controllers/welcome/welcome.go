package welcome

import (
	"gozen/system/rendering"
	"gozen/system/sessionstore"
	"net/http"
)

// welcome page
func Index(w http.ResponseWriter, r *http.Request) {

	sessionstore.SetSession(w, r, "foo", "A session test")
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "welcome", nil)
}

func Install(w http.ResponseWriter, r *http.Request) {

	rendering.RenderTemplate(w, r, "install", nil)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "dashboard", nil)
}

func Session(w http.ResponseWriter, r *http.Request) {

	b := sessionstore.GetSession(r, "foo")
	w.Write([]byte(b))
}

func Destroy(w http.ResponseWriter, r *http.Request) {

	sessionstore.DestroySession(w, r)

	http.Redirect(w, r, "/login", http.StatusFound)
}
