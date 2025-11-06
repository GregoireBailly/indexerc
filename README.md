# 🧮 indexERC — Ethereum ERC-20 Activity Indexer

![CI Status](https://github.com/GregoireBailly/indexerc/actions/workflows/ci.yml/badge.svg)

### Overview

**indexERC** is a lightweight Go project that connects directly to the Ethereum network (for now using infura) and analyzes on-chain ERC-20 transfer activity.  
It was built to explore Go’s language and usual architecture, while diving into the Ethereum blockchain ecosystem in a clean, tested, and CI-ready way.

The current version focuses on setting up a robust foundation — clean abstractions, testing, and CI/CD — before expanding into more complex analytics.

---

## 🎯 Goals

- ✅ Re-learn and explore **Go** and the **Ethereum API** through a real, modular project.  
- ✅ Be as clean as possible while still moving fast
- ✅ Include **tests, documentation, and CI/CD** from day one.  
- 🧩 Keep the project lightweight but extensible for future features.

---

## 🚀 How to Use

1. Set up environment variables for your Ethereum API provider (`ETH_RPC_URL`, `ETH_RPC_API_KEY`)  

2. Run the indexer to connect to Ethereum and count ERC-20 transfers.  
`go run ./cmd/indexerc`

Expected output:

```bash
🚀 Starting indexERC — connecting to Ethereum mainnet…
✅ Connected!
Count of ERC20 transfers is: 42
```

3. Run the tests to validate the implementation and CI setup.

`make test`

Or directly

`go test ./...`


---

## 🧩 Next Steps

### Short Term
- ✅ Add **integration tests** using cassettes of some sort (e.g., with `go-vcr`).  
- 🧪 Implement **error and retry logic** in the client.  
- 🔍 Improve the analyzer to **filter per token** (e.g., USDT, DAI, etc.).  

### Medium Term
- 📈 Compute and **rank the most active ERC-20 tokens** by transfer volume.  
- 💾 Add **caching** for recent blocks in a local database.  
- 🕒 Introduce a **cron job or scheduler** to update results periodically.  
- 🌐 Expose metrics via a **REST or GraphQL API**.

---

## 🧱 Design Principles

- **Separation of concerns** — clear distinction between blockchain data access and analytics logic.  
- **Dependency injection** — all services receive their dependencies as interfaces, improving testability and flexibility.  
- **Testability** — each component can be mocked and verified independently.  
- **Extensibility** — designed to scale with new analyzers, providers, and data backends.  
- **Simplicity first** — prioritize clarity and correctness over premature optimization.  
- **CI/CD ready** — all code paths are linted and tested automatically.

---
