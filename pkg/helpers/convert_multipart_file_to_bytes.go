package helpers

import (
	"bytes"
	"io"
	"mime/multipart"
)

// ConvertMultipartFileToBytes converts a *multipart.FileHeader to []byte
func ConvertMultipartFileToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// Open the file from the file header
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file into a byte buffer
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
