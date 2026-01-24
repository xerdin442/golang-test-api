package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/resend/resend-go/v2"
	"github.com/rs/zerolog/log"
)

type EmailPayload struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

func NewEmailTask(data *EmailPayload) (*asynq.Task, error) {
	payload, _ := json.Marshal(data)
	return asynq.NewTask("email_queue", payload), nil
}

func (h *TaskHandler) HandleEmailTask(ctx context.Context, t *asynq.Task) error {
	var p EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	resendClient := resend.NewClient(h.cfg.ResendEmailApiKey)
	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <%s>", h.cfg.AppName, h.cfg.AppEmail),
		To:      []string{p.Recipient},
		Subject: p.Subject,
		Html:    p.Content,
	}

	_, err := resendClient.Emails.Send(params)
	if err != nil {
		log.Error().Err(err).Msgf("Error occured while sending %s email to %s", p.Subject, p.Recipient)
		return err
	}

	log.Info().Msgf("%s email sent to %s", p.Subject, p.Recipient)
	return nil
}
