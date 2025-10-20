package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	mail "github.com/go-mail/mail/v2"
	"github.com/hibiken/asynq"
)

func HandleEmailTask(c context.Context, t *asynq.Task) error {
	var p EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	fmt.Printf("🚀 Menerima task kirim email ke %s, subject: %s", p.To, p.Subject)
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_ADDRESS"))
	m.SetHeader("To", p.To)
	m.SetHeader("Subject", p.Subject)
	m.SetBody("text/plain", p.Body)

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT_SMTP"))
	if err != nil {
		return err
	}
	d := mail.NewDialer(os.Getenv("MAIL_SMTP"), port, os.Getenv("MAIL_ADDRESS"), os.Getenv("MAIL_PASSWORD"))

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("❌ Gagal mengirim email ke %s, subject: %s\n", p.To, p.Subject)
		return err
	}

	fmt.Printf("Email terkirim ke %s\n", p.To)
	return nil
}
