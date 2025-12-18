# Winmar Chain — Руководство по майнингу и стейкингу (RU)

## Обзор
Winmar Chain использует Proof-of-Stake (PoS). «Майнинг» — это валидация блоков через стейкинг `WNC`.

- Минимальный стейк: `32,000 WNC`
- Время блока: `2с` | Эпоха: `300 слотов` | Финализация: `≤2 эпох`
- RPC: `http://localhost:8545`

## Требования
- ОС: Linux/Mac/Windows
- Порты: `43333/tcp` P2P, `8545/tcp` RPC, `8546/tcp` WS
- Диск: 50+ ГБ, CPU: 4 ядра, RAM: 8 ГБ

## Шаг 1 — Запуск узла
- Docker Compose:
```bash
docker-compose up -d
```
- Бинарник:
```bash
make build
./build/wnc-node --config config/node.toml
```

## Шаг 2 — Создать аккаунт
Используйте EVM-кошелёк или JSON-RPC:
```bash
curl -X POST http://localhost:8545 \
 -H 'Content-Type: application/json' \
 -d '{"jsonrpc":"2.0","method":"personal_newAccount","params":["надежный-пароль"],"id":1}'
```
Сохраните адрес.

## Шаг 3 — Получить WNC
Приобретите `≥ 32,000 WNC` для стейкинга. Для тестнета используйте `config/alloc.json`.

## Шаг 4 — Сгенерировать ключи валидатора (BLS)
Генерируйте BLS12-381 ключи и реквизиты вывода. Добавьте валидатора в `config/validators.json`, внесите в `bootValidators` `config/genesis.json`, запустите тестнет:
```bash
docker-compose -f ops/testnet-compose.yml up -d
```

## Шаг 5 — Регистрация в мейннете (депозит)
Отправьте транзакцию `deposit` в стейкинг-контракт с параметрами `pubkey`, `withdrawalCredentials`, `signature`, `amount`.

## Шаг 6 — Валидация и вознаграждения
- Держите узел онлайн с низкой задержкой
- Вознаграждения распределяются по эпохам; base fee сжигается; priority fee получает пропозер

## Важные файлы
- `config/genesis.json`
- `config/node.toml`
- `ops/testnet-compose.yml`
