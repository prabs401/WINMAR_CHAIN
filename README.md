# Winmar Chain (Winmar Network)

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go](https://img.shields.io/badge/go-1.21-blue.svg)

**Winmar Chain (WNC)** is a high-performance, decentralized Proof-of-Stake blockchain designed for robust infrastructure and community governance.

## Executive Summary

- **Native Coin**: `WNC`
- **Consensus**: Proof-of-Stake (PoS) with VRF Leader Selection & Casper FFG Finality.
- **Block Time**: 2 Seconds.
- **EVM Compatible**: Fully compatible with Ethereum tools (Metamask, Hardhat, Remix).
- **Network ID**: `8822`

## Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose (optional, for easy deployment)
- Make

### Building from Source

```bash
make build
./build/wnc-node --help
```

### Running a Node (Docker)

To run a full node quickly:

```bash
docker-compose up -d
```

This will start a node with:
- P2P Port: `43333`
- RPC Port: `8545`

### Running a Local Testnet

To spin up a 5-node local testnet:

```bash
docker-compose -f ops/testnet-compose.yml up -d
```

### Run from GitHub Container Registry (GHCR)

Once CI publishes the image, you can run directly:

```bash
docker run -p 43333:43333 -p 8545:8545 -p 8546:8546 \
  ghcr.io/prabs401/wnc-node:latest --config /config/node.toml
```

Note: To enable global mining, operators should run their own nodes using the published image and open P2P/RPC ports.

## Mining / Staking Guide

Winmar Chain uses Proof-of-Stake. To participate in consensus and earn `WNC` rewards, follow the guides below:

- English: `docs/mining-staking.en.md`
- Indonesia: `docs/mining-staking.id.md`
- 中文: `docs/mining-staking.zh.md`
- Русский: `docs/mining-staking.ru.md`
- Español: `docs/mining-staking.es.md`
- العربية: `docs/mining-staking.ar.md`

## Technical Specification

### Consensus
- **Algorithm**: PoS + VRF + Casper FFG
- **Epoch Length**: 300 slots (~10 minutes)
- **Finality**: Deterministic after 2 epochs.

### Tokenomics
- **Initial Supply**: 1,000,000,000 WNC
- **Max Supply**: 2,000,000,000 WNC
- **Inflation**: 2% initial, decaying.
- **Fee Burn**: EIP-1559 style base fee burning.

### Network
- **P2P Port**: 43333
- **Bootnodes**:
  - `enode://<pubkey>@seed1.winmar.network:43333`
  - `enode://<pubkey>@seed2.winmar.network:43333`

## Directory Structure

- `cmd/`: Entry points for binaries.
- `config/`: Genesis and node configuration files.
- `ops/`: DevOps and testnet orchestration tools.
- `docs/`: Documentation.

## License

MIT License.
