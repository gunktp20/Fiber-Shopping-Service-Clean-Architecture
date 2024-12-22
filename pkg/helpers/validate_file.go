package helpers

import (
	"errors"
	"fmt"
	"net/http"
)

// ValidateFile checks if the file meets the allowed types and size requirements.
func ValidateFile(fileBytes []byte, allowedTypes []string, maxFileSize int64) error {
	// Check file size
	if int64(len(fileBytes)) > maxFileSize {
		return errors.New("file size exceeds the allowed limit")
	}

	// Detect the MIME type of the file
	contentType := http.DetectContentType(fileBytes)

	// Check if the content type is in the allowed list
	isValidType := false
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		// Generate a dynamic error message
		formattedAllowedTypes := formatAllowedTypes(allowedTypes)
		return fmt.Errorf("invalid file type. Only %s are allowed", formattedAllowedTypes)
	}

	return nil
}
