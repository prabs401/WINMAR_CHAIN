# Winmar Chain — 挖矿与质押指南 (中文)

## 概述
Winmar Chain 采用权益证明（PoS）。“挖矿”即通过质押 `WNC` 参与区块验证。

- 最低质押：`32,000 WNC`
- 区块时间：`2 秒` | 纪元：`300 个时隙` | 终局：`≤2 个纪元`
- RPC：`http://localhost:8545`

## 要求
- 系统：Linux/Mac/Windows
- 端口：`43333/tcp`（P2P）、`8545/tcp`（RPC）、`8546/tcp`（WS）
- 硬件：50+ GB 磁盘、4 核 CPU、8 GB 内存

## 第一步 — 运行节点
- Docker Compose：
```bash
docker-compose up -d
```
- 二进制：
```bash
make build
./build/wnc-node --config config/node.toml
```

## 第二步 — 创建账户
使用 EVM 钱包（MetaMask、Ledger）或 JSON-RPC：
```bash
curl -X POST http://localhost:8545 \
 -H 'Content-Type: application/json' \
 -d '{"jsonrpc":"2.0","method":"personal_newAccount","params":["强密码"],"id":1}'
```
记录地址。

## 第三步 — 获取 WNC
至少准备 `32,000 WNC` 用于质押。测试网可用 `config/alloc.json` 本地分配。

## 第四步 — 生成验证者密钥（BLS）
准备 BLS12-381 密钥与提现凭证。将验证者加入 `config/validators.json`，并添加到 `config/genesis.json` 的 `bootValidators`，启动本地测试网：
```bash
docker-compose -f ops/testnet-compose.yml up -d
```

## 第五步 — 主网注册（质押存款）
向质押合约发送 `deposit` 交易，参数包含 `pubkey`、`withdrawalCredentials`、`signature`、`amount`。

## 第六步 — 验证与获得奖励
- 保持节点在线、延迟低
- 奖励按纪元分配；基础费销毁；优先费归提议者

## 重要文件
- `config/genesis.json`
- `config/node.toml`
- `ops/testnet-compose.yml`
