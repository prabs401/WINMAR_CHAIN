# Winmar Chain — Guía de minería y staking (ES)

## Resumen
Winmar Chain usa Proof-of-Stake (PoS). "Minería" significa validar bloques apostando `WNC`.

- Participación mínima: `32,000 WNC`
- Tiempo de bloque: `2s` | Época: `300 slots` | Finalidad: `≤2 épocas`
- RPC: `http://localhost:8545`

## Requisitos
- SO: Linux/Mac/Windows
- Puertos: `43333/tcp` P2P, `8545/tcp` RPC, `8546/tcp` WS
- Almacenamiento: 50+ GB, CPU: 4 núcleos, RAM: 8 GB

## Paso 1 — Ejecutar un nodo
- Docker Compose:
```bash
docker-compose up -d
```
- Binario:
```bash
make build
./build/wnc-node --config config/node.toml
```

## Paso 2 — Crear una cuenta
Use una billetera EVM o JSON-RPC:
```bash
curl -X POST http://localhost:8545 \
 -H 'Content-Type: application/json' \
 -d '{"jsonrpc":"2.0","method":"personal_newAccount","params":["contraseña-segura"],"id":1}'
```
Guarde su dirección.

## Paso 3 — Obtener WNC
Adquiera `≥ 32,000 WNC` para staking. En testnet, use `config/alloc.json`.

## Paso 4 — Generar claves de validador (BLS)
Genere claves BLS12-381 y credenciales de retiro. Añada al `config/validators.json`, incluya en `bootValidators` `config/genesis.json` y lance el testnet:
```bash
docker-compose -f ops/testnet-compose.yml up -d
```

## Paso 5 — Registro en mainnet (depósito)
Envíe una transacción `deposit` al contrato de staking con `pubkey`, `withdrawalCredentials`, `signature`, `amount`.

## Paso 6 — Validar y ganar recompensas
- Mantenga el nodo en línea y con baja latencia
- Recompensas por época; base fee quemada; priority fee para el proponente

## Archivos útiles
- `config/genesis.json`
- `config/node.toml`
- `ops/testnet-compose.yml`
