FROM node:14
FROM golang:latest

WORKDIR /go-server
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY main.go quotes.go .
RUN go build -o build/quotes .
