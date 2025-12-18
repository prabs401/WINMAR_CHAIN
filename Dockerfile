FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates
COPY wnc-node /usr/local/bin/wnc-node
COPY config /config
ENTRYPOINT ["/usr/local/bin/wnc-node","--config","/config/node.toml"]
