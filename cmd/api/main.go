package main

import (
	"log"
	"net/http"

	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	tracer := otel.Tracer("spotify-etl-api")
	_ = tracer

	router := chi.NewRouter()

	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	_ = middleware.OapiRequestValidator

	var _ *pgxpool.Pool

	_ = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_ = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "spotify-events",
		GroupID: "spotify-etl",
	})

	requestsTotal := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	})
	prometheus.MustRegister(requestsTotal)

	router.Handle("/metrics", promhttp.Handler())

	logger.Info("starting api", zap.String("addr", ":8080"))
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal("api exited", zap.Error(err))
	}
}

func _noopTracerProvider() trace.TracerProvider {
	return otel.GetTracerProvider()
}
