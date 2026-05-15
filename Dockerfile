FROM golang:1.26.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o cloudrun cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/cloudrun .
EXPOSE 8080
CMD ["./cloudrun"]
