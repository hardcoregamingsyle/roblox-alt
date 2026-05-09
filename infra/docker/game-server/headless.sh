#!/bin/sh
# NexusEngine Headless Server Wrapper
# Ensures graceful shutdown and environment validation.

set -e

cleanup() {
    echo "Signal received, shutting down NexusClient..."
    kill -TERM "$PID" 2>/dev/null
    wait "$PID"
    exit 0
}

trap cleanup TERM INT

if [ -z "$INSTANCE_ID" ] || [ -z "$AUTH_TOKEN" ]; then
    echo "Error: MISSING_REQUIRED_ENV_VARS"
    exit 1
fi

./NexusClient --headless \
    --main-pack game.pck \
    --fixed-fps 60 \
    --instance-id "$INSTANCE_ID" \
    --auth-token "$AUTH_TOKEN" &

PID=$!
wait $PID