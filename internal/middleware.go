package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func RequestID(ctx *fiber.Ctx) error {
	id := xid.New()
	idStr := id.String()
	ctx.Set(fiber.HeaderXRequestID, idStr)
	ctx.Locals("request-id", idStr)
	return ctx.Next()
}

func Logger(ctx *fiber.Ctx) error {
	idStr, _ := ctx.Locals("request-id").(string)
	start := time.Now()

	chainErr := ctx.Next()

	end := time.Now()
	url := ctx.OriginalURL()
	ip := ctx.IP()
	method := ctx.Method()
	hostname := ctx.Hostname()
	origin := ctx.Get(fiber.HeaderOrigin)
	log.
		Info().
		Str("requestId", idStr).
		Float64("responseTimeMilli", float64(end.Sub(start).Microseconds())/1000).
		Int("status", ctx.Context().Response.StatusCode()).
		Str("url", url).
		Str("ip", ip).
		Str("method", method).
		Str("hostname", hostname).
		Str("origin", origin).
		Err(chainErr).
		Msg("request received")

	return chainErr
}
