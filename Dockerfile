FROM golang:1.18-alpine as builder

# cを使うlibraryのため
RUN apk add build-base

WORKDIR /opt/action

# copy → [currentDir] [WORKDIR]
COPY go.mod go.sum ./

RUN ls

RUN go mod download

COPY . ./

RUN go build -o action-app cmd/api/main.go