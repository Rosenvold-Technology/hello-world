
# ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
# ┃             Multi‑Stage Docker Build             ┃
# ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o hello ./cmd/server

# ┏━━━ Final Image ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
FROM gcr.io/distroless/static-debian12
WORKDIR /root/
COPY --from=builder /app/hello .
COPY web ./web
EXPOSE 8080
ENTRYPOINT ["./hello"]
