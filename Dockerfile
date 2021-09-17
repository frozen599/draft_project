FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

RUN go build