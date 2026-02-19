FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY *.go .
RUN apk add --no-cache git
RUN CGO_ENABLED=0 go build -o /app/phonebook .

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/phonebook /app/phonebook
EXPOSE 8080
VOLUME ["/app"]
ENTRYPOINT ["/app/phonebook"]
