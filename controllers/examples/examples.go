package examples

import (
	"fmt"
	"gozen/system/formutils"
	"gozen/system/rendering"
	"gozen/system/validation"
	"net/http"
)

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "examples", nil)
}

func Socket(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "socket", nil)
}

func Form(w http.ResponseWriter, r *http.Request) {
	data := formutils.TemplateData{
		Foo: "hi", //some data mostly a model
	}
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "forms", data)
}

// post request
func FormValidate(w http.ResponseWriter, r *http.Request) {
	validator := &validation.Validator{}

	validator.Required("name", r.FormValue("name")).
		Email("email", r.FormValue("email"))

	postData := formutils.SetAndGetPostData(w, r)

	//get checkbox postdata
	fmt.Println(r.FormValue("numbers"))
	fmt.Println(r.FormValue("foo"))
	//we NEED to use just Form for checkboxes
	arr := r.Form["animals"]

	for index := range arr {
		fmt.Print(arr[index])
	}

	if validator.HasErrors() {
		formutils.HandleValidationErrors(w, r, validator, postData, "forms")
		return
	}
	//else success

	http.Redirect(w, r, "/forms", http.StatusFound)
}
