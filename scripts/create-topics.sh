#!/usr/bin/env bash
set -euo pipefail

# Host-side bootstrap; Kafka itself runs in Docker.
# We exec into the Kafka container so you don't need kafka CLI installed locally.

KAFKA_CONTAINER="${KAFKA_CONTAINER:-spotify-etl-kafka}"
INTERNAL_BOOTSTRAP="${INTERNAL_BOOTSTRAP:-kafka:29092}"

create_topic () {
  local topic="$1"
  local partitions="${2:-1}"

  docker exec -it "${KAFKA_CONTAINER}" bash -lc "\
    kafka-topics --bootstrap-server ${INTERNAL_BOOTSTRAP} \
      --create --if-not-exists \
      --topic ${topic} \
      --partitions ${partitions} \
      --replication-factor 1"
}

create_topic "spotify-events" 1
create_topic "spotify-events-retry" 1
create_topic "spotify-events-dlq" 1

echo "Topics created (or already existed)."
