module github.com/pulap/pulap/pkg/lib/telemetry

go 1.24.0

toolchain go1.24.7

require (
	github.com/go-chi/chi/v5 v5.2.3
	github.com/pulap/pulap/pkg/lib/core v0.0.0
)

require (
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)

replace github.com/pulap/pulap/pkg/lib/core => ../core
