package formutils

import (
	"gozen/system/rendering"
	"gozen/system/sessionstore"
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
	sessionstore.SetOldPostData(w, r)
	return sessionstore.GetOldPostData(w, r)
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
	rendering.RenderTemplate(w, r, templatePath, data)
}
