package utils

import (
	"errors"
	"net/http"
)

func GetImageFileExtension(fileBytes []byte) (string, error) {
	contentType := http.DetectContentType(fileBytes)

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
