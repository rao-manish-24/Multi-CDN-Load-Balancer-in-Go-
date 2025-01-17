# Multi-CDN-Load-Balancer-in-Go-


# Multi-CDN Load Balancer

## Overview
The Multi-CDN Load Balancer is a lightweight and efficient Go-based application designed to distribute traffic across multiple Content Delivery Network (CDN) providers. It includes:
- **Automatic failover:** Ensures traffic is routed to healthy endpoints.
- **Health checks:** Periodically validates the availability of each CDN endpoint.
- **Metrics collection:** Exposes real-time metrics via Prometheus for monitoring.
- **Dashboard integration:** Provides visual insights into system health using Grafana.

This project is containerized using Docker and orchestrated with Docker Compose, making it easy to deploy and run in any environment.

---

## Features
1. **Multi-CDN Traffic Balancing**: Efficient round-robin routing across multiple CDNs.
2. **Automatic Failover**: Traffic is rerouted to healthy endpoints if a CDN fails.
3. **Customizable Health Checks**: Configurable intervals and timeouts for endpoint validation.
4. **Prometheus Metrics**: Tracks the number of healthy endpoints and other key metrics.
5. **Grafana Dashboard**: Pre-configured dashboard to visualize endpoint health in real-time.

---

## Prerequisites
- **Go** 1.20+  
- **Docker** and **Docker Compose**  
