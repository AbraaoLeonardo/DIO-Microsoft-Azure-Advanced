FROM golang:1.25.3-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o bin/main

FROM alpine:3.22.2
WORKDIR /app
COPY --from=builder /app/bin/main .
EXPOSE 8080
ENTRYPOINT ["./main"]