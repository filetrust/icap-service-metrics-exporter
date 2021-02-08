package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	icap "github.com/filetrust/icap-service-metrics-exporter/pkg"
)

var (
	icapHost   = os.Getenv("ICAP_HOST")
	icapPort   = os.Getenv("ICAP_PORT")
	service    = os.Getenv("SERVICE")
	metricPort = os.Getenv("METRICS_PORT")
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	alive, err := icap.CheckHealth(icapHost, icapPort)
	if err != nil {
		http.Error(w, "icap-service is not available", http.StatusServiceUnavailable)
		return
	}

	if !alive {
		http.Error(w, "icap-service is not available", http.StatusServiceUnavailable)
		return
	}

	w.Write([]byte("icap-service is available"))
}

func run() error {
	registry := prometheus.NewPedanticRegistry()

	c := icap.NewIcapChecker(icapHost, icapPort, service)
	if err := registry.Register(c); err != nil {
		return fmt.Errorf("failed to register icap checker: %v", err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", metricPort))
	if err != nil {
		return fmt.Errorf("failed to listen at %q: %v", metricPort, err)
	}
	defer listen.Close()
	log.Println("listening on", listen.Addr())

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/health/live", healthCheckHandler)
	http.HandleFunc("/health/readiness", healthCheckHandler)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  5 * time.Minute,
	}
	if err := srv.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
