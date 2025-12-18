package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Persistence
const StateFile = "chaindata.json"

type ChainState struct {
	Height   int64               `json:"height"`
	Balances map[string]*big.Int `json:"balances"`
}

func saveState() {
	mu.Lock()
	defer mu.Unlock()

	state := ChainState{
		Height:   atomic.LoadInt64(&height),
		Balances: balances,
	}

	log.Println("Saving state to chaindata.json...") // Debug log
	file, err := json.MarshalIndent(state, "", " ")
	if err == nil {
		err = os.WriteFile(StateFile, file, 0644)
		if err != nil {
			log.Printf("❌ Failed to write state file: %v", err)
		}
	} else {
		log.Printf("❌ Failed to marshal state: %v", err)
	}
}

func loadState() {
	file, err := os.ReadFile(StateFile)
	if err != nil {
		return // Start from scratch
	}

	var state ChainState
	if err := json.Unmarshal(file, &state); err == nil {
		atomic.StoreInt64(&height, state.Height)
		mu.Lock()
		balances = state.Balances
		mu.Unlock()
		fmt.Printf("⚠️ RECOVERED CHAIN STATE! Resuming from Block #%d\n", state.Height)
	}
}

// Global state
var (
	height         int64
	currentHash    string
	mu             sync.Mutex
	balances       = make(map[string]*big.Int)
	rewardAddr     string
	rewardPerBlock = new(big.Int)
	initialReward  = new(big.Int)
	minReward      = new(big.Int)
)

const HalvingInterval = 50 // Fast halving for demo purposes (usually 210,000)

// RPC types
type JSONRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      any    `json:"id"`
}

