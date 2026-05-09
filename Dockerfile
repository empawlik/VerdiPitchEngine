# Stage 1: Builder
FROM golang:alpine AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod ./
# COPY go.sum ./ # Not created yet, but would be here if there were dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/verdi-pitch-engine ./cmd/verdi

# Stage 2: Runtime
FROM ubuntu:22.04

# Avoid tzdata interactive prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Install ffmpeg with librubberband support and flac for metadata tagging
RUN apt-get update && \
    apt-get install -y ffmpeg flac && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /go/bin/verdi-pitch-engine /usr/local/bin/verdi-pitch-engine

# Copy the orchestration wrapper script
COPY scripts/verdi-process.sh /usr/local/bin/verdi-process
RUN chmod +x /usr/local/bin/verdi-process

# Default environment variables for worker count and paths
ENV VERDI_WORKERS=4
ENV VERDI_IN=/music_in
ENV VERDI_OUT=/music_out

# Execute the engine
ENTRYPOINT ["/usr/local/bin/verdi-pitch-engine"]
