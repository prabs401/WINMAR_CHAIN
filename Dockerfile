FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY cmd/wnc-node ./cmd/wnc-node
RUN go build -o /out/wnc-node ./cmd/wnc-node

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/wnc-node /usr/local/bin/wnc-node
COPY config /config
EXPOSE 43333 8545 8546
ENTRYPOINT ["/usr/local/bin/wnc-node","--config","/config/node.toml"]