type JSONRPCResponse struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      any       `json:"id"`
	Result  any       `json:"result,omitempty"`
	Error   *RPCError `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Winmar Chain (WNC) Node v1.1.0 (Golden Protocol)")

	// Load previous state if exists
	loadState()

	fmt.Println("Initializing Winmar Network...")
	fmt.Println("Loading configuration...")

	// Initialize configuration
	rewardAddr = os.Getenv("WNC_REWARD_ADDRESS")
	if rewardAddr == "" {
		rewardAddr = "0xlocal-validator"
	}
	// Initial reward: 50 WNC
	initialReward.SetString("50000000000000000000", 10)
	// Minimum Reward (Tail Emission): 1 WNC
	minReward.SetString("1000000000000000000", 10)

	rewardPerBlock.Set(initialReward)

	// Simulate startup
	time.Sleep(1 * time.Second)
	fmt.Println("Genesis block loaded: 0x0000000000000000")
	fmt.Println("Network ID: 8822")
	fmt.Println("Listening on P2P port 43333")
	fmt.Println("JSON-RPC endpoint active at 0.0.0.0:8545")

	// Standard endpoints
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	http.HandleFunc("/block", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.Header().Set("Content-Type", "application/json")
		h := atomic.LoadInt64(&height)
		mu.Lock()
		hash := currentHash
		mu.Unlock()
		if hash == "" {
			hash = "0x0000000000000000000000000000000000000000000000000000000000000000"
		}
		fmt.Fprintf(w, `{"height":%d,"hash":"%s"}`, h, hash)
	})

	http.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		w.Header().Set("Content-Type", "application/json")
		addr := r.URL.Query().Get("address")
		if addr == "" {
			addr = rewardAddr
		}
		mu.Lock()
		bal, ok := balances[addr]
		if !ok {
			bal = new(big.Int)
		}
		balStr := bal.String()
		mu.Unlock()
		wncf := toWNC(bal)
		fmt.Fprintf(w, `{"address":"%s","balanceWei":"%s","balanceWNC":"%s"}`, addr, balStr, wncf)
	})

	// JSON-RPC handler
	http.HandleFunc("/", handleRPC)

	// Mining simulation
	go func() {
		t := time.NewTicker(2 * time.Second)
		for range t.C {
			h := atomic.AddInt64(&height, 1)

			// Generate deterministic simulated hash
			sum := sha256.Sum256([]byte(fmt.Sprintf("winmar-block-%d", h)))
			hash := fmt.Sprintf("0x%x", sum)

			// Calculate Halving
			halvings := uint(h / HalvingInterval)
			// Right shift initial reward by number of halvings (equivalent to dividing by 2^n)
			currentBlockReward := new(big.Int).Rsh(initialReward, halvings)

			// 1. ABADI MECHANISM: Tail Emission
			// If reward drops below MinReward, clamp it to MinReward
			if currentBlockReward.Cmp(minReward) < 0 {
				currentBlockReward.Set(minReward)
			}

			// 2. MENYENANGKAN MECHANISM: Lucky Critical Block
			// 10% Chance to get 2x Reward
			isCritical := false
			if rand.Intn(100) < 10 {
				multiplier := big.NewInt(2)
				currentBlockReward.Mul(currentBlockReward, multiplier)
				isCritical = true
			}

			logMsg := fmt.Sprintf("Proposed block #%d Hash: %s | Reward: %s WNC", h, hash, toWNC(currentBlockReward))
			if isCritical {
				logMsg += " [🔥 CRITICAL HIT! 2x REWARD 🔥]"
			}
			log.Println(logMsg)

			mu.Lock()
			currentHash = hash
			rewardPerBlock.Set(currentBlockReward) // Update global state for API
			if _, ok := balances[rewardAddr]; !ok {
				balances[rewardAddr] = new(big.Int)
			}
			balances[rewardAddr].Add(balances[rewardAddr], currentBlockReward)
			mu.Unlock()

			// Save state every block
			saveState()

			// Announce halving event
			if h > 0 && h%HalvingInterval == 0 {
				log.Printf("⚠️ HALVING EVENT! Block Reward reduced to %s WNC", toWNC(currentBlockReward))
			}
		}
	}()

	// Start server with CORS support for MetaMask
	log.Fatal(http.ListenAndServe(":8545", nil))
}

func handleRPC(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	enableCors(&w)

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req JSONRPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	res := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      req.ID,
	}

	switch req.Method {
	case "eth_chainId":
		res.Result = "0x2276" // 8822
	case "net_version":
		res.Result = "8822"
	case "web3_clientVersion":
		res.Result = "WNC/v1.0.0/darwin/go1.21"
	case "eth_blockNumber":
		h := atomic.LoadInt64(&height)
		res.Result = fmt.Sprintf("0x%x", h)
	case "eth_gasPrice":
		res.Result = "0x3b9aca00" // 1 Gwei
	case "eth_getBalance":
		if len(req.Params) > 0 {
			addr, ok := req.Params[0].(string)
			if ok {
				// Normalize address (lowercase)
				addr = strings.ToLower(addr)
				mu.Lock()
				// Check exact match or case-insensitive match
				var found *big.Int
				for k, v := range balances {
					if strings.EqualFold(k, addr) {
						found = v
						break
					}
				}
				if found == nil {
					found = new(big.Int)
				}
				balStr := fmt.Sprintf("0x%x", found)
				mu.Unlock()
				res.Result = balStr
			} else {
				res.Result = "0x0"
			}
		} else {
			res.Result = "0x0"
		}
	case "eth_getCode":
		res.Result = "0x"
	case "eth_estimateGas":
		res.Result = "0x5208" // 21000
	case "eth_call":
		res.Result = "0x"
	default:
		// Return zero for unknown read methods to avoid MetaMask errors
		if strings.HasPrefix(req.Method, "eth_") {
			res.Result = "0x0"
		} else {
			res.Error = &RPCError{Code: -32601, Message: "Method not found"}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func toWNC(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	base := new(big.Float).SetInt(wei)
	denom := new(big.Float).SetFloat64(1e18)
	val := new(big.Float).Quo(base, denom)
	s := val.Text('f', 18)
	i := len(s) - 1
	for i >= 0 && s[i] == '0' {
		i--
	}
	if i >= 0 && s[i] == '.' {
		i--
	}
	return s[:i+1]
}
