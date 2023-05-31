FROM golang:1.20-alpine as builder

WORKDIR /build
COPY go* .
RUN go mod download

RUN apk add --no-cache build-base

COPY internal internal
COPY cmd cmd
RUN CGO_ENABLED=1 go build -o meowvie ./cmd/meowvie

FROM alpine

WORKDIR /app
COPY --from=builder /build/meowvie meowvie
CMD ["./meowvie"]