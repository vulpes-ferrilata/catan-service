FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./catan-service ./cmd

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/catan-service .

COPY ./locales ./locales
COPY ./config.yaml .

EXPOSE 8080

ENTRYPOINT ["/catan-service"]