FROM golang:1.24-alpine AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/app .

ENV SWAGGER_HOST=snet-api.fly.dev

ENV PORT=8080

EXPOSE 8080

CMD ["./app"]
