package util

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/xerdin442/api-practice/internal/env"
)

func ParseImageMimetype(file multipart.File) error {
	buffer := make([]byte, 512)
	file.Read(buffer)

	// Detect MIME type of file
	contentType := http.DetectContentType(buffer)

	// Verify if the MIME type is supported
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/heic" {
		return errors.New("Unsupported MIME type")
	}

	return nil
}

func ProcessFileUpload(file multipart.File, folder string) (*uploader.UploadResult, error) {
	cloudName := env.GetStr("CLOUDINARY_NAME")
	cloudSecret := env.GetStr("CLOUDINARY_SECRET")
	apiKey := env.GetStr("CLOUDINARY_API_KEY")

	cld, _ := cloudinary.NewFromParams(cloudName, cloudSecret, apiKey)

	return cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: folder,
	})
}
