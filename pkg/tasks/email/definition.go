package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeEmailWelcome = "email:welcome"

type EmailPayload struct {
	To      string
	Subject string
	Body    string
}

func NewEmailWelcomeTask(to, subject, body string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailPayload{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailWelcome, payload), nil
}
