# Spotify ETL Pipeline (Go)

Go-based ETL pipeline for Spotify podcasts.

## Tech stack (Phase 0)

- Go 1.22
- API framework: chi
- Postgres: pgx
- Redis: go-redis/v9
- Kafka: kafka-go
- Observability: prometheus client, zap logging, OpenTelemetry
- OpenAPI tooling: oapi-codegen

## Phases

- **Phase 0 — Repo scaffold & local infra**
  - [x] Go toolchain & module setup
  - [ ] Local Postgres/Redis/Kafka with Docker
  - [ ] Basic API skeleton

## Running with Docker

Build and run the API locally:

```bash
docker build -t spotify-etl-api .
docker run --rm -p 8080:8080 spotify-etl-api
```
or with docker compose:
```bash
docker compose up --build
```

More details to come as the project evolves.
