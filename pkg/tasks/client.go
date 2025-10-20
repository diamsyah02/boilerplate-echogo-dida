package tasks

import (
	"log"

	emailTask "boilerplate-echogo-dida/pkg/tasks/email"

	"github.com/hibiken/asynq"
)

type TaskDistributor struct {
	client *asynq.Client
}

func NewTaskDistributor(redisAddr string) *TaskDistributor {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	return &TaskDistributor{client: client}
}

func (d *TaskDistributor) DistributeEmail(to, subject, body string) error {
	task, err := emailTask.NewEmailWelcomeTask(to, subject, body)
	if err != nil {
		return err
	}

	// Optionally add delay, retry, etc
	_, err = d.client.Enqueue(task)
	if err != nil {
		return err
	}

	log.Printf("📨 Task dikirim ke queue untuk: %s\n", to)
	return nil
}
