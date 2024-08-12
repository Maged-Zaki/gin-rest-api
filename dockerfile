# Build stage
FROM golang:1.22-bookworm as builder
# Must enable CGO cuz of sqlite3
ENV CGO_ENABLED=1

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      build-essential \
      libsqlite3-dev

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server

# Final Stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /src/server /app/server
COPY --from=builder /src/.env /app/.env

# Use a shell to change directory and run the server binary
CMD ["/bin/sh", "-c", "cd /app && ./server"]
