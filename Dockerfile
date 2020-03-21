FROM golang:1.14.1-alpine3.11

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /app

RUN go build -o bin/
RUN ls /app/bin

CMD "/app/bin/gotest1"
