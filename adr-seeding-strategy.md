# ADR: Initial Seeding Strategy for Distributed Go Services using MongoDB

## Context

Our distributed system consists of multiple Go microservices, each responsible for managing its own aggregates and persistence layer.  
We use a **Repository pattern** to abstract the underlying datastore, allowing the use of either **MongoDB** (default) or an SQL database for persistence.

Certain collections must always contain a predefined set of values (e.g., roles, default configuration, status enums). These values must be **seeded deterministically**, **idempotently**, and **tracked over time** to ensure consistent environments across deployments and services.

Running in a containerized and orchestrated environment (e.g., Kubernetes), we cannot rely on shared filesystems or manual seeding steps. Therefore, seeds must be **self-contained** within the service image.

---

## Decision

We will implement a **code-based seeding system**, versioned and executed automatically at service startup (or via a dedicated bootstrap command).

The system will be **declarative, idempotent, and auditable**, storing metadata of applied seeds in the database to detect newly added ones.

---

## Design Overview

### 1. Seed Definition

Each seed will be represented as a versioned structure:

```go
type Seed struct {
	ID          string
	Description string
	Run         func(ctx context.Context, db *mongo.Database) error
}
```

Seeds are defined in code, ensuring they are version-controlled, type-safe, and deployed alongside the service.  
Example:

```go
var seeds = []Seed{
	{
		ID: "2025-10-30_initial_configs",
		Description: "Insert default configuration entries",
		Run: func(ctx context.Context, db *mongo.Database) error {
			_, err := db.Collection("configs").UpdateOne(
				ctx,
				bson.M{"_id": "default_currency"},
				bson.M{"$setOnInsert": bson.M{"_id": "default_currency", "value": "USD"}},
				options.Update().SetUpsert(true),
			)
			return err
		},
	},
}
```

---

### 2. Seed Tracking

A dedicated collection (`_seeds`) will track applied seeds:

```go
type SeedRecord struct {
	ID          string    `bson:"_id"`
	Application string    `bson:"application"`
	Description string    `bson:"description"`
	AppliedAt   time.Time `bson:"applied_at"`
}
```

This enables:
- **Auditing** of what seeds ran and when.  
- **Idempotence** — seeds won’t run twice.  
- **Detection** of newly added seeds.

The tracking layer will be abstracted behind a `SeedTracker` interface:

```go
type SeedTracker interface {
	HasRun(ctx context.Context, id string) (bool, error)
	MarkRun(ctx context.Context, record SeedRecord) error
}
```

---

### 3. Seed Application Logic

A generic runner ensures all seeds are executed exactly once per environment:

```go
func ApplySeeds(ctx context.Context, tracker SeedTracker, seeds []Seed) error {
	for _, s := range seeds {
		ok, _ := tracker.HasRun(ctx, s.ID)
		if ok {
			continue
		}
		if err := s.Run(ctx); err != nil {
			return fmt.Errorf("seed %s failed: %w", s.ID, err)
		}
		if err := tracker.MarkRun(ctx, SeedRecord{
			ID:          s.ID,
			Application: "orders",
			Description: s.Description,
			AppliedAt:   time.Now(),
		}); err != nil {
			return err
		}
		log.Printf("Applied seed: %s", s.ID)
	}
	return nil
}
```

The function can be triggered automatically at startup or exposed as a dedicated command (`app seed`).

---

### 4. Repository Integration

Each repository implementing the domain persistence layer (e.g., `ConfigRepo`, `RoleRepo`) will also implement the following optional interface:

```go
type Seeder interface {
	EnsureDefaults(ctx context.Context) error
}
```

This enables domain-level seeding independent of storage backend.

For MongoDB:

```go
func (r *MongoConfigRepo) EnsureDefaults(ctx context.Context) error {
	_, err := r.coll.UpdateOne(
		ctx,
		bson.M{"_id": "default_currency"},
		bson.M{"$setOnInsert": bson.M{"_id": "default_currency", "value": "USD"}},
		options.Update().SetUpsert(true),
	)
	return err
}
```

For SQL:

```go
func (r *SQLConfigRepo) EnsureDefaults(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
	INSERT INTO configs (id, value)
	VALUES ('default_currency', 'USD')
	ON CONFLICT (id) DO NOTHING`)
	return err
}
```

Both can be run under the same `Bootstrap()` orchestration:

```go
func Bootstrap(ctx context.Context, repos ...Seeder) error {
	for _, repo := range repos {
		if err := repo.EnsureDefaults(ctx); err != nil {
			return err
		}
	}
	return nil
}
```

---

### 5. Deployment Considerations

- **No filesystem dependency:** all seeds live as Go code inside the binary.  
- **Automatic execution:** seeds can run on service startup or via a CI/CD init container.  
- **Immutability:** new seeds require code changes and redeploy — consistent with immutable image philosophy.  
- **Shared semantics:** if SQL support is enabled, both backends share the same logical seed list.  

---

## Alternatives Considered

| Approach | Pros | Cons |
|-----------|------|------|
| **Filesystem-based JSON seeds** | Editable, human-readable | Requires volume mounts; error-prone in containers |
| **Embedded JSON (`embed.FS`)** | Self-contained; reproducible | Rebuilds required; less readable; still detached from domain logic |
| **Code-based seeds (chosen)** | Type-safe, version-controlled, easily testable | Requires code release to change data |

---

## Outcome

This design provides a unified, auditable, and backend-agnostic seeding mechanism.  
It integrates naturally with the repository abstraction, follows immutable infrastructure principles, and ensures reproducibility across environments and deployments.
