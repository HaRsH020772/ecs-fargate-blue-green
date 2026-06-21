# ─── Stage 1: Build ──────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

WORKDIR /build

# Copy go.mod first — cached layer as long as deps don't change
COPY go.mod ./
RUN go mod download

COPY . .

# Build args injected by CodeBuild in buildspec.yml
ARG VERSION=dev
ARG COMMIT_SHA=unknown
ARG BUILD_TIME=unknown

# CGO_ENABLED=0 produces a fully static binary safe for alpine/scratch
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s \
        -X main.version=${VERSION} \
        -X main.commitSHA=${COMMIT_SHA} \
        -X main.buildTime=${BUILD_TIME}" \
    -o app .

# ─── Stage 2: Runtime ─────────────────────────────────────────────────────────
FROM alpine:3.19

# CA certificates for any outbound TLS calls the app might make
RUN apk add --no-cache ca-certificates tzdata

# Non-root user — matches ECS best-practice for Fargate
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /build/app .

USER appuser

EXPOSE 8080

# ALB health check will also validate this before registering the task
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
    CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["./app"]
