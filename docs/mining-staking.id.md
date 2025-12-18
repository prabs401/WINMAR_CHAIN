# Winmar Chain — Panduan Mining & Staking (ID)

## Gambaran Umum
Winmar Chain memakai Proof-of-Stake (PoS). "Mining" berarti memvalidasi blok dengan staking `WNC`.

- Stake minimum: `32,000 WNC`
- Waktu blok: `2s` | Epoch: `300 slot` | Finalitas: `≤2 epoch`
- RPC: `http://localhost:8545`

## Persyaratan
- OS: Linux/Mac/Windows
- Port: `43333/tcp` P2P, `8545/tcp` RPC, `8546/tcp` WS
- Storage: 50+ GB, CPU: 4 core, RAM: 8 GB

## Langkah 1 — Jalankan Node
- Docker Compose:
```bash
docker-compose up -d
```
- Binary:
```bash
make build
./build/wnc-node --config config/node.toml
```

Konfigurasi: `config/node.toml`
- P2P: `p2p_port = 43333`
- RPC: `rpc_http_port = 8545`

## Langkah 2 — Buat Akun
Pakai dompet EVM (MetaMask, Ledger) atau JSON-RPC:
```bash
curl -X POST http://localhost:8545 \
 -H 'Content-Type: application/json' \
 -d '{"jsonrpc":"2.0","method":"personal_newAccount","params":["kata-sandi-kuat"],"id":1}'
```
Catat alamat Anda.

## Langkah 3 — Dapatkan WNC
Kumpulkan `≥ 32,000 WNC` untuk staking. Untuk testnet, gunakan alokasi lokal di `config/alloc.json`.

## Langkah 4 — Buat Kunci Validator (BLS)
Siapkan pasangan kunci BLS12-381 dan withdrawal credentials. Tambahkan ke `config/validators.json` untuk testnet lokal, lalu masukkan ke `bootValidators` di `config/genesis.json` dan jalankan:
```bash
docker-compose -f ops/testnet-compose.yml up -d
```

## Langkah 5 — Registrasi Mainnet (Deposit Staking)
Kirim transaksi `deposit` ke kontrak staking berisi `pubkey`, `withdrawalCredentials`, `signature`, `amount`. Lihat ABI contoh pada dokumentasi.

## Langkah 6 — Validasi & Dapatkan Reward
- Pastikan node selalu online dan latensi rendah
- Reward didistribusikan per-epoch; base fee dibakar; priority fee ke proposer

## Berkas Penting
- Genesis: `config/genesis.json`
- Node: `config/node.toml`
- Testnet: `ops/testnet-compose.yml`

## Pemecahan Masalah
- RPC: pastikan port `8545` terbuka dan servis aktif
- P2P: cek bootnodes dan firewall
