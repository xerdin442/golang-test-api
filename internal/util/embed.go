package util

import (
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templateFS embed.FS

func FetchTemplateFile(templateName string) (*template.Template, error) {
	return template.ParseFS(templateFS, "templates/"+templateName)
}
