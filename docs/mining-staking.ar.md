# شبكة وينمار — دليل التعدين والـ Staking (AR)

## نظرة عامة
تستخدم شبكة Winmar خوارزمية Proof-of-Stake. يعني "التعدين" التحقق من الكتل عبر رهن `WNC`.

- الحد الأدنى للرهن: `32,000 WNC`
- زمن الكتلة: `2 ثانية` | الحقبة: `300 فتحة` | النهائية: `≤2 حقب`
- واجهة RPC: `http://localhost:8545`

## المتطلبات
- أنظمة: Linux/Mac/Windows
- المنافذ: `43333/tcp` (P2P)، `8545/tcp` (RPC)، `8546/tcp` (WS)
- مساحة: 50+ GB، معالج: 4 أنوية، ذاكرة: 8 GB

## الخطوة 1 — تشغيل عقدة
- Docker Compose:
```bash
docker-compose up -d
```
- ملف ثنائي:
```bash
make build
./build/wnc-node --config config/node.toml
```

## الخطوة 2 — إنشاء حساب
استخدم محفظة EVM أو JSON-RPC:
```bash
curl -X POST http://localhost:8545 \
 -H 'Content-Type: application/json' \
 -d '{"jsonrpc":"2.0","method":"personal_newAccount","params":["كلمة-مرور-قوية"],"id":1}'
```
احتفظ بعنوانك.

## الخطوة 3 — الحصول على WNC
احصل على `≥ 32,000 WNC` للرهن. في شبكة الاختبار استخدم `config/alloc.json`.

## الخطوة 4 — إنشاء مفاتيح المدقق (BLS)
أنشئ مفاتيح BLS12-381 وبيانات السحب. أضفها إلى `config/validators.json`، ثم إلى `bootValidators` في `config/genesis.json`، وشغّل شبكة الاختبار:
```bash
docker-compose -f ops/testnet-compose.yml up -d
```

## الخطوة 5 — التسجيل على الشبكة الرئيسية (إيداع)
أرسل معاملة `deposit` إلى عقد الرهن تتضمن `pubkey` و`withdrawalCredentials` و`signature` و`amount`.

## الخطوة 6 — التحقق وكسب المكافآت
- حافظ على العقدة تعمل وبزمن استجابة منخفض
- المكافآت تُوزع لكل حقبة؛ يتم حرق الرسوم الأساسية؛ أولوية الرسوم للمُقترِح

## ملفات مهمة
- `config/genesis.json`
- `config/node.toml`
- `ops/testnet-compose.yml`
