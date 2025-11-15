# Stage 1: build
FROM golang:1.24 AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# STATIC BUILD
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags "-s -w" -o app .

# Stage 2: runtime
FROM alpine:3.22.2

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 5700

CMD ["./app"]

