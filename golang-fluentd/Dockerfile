FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o webserver

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/webserver .

EXPOSE 8080

CMD ["./webserver"]