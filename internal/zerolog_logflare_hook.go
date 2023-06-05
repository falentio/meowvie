package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type ZerologLogflareWriter struct {
	url    string
	secret string
}

func NewZerologLogflare(sourceID string, secret string) *ZerologLogflareWriter {
	return &ZerologLogflareWriter{
		url:    "https://api.logflare.app/api/logs?source=" + sourceID,
		secret: secret,
	}
}

func (z *ZerologLogflareWriter) Write(b []byte) (int, error) {
	var data fiber.Map

	if err := json.Unmarshal(b, &data); err != nil {
		return 0, err
	}

	msg := fmt.Sprintf("%s [%s] %s | %.0f | %f | %s", data["message"], data["method"], data["url"], data["status"], data["responseTimeMilli"], data["requestId"])
	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(fiber.Map{
		"event_message": msg,
		"metadata":      data,
	}); err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", z.url, body)
	if err != nil {
		return 0, err
	}
	req.Header.Set("X-Api-Key", z.secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		msg := bytes.NewBuffer(nil)
		_, err := msg.ReadFrom(resp.Body)
		if err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("%d status code received, with message %q", resp.StatusCode, msg.String())
	}

	return len(b), nil
}

var _ io.Writer = new(ZerologLogflareWriter)
