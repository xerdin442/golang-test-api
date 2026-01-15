package util

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templateFS embed.FS

func ParseEmailTemplate(data any, templateName string) (string, error) {
	tmpl, err := template.ParseFS(templateFS, "templates/"+templateName)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
