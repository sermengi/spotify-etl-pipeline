FROM golang:1.23-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api ./cmd/api

FROM alpine:3.20

RUN adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=builder /api /app/api

EXPOSE 8080

USER appuser

ENTRYPOINT ["/app/api"]
