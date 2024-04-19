package formutils

import (
	"gozen/system/session"
	"gozen/system/templates"
	"gozen/system/validation"
	"net/http"
)

type TemplateData struct {
	PostData       map[string]interface{}
	PostDataErrors map[string]interface{}
	FlashData      string
	Foo            string
}

func SetAndGetPostData(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	session.SetOldPostData(w, r)
	return session.GetOldPostData(w, r)
}

func HandleValidationErrors(w http.ResponseWriter,
	r *http.Request,
	validator *validation.Validator,
	postData map[string]interface{}, templatePath string) {
	postDataErrors := make(map[string]interface{})
	for _, err := range validator.GetErrors() {
		postDataErrors[err.Field] = err.Message
	}

	data := TemplateData{
		PostData:       postData,
		PostDataErrors: postDataErrors,
		FlashData:      "Failed, error occurred",
	}
	templates.RenderTemplate(w, r, templatePath, data)
}
