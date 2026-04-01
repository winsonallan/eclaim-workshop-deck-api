# syntax=docker.io/docker/dockerfile:1

# ========= BUILDER =========
FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o server ./cmd/server/

# ========= RUNNER =========
FROM alpine:3.21 AS runner
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

RUN addgroup --system --gid 1001 goapp && \
    adduser --system --uid 1001 --ingroup goapp ginuser

COPY --from=builder /app/server .

USER ginuser

EXPOSE 8080

CMD ["./server"]