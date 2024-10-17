FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o api-server .

FROM alpine:latest 

WORKDIR /app

COPY --from=builder /app/api-server .

EXPOSE 8000

CMD ["./api-server"]
