### Tracing

This document summarizes the complete tracing discussion and implementation details for the system.

**Contents:**

1. Concept: Request correlation and tracing propagation
2. Middleware for Request ID
3. Context helpers for Request ID
4. Integration with `Tracer`
5. Integration points (admin, business frontend)
6. Example usage

---

#### Request correlation and tracing propagation

Each incoming request must carry a unique correlation ID (e.g., `X-Request-ID`). If no ID is present, it is generated at the first entrypoint (e.g., `admin` or `business frontend`).

* The ID is stored in the request context and echoed in the response headers.
* Downstream services must preserve this ID in outbound requests.
* The `Tracer` can use this ID for local span correlation or distributed trace integration (e.g., OpenTelemetry's `traceparent`).

Later, when an API Gateway is introduced, it becomes the new top-level generator of these IDs.

---

#### Middleware for Request ID

**File:** `pkg/lib/core/requestid_middleware.go`

```go
package middleware

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/yourorg/yourapp/pkg/lib/core"
)

const requestIDHeader = "X-Request-ID"

func RequestIDMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        reqID := r.Header.Get(requestIDHeader)
        if reqID == "" {
            reqID = uuid.NewString()
        }

        ctx := core.WithRequestID(r.Context(), reqID)
        w.Header().Set(requestIDHeader, reqID)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

---

#### Context helpers for Request ID

**File:** `pkg/lib/core/request_context.go`

```go
package core

import "context"

type ctxKey string

const requestIDKey ctxKey = "request_id"

func WithRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFrom(ctx context.Context) string {
    if v, ok := ctx.Value(requestIDKey).(string); ok {
        return v
    }
    return ""
}
```

---

#### Integration with `Tracer`

The tracer can automatically pick up the request ID from the context when creating a span.

**Simplified example:**

```go
ctx, span := xparams.Tracer.Start(r.Context(), "ListUsers")
defer span.End()
```

If no distributed tracing backend is active, this still provides correlation between logs and local spans.

---

#### Integration points (admin, business frontend)

Entry services (`admin`, `business frontend`) should apply the `RequestIDMiddleware` at router level:

```go
r := chi.NewRouter()
r.Use(middleware.RequestIDMiddleware)
r.Use(middleware.Recoverer)
r.Use(middleware.Logger)
```

All other internal services simply preserve and forward the header.

---

#### 6. Example request flow

1. A user sends a request to `admin`.
2. Middleware checks for `X-Request-ID` → not found → generates one.
3. `admin` logs the request and passes the header to downstream services.
4. `orders` and `inventory` read `X-Request-ID` and keep it unchanged.
5. All logs and spans across the call chain share the same correlation ID.

---

This setup provides a lightweight, framework-independent foundation for distributed tracing and correlation, ready to scale into full OpenTelemetry integration later without changing service code.
