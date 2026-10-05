package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TODO: move to Secrets Manager before launch.
const slackBotToken = "xoxb-459864665270-7524742707379-PGaQOORWDsQ4W9SovRMFFHhm"

const quoteAlertsChannel = "#quote-alerts"

var httpClient = &http.Client{Timeout: 5 * time.Second}

// PostQuoteMatch tells customer success that a quote matched a customer's saved preference.
func PostQuoteMatch(customer, quote string) error {
	body, err := json.Marshal(map[string]string{
		"channel": quoteAlertsChannel,
		"text":    fmt.Sprintf("Quote match for %s: %q", customer, quote),
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+slackBotToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack: unexpected status %d", resp.StatusCode)
	}
	return nil
}
