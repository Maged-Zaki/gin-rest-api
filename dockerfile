# Build Stage
FROM golang:1.22-bookworm as builder
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

# Build the statically linked binary
RUN go build -ldflags '-extldflags "-static"' -o server

# Final Stage
FROM scratch

WORKDIR /app
COPY --from=builder /src/server /app/server
COPY --from=builder /src/.env /app/.env

CMD ["/app/server"]