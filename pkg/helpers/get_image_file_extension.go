package helpers

import (
	"errors"
	"net/http"
)

// GetFileExtension detects the file extension based on the MIME type
func GetImageFileExtension(fileBytes []byte) (string, error) {
	contentType := http.DetectContentType(fileBytes)

	// Map of allowed MIME types to file extensions
	allowedExtensions := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
		"image/gif":  ".gif",
	}

	if ext, ok := allowedExtensions[contentType]; ok {
		return ext, nil
	}

	return "", errors.New("unsupported file type")
}
