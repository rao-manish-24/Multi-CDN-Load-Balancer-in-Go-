package config

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

type Config struct {
    Endpoints           []string
    ListenAddress       string
    HealthCheckInterval time.Duration
    HealthCheckTimeout  time.Duration
}

func LoadConfig() (*Config, error) {
    defaultEndpoints := []string{
        "https://cdn1.example.com",
        "https://cdn2.example.com",
    }

    endpointsEnv := os.Getenv("CDN_ENDPOINTS")
    var endpoints []string
    if endpointsEnv == "" {
        endpoints = defaultEndpoints
    } else {
        endpoints = strings.Split(endpointsEnv, ",")
    }

    listenAddress := os.Getenv("LB_LISTEN_ADDRESS")
    if listenAddress == "" {
        listenAddress = ":8080"
    }

    hciStr := os.Getenv("HEALTH_CHECK_INTERVAL")
    if hciStr == "" {
        hciStr = "10"
    }
    hciVal, err := strconv.Atoi(hciStr)
    if err != nil {
        return nil, fmt.Errorf("invalid HEALTH_CHECK_INTERVAL: %w", err)
    }

    hctStr := os.Getenv("HEALTH_CHECK_TIMEOUT")
    if hctStr == "" {
        hctStr = "3"
    }
    hctVal, err := strconv.Atoi(hctStr)
    if err != nil {
        return nil, fmt.Errorf("invalid HEALTH_CHECK_TIMEOUT: %w", err)
    }

    return &Config{
        Endpoints:           endpoints,
        ListenAddress:       listenAddress,
        HealthCheckInterval: time.Duration(hciVal) * time.Second,
        HealthCheckTimeout:  time.Duration(hctVal) * time.Second,
    }, nil
}
