package users

import (
	"gozen/models/users"
	"gozen/system/formutils"
	"gozen/system/rendering"
	"gozen/system/validation"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	users, err := users.All()
	if err != nil {
		// Handle error
		return
	}
	rendering.RenderTemplate(w, r, "users-all", users)
}

func CreateView(w http.ResponseWriter, r *http.Request) {
	data := formutils.TemplateData{
		// You can set data here as needed
	}
	rendering.RenderTemplate(w, r, "users-create", data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	validator := &validation.Validator{}
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	validator.Required("name", name)
	validator.Required("email", email)
	validator.Required("password", password)
	postData := formutils.SetAndGetPostData(w, r)
	if validator.HasErrors() {
		formutils.HandleValidationErrors(w, r, validator, postData, "users-create")
		return
	}
	// Else, no validation errors, proceed with creation
	t, _ := users.Create(name, email, password)
	fmt.Print(t)
	http.Redirect(w, r, "/users", http.StatusFound)
}

func UpdateView(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := users.Read(id)
	if err != nil {
		// Handle error
		return
	}
	data := map[interface{}]interface{}{
		"Id":    id,
		"Users": user,
	}
	rendering.RenderTemplate(w, r, "users-update", data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	if err := users.Update(id, name, email, password); err != nil {
		// Handle error
		return
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := users.Delete(id); err != nil {
		// Handle error
		return
	}
	http.Redirect(w, r, "/users", http.StatusFound)
}
