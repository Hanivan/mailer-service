package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/mrz1836/postmark"
)

type EmailParams struct {
	PostmarkToken  string `json:"postmark_token"`
	From           string `json:"from"`
	To             string `json:"to"`
	Subject        string `json:"subject"`
	Base64HTMLBody string `json:"base64_html_body"`
}

func SendToGmail(postmarkToken string, accountToken string, params EmailParams) (postmark.EmailResponse, error) {
	client := postmark.NewClient(postmarkToken, accountToken)
	currentDate := time.Now()
	dayOfMonth := currentDate.Day()
	nextMonth := currentDate.AddDate(0, 1, 0)
	startOfNextMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 1, 0, 0, 0, 0, nextMonth.Location())
	dayEndOfThisMonth := startOfNextMonth.Add(-time.Second).Day()

	decodedHtmlBody, _ := base64.StdEncoding.DecodeString(params.Base64HTMLBody)
	email := postmark.Email{
		From:     params.From,
		To:       params.To,
		Subject:  params.Subject,
		HTMLBody: string(decodedHtmlBody),
	}

	usedQuota, _ := client.GetSentCounts(context.Background(), nil)
	if dayOfMonth <= dayEndOfThisMonth && usedQuota.Sent == 100 {
		return postmark.EmailResponse{}, errors.New("monthly quota limit reached")
	}
	postmarkResponse, err := client.SendEmail(context.Background(), email)

	if err != nil {
		return postmark.EmailResponse{}, err
	}

	return postmarkResponse, nil
}
