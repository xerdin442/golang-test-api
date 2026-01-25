package util

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadConfig struct {
	Folder      string
	CloudName   string
	ApiKey      string
	CloudSecret string
}

func ParseImageMimetype(file multipart.File) error {
	buffer := make([]byte, 512)
	file.Read(buffer)

	// Detect MIME type of file
	contentType := http.DetectContentType(buffer)

	// Reset file pointer
	file.Seek(0, io.SeekStart)

	// Verify if the MIME type is supported
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/heic" {
		return errors.New("Unsupported MIME type")
	}

	return nil
}

func ProcessFileUpload(file multipart.File, config *UploadConfig) (*uploader.UploadResult, error) {
	cld, _ := cloudinary.NewFromParams(config.CloudName, config.ApiKey, config.CloudSecret)
	return cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: config.Folder,
	})
}
