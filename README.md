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

## Winmar Golden Protocol 🌟
Winmar Chain implements a unique economic model designed to be Safe, Fun, Durable, and Eternal.

### 1. Halving Schedule (Safe)
Block rewards reduce by 50% every 210,000 blocks (Demo: every 50 blocks).
- Phase 1: 50 WNC
- Phase 2: 25 WNC
- Phase 3: 12.5 WNC
- ...

### 2. Lucky Critical Block (Fun) 🎲
Every block has a **10% chance** to be a "Critical Hit", awarding **2x the normal reward**.
Mining is no longer boring; it's a treasure hunt!

### 3. Tail Emission (Eternal) 🛡️
To ensure the chain lives forever, the block reward will **never drop below 1 WNC**.
This guarantees miners are always incentivized to secure the network, even 100 years from now.

### 4. Persistence & Safety (Durable) 💾
- **Auto-Save**: Chain state is saved to disk (`chaindata.json`) every block.
- **Crash Recovery**: Node automatically resumes from the last block after a restart.
- **Backup Script**: Use `./backup.sh` to instantly backup your entire chain data.

## Dashboard
A built-in dashboard is available to monitor the chain status, rewards, and halving progress.
- Local: http://localhost:8001
- VPS: http://YOUR_VPS_IP:8001

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
