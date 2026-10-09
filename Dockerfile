FROM alpine:3.20.3
RUN apk add --no-cache go

ARG GIT_SHA=unknown
LABEL application_name="backend"
LABEL description="Backend service for pantalasa quotes application"
LABEL owner="earthly"
LABEL source_uri="https://github.com/earthly/pantalasa"
LABEL git_sha="${GIT_SHA}"

WORKDIR /go-server
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY main.go quotes.go .
RUN go build -o build/quotes .
