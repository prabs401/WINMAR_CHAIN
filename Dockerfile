FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy dependency files first (if we had go.mod/sum, but we are using simple structure)
COPY . .

# Build the binary
RUN go build -o wnc-node ./cmd/wnc-node

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy binary and config
COPY --from=builder /app/wnc-node .
COPY --from=builder /app/config ./config

# Expose ports (RPC and P2P)
EXPOSE 8545 43333

# Run the node
CMD ["./wnc-node", "--config", "config/node.toml"]
