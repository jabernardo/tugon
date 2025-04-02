package v1

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"runtime"
)

type MemStats struct {
	Alloc   float64 `json:"alloc"`
	Sys     float64 `json:"sys"`
	LastGc  uint64  `json:"gc_next"`
	NextGc  uint64  `json:"gc_last"`
	CountGc uint64  `json:"gc_cycle"`
}

type Stats struct {
	Message  string   `json:"message"`
	MemStats MemStats `json:"memstats"`
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func byteToMb(b uint64) float64 {
	return float64(b) / 1024 / 1024
}

func Ping(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memstats := &MemStats{
		Alloc:   roundFloat(byteToMb(m.Alloc), 5),
		Sys:     roundFloat(byteToMb(m.Sys), 0),
		LastGc:  m.LastGC,
		NextGc:  m.NextGC,
		CountGc: uint64(m.NumGC),
	}

	err := json.NewEncoder(w).Encode(&Stats{Message: "Up and running!", MemStats: *memstats})

	if err != nil {
		fmt.Fprintf(w, "{\"success\": false}")
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
}
