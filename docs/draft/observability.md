# Observability Notes

## Metrics

- `pkg/lib/core/metrics.go` defines the `core.Metrics` interface with a noop implementation.
- `core.ApplyStack` now wires `core.NewMetricsMiddleware`, so every request reports duration, method, status.
- Each service receives a metrics collector via `XParams.Metrics()`; override with `SetMetrics` when wiring a real backend (Prometheus, OTEL, etc.).

## Health & Lifecycle

- `core.NewHealthRegistry` aggregates liveness/readiness checks contributed by components implementing `HealthReporter`.
- `core.Setup` now returns `*HealthRegistry` (ignored by default) and exposes `/healthz`, `/livez`, `/readyz` with per-check results.
- Custom checks: implement `HealthReporter` or call `registry.Register(Readiness|Liveness)` after `Setup`.
- Default checks keep services reporting `ok` until dependencies are registered.
