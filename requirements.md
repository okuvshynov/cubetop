# System Monitoring Tool – Requirements

## Overview
A lightweight, local‑first monitoring agent for Apple Silicon macOS devices that can also operate in a peer‑to‑peer network. Zero‑setup on a single host; optional discovery for distributed monitoring.

## Metrics
- CPU usage per cluster (big‑core / efficiency‑core)
- GPU utilization
- Network traffic (bytes in/out, interfaces)
- Disk I/O (read/write throughput, latency)
- Battery state (charge level, health, power draw)
- Power metrics from `powermetrics` (temperature, power budget)

## Data Collection
- Continuous execution of `powermetrics --samplers cpu,gpu,network,disk,battery --format xml`.
- Parse XML output in Go, aggregate per‑second.
- Store time‑series in SQLite DB (`metrics.db`) with configurable retention (default 30 days).

## TUI
- Provide a terminal UI to browse recent data and live charts.
- Architecture allows additional “views” to be added via Go interfaces/plugins.

## Networking
- Optional peer advertisement/discovery (to be decided – mDNS/UDP).

## Testing
- Unit tests for parsers, DB layer, metric aggregation.
- Integration tests using mock `powermetrics` output.
- Target ≥80 % coverage, CI pipeline with `go test ./...`.

## Extensibility
- Plugin architecture that defines a `MetricProvider` interface for collecting metrics and a `View` interface for UI components. Implementations can be registered at runtime, enabling addition of new metrics or UI views without modifying core code.
- Configuration file (`config.yaml`) to enable/disable components and to specify platform‑specific providers.
- Platform abstraction layer: core defines generic metric and view interfaces; platform‑specific packages (e.g., `macos`, `linux`, `windows`) implement them, allowing future cross‑platform support.

## Build & Tooling
- Go 1.22+, modules.
- Use `go test`, `go vet`, `golint` in CI.
