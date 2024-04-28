package upload

import (
	"fmt"
	"gozen/system/fileupload"
	"gozen/system/templates"
	"net/http"
)

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	templates.Render(w, r, "upload", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Perform file upload
	randomFilename, err := fileupload.New().
		File(file, handler).
		MaxFileSize("5mb"). // 5 MB limit
		AllowedTypes("image/jpeg", "image/png", "application/pdf", "application/zip", "image/svg+xml", "text/xml").
		DestinationFolder("./uploads").
		Upload()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully with filename: %s", randomFilename)
}
