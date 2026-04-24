FROM golang:1.22-alpine

ARG GIT_SHA=unknown
LABEL application_name="notifier"
LABEL description="Outbound webhook dispatcher for quote-match notifications"
LABEL owner="backend@pantalasa.org"
LABEL source_uri="https://github.com/pantalasa/backend"
LABEL git_sha="${GIT_SHA}"

RUN apk add --no-cache ca-certificates

WORKDIR /app

CMD ["/bin/sh", "-c", "echo notifier placeholder; sleep infinity"]
