# Build the manager binary
FROM golang:1.24.5 AS builder

WORKDIR /workspace

# Copy Go modules first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY cmd/controller/main.go cmd/controller/main.go
COPY api/ api/
COPY internal/ internal/

# Build binary for linux/amd64 only
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o manager cmd/controller/main.go

# Minimal image with distroless
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
