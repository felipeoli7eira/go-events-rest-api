FROM golang:1.23-alpine3.21 AS builder

ENV TZ America/Sao_Paulo

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -o api

EXPOSE 8080

CMD [ "./api" ]
