FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./cmd/bin/app ./cmd/app/main.go

FROM alpine AS prod

WORKDIR /

COPY --from=builder /app/cmd/bin/app /cmd/bin/app
COPY --from=builder /app/.env .
COPY --from=builder /app/config/config.yml /config/

EXPOSE 8080

CMD ["/cmd/bin/app"]
