package events

import "time"

type Envelope[T any] struct {
	SchemaVersion int       `json:"schema_version"`
	EventType     string    `json:"event_type"`
	EventID       string    `json:"event_id"`
	OccurredAt    time.Time `json:"occurred_at"`

	Attempt   int    `json:"attempt,omitempty"`
	LastError string `json:"last_error,omitempty"`

	Payload T `json:"payload"`
}

type EpisodeFetchedV1 struct {
	EpisodeID string `json:"episode_id"`
	ShowID    string `json:"show_id"`
}