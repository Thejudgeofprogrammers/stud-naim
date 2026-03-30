FROM golang:1.26-alpine3.23 AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api/main.go


FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .

COPY ./static ./static
COPY .env.dev .
ENV CONFIG_FILE=.env.dev

RUN mkdir -p /app/uploads

CMD ["./app"]