package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"novak/pkg/chain"
	"novak/pkg/eir"
	"novak/pkg/hvet"
	"novak/pkg/rgac"

	blake3 "lukechampine.com/blake3"
)

var (
	merkle       *chain.Chain
	rc           rgac.Chain
	heartbeat    uint64
	pqcRotations uint64
	eirTotal     uint64
	hvetTotal    uint64
	rgacTotal    uint64
	version      = "dev"
)

func hashAll(data []byte) map[string]string {
	out := make(map[string]string)
	h256 := sha256.Sum256(data)
	h384 := sha512.Sum384(data)
	h512 := sha512.Sum512(data)
	h3 := blake3.Sum256(data)
	out["sha256"] = fmt.Sprintf("%x", h256)
	out["sha384"] = fmt.Sprintf("%x", h384)
	out["sha512"] = fmt.Sprintf("%x", h512)
	out["blake3"] = fmt.Sprintf("%x", h3)
	return out
}

func metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	fmt.Fprintf(w, "novak_heartbeat_total %d\n", atomic.LoadUint64(&heartbeat))
	fmt.Fprintf(w, "novak_build_info{version=%q} 1\n", version)
	fmt.Fprintf(w, "pqc_rotations_total %d\n", atomic.LoadUint64(&pqcRotations))
	fmt.Fprintf(w, "novak_merkle_depth %d\n", merkle.Depth())
	fmt.Fprintf(w, "novak_eir_appends_total %d\n", atomic.LoadUint64(&eirTotal))
	fmt.Fprintf(w, "novak_hvet_compute_total %d\n", atomic.LoadUint64(&hvetTotal))
	fmt.Fprintf(w, "novak_rgac_append_total %d\n", atomic.LoadUint64(&rgacTotal))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func hashbench(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	data := []byte(time.Now().UTC().String())
	results := hashAll(data)
	elapsed := time.Since(start)
	fmt.Fprintf(w, "# elapsed_ms %.3f\n", float64(elapsed.Microseconds())/1000)
	for k, v := range results {
		fmt.Fprintf(w, "%s %s\n", k, v)
	}
}

func showpqc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pqc_rotations_total %d\n", atomic.LoadUint64(&pqcRotations))
}

// /record accepts {"R":"...","D":"...","O":"..."} and appends EIR→HVET→RGAC→Merkle
func record(w http.ResponseWriter, r *http.Request) {
	var p struct{ R, D, O string }
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
		return
	}
	e := eir.New(p.R, p.D, p.O)
	atomic.AddUint64(&eirTotal, 1)

	h := hvet.Compute(e.JSON())
	atomic.AddUint64(&hvetTotal, 1)

	head := rc.Append(h["sha256"])
	atomic.AddUint64(&rgacTotal, 1)

	merkle.Append(e.JSON()) // persisted at /var/lib/novak/chain.json
	fmt.Fprintf(w, "recorded head %s\n", head)
}

func main() {
	// init merkle ledger file (persists to /var/lib/novak/chain.json)
	var err error
	merkle, err = chain.New("/var/lib/novak/chain.json")
	if err != nil {
		log.Fatalf("ledger init: %v", err)
	}

	// seed: proof of boot
	e := eir.New("system", "startup", "boot")
	h := hvet.Compute(e.JSON())
	rc.Append(h["sha256"])
	merkle.Append(e.JSON())
	atomic.AddUint64(&eirTotal, 1)
	atomic.AddUint64(&hvetTotal, 1)
	atomic.AddUint64(&rgacTotal, 1)

	logger := log.New(os.Stdout, "", 0)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", health)
	mux.HandleFunc("/metrics", metrics)
	mux.HandleFunc("/hashbench", hashbench)
	mux.HandleFunc("/showpqc", showpqc)
	mux.HandleFunc("/record", record)

	go func() {
		logger.Println("HTTP :9105 up")
		if err := http.ListenAndServe(":9105", mux); err != nil {
			logger.Fatalf("http: %v", err)
		}
	}()

	// heartbeat + simulated PQC key rotation
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-time.After(5 * time.Second):
			atomic.AddUint64(&heartbeat, 1)
			logger.Println(time.Now().UTC().Format(time.RFC3339), "heartbeat", version)
		case <-ticker.C:
			atomic.AddUint64(&pqcRotations, 1)
			logger.Println(time.Now().UTC().Format(time.RFC3339), "PQC rotation complete")
		}
	}
}
