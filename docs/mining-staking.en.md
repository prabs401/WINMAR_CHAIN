# Winmar Chain — Mining & Staking Guide (EN)

## Overview
Winmar Chain uses Proof-of-Stake (PoS). "Mining" refers to validating blocks by staking `WNC`.

- Minimum stake: `32,000 WNC`
- Block time: `2s` | Epoch: `300 slots` | Finality: `≤2 epochs`
- RPC: `http://localhost:8545` (default)

## Requirements
- OS: Linux/Mac/Windows
- Ports: `43333/tcp` P2P, `8545/tcp` RPC, `8546/tcp` WS
- Disk: 50+ GB (mainnet), CPU: 4 cores, RAM: 8 GB

## Step 1 — Run a Node
- Docker Compose:
```bash
docker-compose up -d
```
- Binary:
```bash
make build
./build/wnc-node --config config/node.toml
```

Config file: `config/node.toml`
- P2P: `p2p_port = 43333`
- RPC: `rpc_http_port = 8545`

## Step 2 — Create an Account
Use any EVM wallet (MetaMask, Ledger) or JSON-RPC:
```bash
curl -X POST http://localhost:8545 \
 -H 'Content-Type: application/json' \
 -d '{"jsonrpc":"2.0","method":"personal_newAccount","params":["strong-passphrase"],"id":1}'
```
Record your address.

## Step 3 — Acquire WNC
Obtain `≥ 32,000 WNC` for staking. For testnet, use local allocations in `config/alloc.json`.

## Step 4 — Generate Validator Keys (BLS)
Prepare BLS12-381 validator keypair and withdrawal credentials. Add to `config/validators.json` for local testnet:
```json
[
  {
    "pubkeyBls": "0x...",
    "stake": "64000",
    "withdrawalCredentials": "0x...",
    "signature": "0x..."
  }
]
```
Then include your validator in `bootValidators` inside `config/genesis.json` and start the testnet:
```bash
docker-compose -f ops/testnet-compose.yml up -d
```

## Step 5 — Register on Mainnet (Staking Deposit)
On mainnet, send a `deposit` transaction to the staking contract with fields `pubkey`, `withdrawalCredentials`, `signature`, `amount`. ABI example:
```json
[
  {"type":"function","name":"deposit","inputs":[
    {"name":"pubkey","type":"bytes"},
    {"name":"withdrawalCredentials","type":"bytes"},
    {"name":"signature","type":"bytes"},
    {"name":"amount","type":"uint256"}
  ],"outputs":[]}
]
```

## Step 6 — Validate and Earn Rewards
- Keep node online and low-latency
- Rewards are distributed per epoch; base fee is burned, priority fee goes to proposer

## Useful Files
- Genesis: `config/genesis.json`
- Node: `config/node.toml`
- Testnet compose: `ops/testnet-compose.yml`

## Troubleshooting
- RPC errors: ensure `8545` open and service running
- P2P sync slow: check bootnodes and firewall
