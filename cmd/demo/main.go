package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a classic histogram with custom buckets
var classicHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "request_latency",
	Help:    "A classic histogram with custom buckets",
	Buckets: prometheus.DefBuckets,
})

// Define a native histogram (available in Prometheus 2.40+)
var nativeHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:                        "native_request_latency",
	Help:                        "A native histogram",
	Buckets:                     nil, // Must be nil for native histograms
	NativeHistogramBucketFactor: 1.2, // Optional: controls density/resolution
})

func refreshMetrics() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for range ticker.C {
			// Simulate some new observations
			for i := 0; i < 10; i++ {
				latency := rand.ExpFloat64() / 100
				classicHistogram.Observe(latency)
				nativeHistogram.Observe(latency)
			}
			log.Println("Updated metrics")
		}
	}()
}

func main() {
	reg := prometheus.NewRegistry()

	// Register metrics
	reg.MustRegister(classicHistogram)
	reg.MustRegister(nativeHistogram)

	// Start recording metrics
	go refreshMetrics()

	// Setup HTTP handler
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Println("Serving metrics at :8080/metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
