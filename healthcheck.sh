#!/bin/sh

HEALTHCHECK_URL=${HEALTHCHECK_URL:-http://127.0.0.1:8777}

exec wget --spider "$HEALTHCHECK_URL"
