/*
|---------------------------------------------------------------
| fileupload
|---------------------------------------------------------------
|
| A file utility to handle uploads, set upload directory,
| set maxiumum file size, and file type restrictions
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package fileupload

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// FileUpload provides methods for handling file uploads
type FileUpload struct {
	file         multipart.File
	handler      *multipart.FileHeader
	maxFileSize  int64
	allowedTypes []string
	destFolder   string
}

// New creates a new FileUpload instance
func New() *FileUpload {
	return &FileUpload{}
}

// File sets the file to be uploaded
func (fu *FileUpload) File(file multipart.File, handler *multipart.FileHeader) *FileUpload {
	fu.file = file
	fu.handler = handler
	return fu
}

// MaxFileSize sets the maximum file size allowed for upload
func (fu *FileUpload) MaxFileSize(size string) *FileUpload {
	if size == "" {
		return fu
	}

	unit := strings.ToLower(size[len(size)-2:])
	if unit != "kb" && unit != "mb" && unit != "gb" {
		return fu
	}

	value, err := strconv.ParseInt(size[:len(size)-2], 10, 64)
	if err != nil {
		return fu
	}

	switch unit {
	case "kb":
		fu.maxFileSize = value * 1024
	case "mb":
		fu.maxFileSize = value * 1024 * 1024
	case "gb":
		fu.maxFileSize = value * 1024 * 1024 * 1024
	}

	return fu
}

// AllowedTypes sets the allowed MIME types for upload
func (fu *FileUpload) AllowedTypes(types ...string) *FileUpload {
	fu.allowedTypes = types
	return fu
}

// DestinationFolder sets the destination folder where the file will be uploaded
func (fu *FileUpload) DestinationFolder(folder string) *FileUpload {
	fu.destFolder = folder
	return fu
}

// Upload performs the file upload
func (fu *FileUpload) Upload() (string, error) {
	// Check max file size
	if fu.maxFileSize > 0 && fu.handler.Size > fu.maxFileSize {
		return "", fmt.Errorf("File size exceeds the maximum allowed size")
	}

	// Check file type
	if len(fu.allowedTypes) > 0 && !fu.isValidFileType() {
		return "", fmt.Errorf("Invalid file type")
	}

	// Generate a random filename
	randomFilename := randomFilename() + filepath.Ext(fu.handler.Filename)

	// Create the uploads directory if it doesn't exist
	uploadsDir := fu.destFolder
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		os.Mkdir(uploadsDir, 0755)
	}

	// Create a new file in the uploads directory to store the uploaded file
	uploadedFile, err := os.Create(filepath.Join(uploadsDir, randomFilename))
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	// Copy the file to the destination on the server
	_, err = io.Copy(uploadedFile, fu.file)
	if err != nil {
		return "", err
	}

	return randomFilename, nil
}

func (fu *FileUpload) isValidFileType() bool {
	// Detect the MIME type of the file
	ext := filepath.Ext(fu.handler.Filename)

	// Detect the MIME type based on the file extension
	mimeType := mime.TypeByExtension(ext)

	// Check if the MIME type is allowed
	for _, allowedType := range fu.allowedTypes {
		if mimeType == allowedType {
			return true
		}
	}
	return false
}

func randomFilename() string {
	return uuid.New().String()
}
