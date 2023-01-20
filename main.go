package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()

			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	sizeRequest = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "myapp_processed",
		Help: "The total bytes",
	})
)

func main() {
	recordMetrics()

	http.DefaultServeMux.HandleFunc("/echo", handler)
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	log.Println("listen on 8080")

	http.ListenAndServe("0.0.0.0:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	sizeRequest.Observe(float64(r.ContentLength))
	w.Write([]byte("hello, world"))
}
