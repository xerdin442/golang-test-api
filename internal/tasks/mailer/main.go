package mailer

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/hibiken/asynq"
	"github.com/resend/resend-go/v2"
	"github.com/rs/zerolog/log"
	"github.com/xerdin442/api-practice/internal/env"
)

type EmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

type OnboardingTemplateData struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

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

func NewEmailTask(data *EmailPayload) (*asynq.Task, error) {
	payload, _ := json.Marshal(data)
	return asynq.NewTask("email_queue", payload), nil
}

func HandleEmailTask(ctx context.Context, t *asynq.Task) error {
	var p EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	resendClient := resend.NewClient(env.GetStr("RESEND_EMAIL_API_KEY"))
	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <%s>", env.GetStr("APP_NAME"), env.GetStr("APP_EMAIL")),
		To:      []string{p.Recipient},
		Subject: p.Subject,
		Html:    p.Content,
	}

	_, err := resendClient.Emails.Send(params)
	if err != nil {
		return err
	}

	log.Info().Msgf("Onboarding email sent to %s", p.Recipient)
	return nil
}
