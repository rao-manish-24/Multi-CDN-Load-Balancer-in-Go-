package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/mux"
    "github.com/yourusername/multi-cdn-load-balancer/internal/config"
    "github.com/yourusername/multi-cdn-load-balancer/internal/healthcheck"
    "github.com/yourusername/multi-cdn-load-balancer/internal/loadbalancer"
    "github.com/yourusername/multi-cdn-load-balancer/internal/metrics"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    lb := loadbalancer.NewLoadBalancer(cfg.Endpoints, cfg.HealthCheckTimeout)

    go func() {
        ticker := time.NewTicker(cfg.HealthCheckInterval)
        defer ticker.Stop()
        for range ticker.C {
            healthcheck.UpdateHealthStatus(lb)
        }
    }()

    metrics.RegisterMetrics(lb)

    router := mux.NewRouter()
    router.Handle("/metrics", metrics.GetHandler())
    router.HandleFunc("/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {
        loadbalancer.HandleRequest(lb, w, r)
    }).Methods(http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions)

    srv := &http.Server{
        Addr:    cfg.ListenAddress,
        Handler: router,
    }

    stopChan := make(chan os.Signal, 1)
    signal.Notify(stopChan, os.Interrupt)

    go func() {
        log.Printf("Starting Multi-CDN Load Balancer on %s", cfg.ListenAddress)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Load Balancer failed: %v", err)
        }
    }()

    <-stopChan
    log.Println("Shutting down Multi-CDN Load Balancer...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown error: %v", err)
    }

    log.Println("Load Balancer stopped gracefully.")
}
