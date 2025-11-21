FROM cgr.dev/chainguard/go:1.21

LABEL application_name="backend" \
      description="Backend service for pantalasa quotes application" \
      owner="earthly" \
      source_uri="https://github.com/earthly/pantalasa"

WORKDIR /go-server
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY main.go quotes.go .
RUN go build -o build/quotes .
