package main

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendEmail(domain, apiKey, sender, recipient, subject, body string) error {
	mg := mailgun.NewMailgun(domain, apiKey)
	message := mailgun.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := mg.Send(ctx, message)
	return err
}
