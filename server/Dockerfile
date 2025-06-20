FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

FROM ubuntu:22.04

WORKDIR /app

COPY --from=build /app/app .
COPY .env .env

COPY database ./database

ENV PORT=8080

EXPOSE 8080

CMD ["./app"]
