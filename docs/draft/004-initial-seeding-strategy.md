# Initial Seeding Strategy for Distributed Go Services using MongoDB

**Status:** Accepted | **Updated:** 2025-11-01 | **Version:** 0.1

## Overview

This document defines the **initial seeding strategy** for Pulap’s distributed Go microservices using **MongoDB** as the primary persistence backend.
Each service manages its own aggregates and repositories, abstracting data storage via a **Repository pattern** compatible with both MongoDB and SQL backends.

Certain collections (e.g., roles, configuration, enums) require predefined values that must be **seeded deterministically**, **idempotently**, and **tracked across deployments**.
Because services run in containerized environments, seeding must be **self-contained**, **automatic**, and **independent of external filesystems**.

## 1. Context

Pulap’s distributed architecture is composed of independent Go microservices, each responsible for a specific domain. To ensure consistency between environments, each service must bootstrap its data automatically at startup.
Manual or filesystem-based seeding introduces deployment inconsistencies and operational complexity in container orchestration environments (e.g., Kubernetes).

Therefore, seeds must be **declared in code**, **executed deterministically**, and **tracked for auditing and idempotence**.

## 2. Decision

We will implement a **code-based seeding mechanism**, executed automatically during service startup or through a dedicated bootstrap command.

The approach is:

* **Declarative:** seeds are defined as versioned Go structures
* **Idempotent:** already applied seeds are skipped
* **Auditable:** each seed is logged in a `_seeds` collection with metadata

Metadata allows tracking of when and by which service a seed was applied.

## 3. Seed Definition

Each seed is a Go structure containing an identifier, description, and execution logic:

```go
type Seed struct {
	ID          string
	Description string
	Run         func(ctx context.Context, db *mongo.Database) error
}
```

### Example

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

Seeds are part of the compiled binary, ensuring version control and reproducibility.

## 4. Seed Tracking

A dedicated collection named `_seeds` will track all executed seeds:

```go
type SeedRecord struct {
	ID          string    `bson:"_id"`
	Application string    `bson:"application"`
	Description string    `bson:"description"`
	AppliedAt   time.Time `bson:"applied_at"`
}
```

### Benefits

* Enables **auditing** (when and which seeds ran)
* Guarantees **idempotence** (no duplicate execution)
* Detects **newly added** seeds automatically

The tracking layer is abstracted behind a common interface:

```go
type SeedTracker interface {
	HasRun(ctx context.Context, id string) (bool, error)
	MarkRun(ctx context.Context, record SeedRecord) error
}
```

## 5. Seed Execution

A generic function applies all registered seeds in order, skipping those already marked as applied:

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

This can be triggered automatically on service startup or via a CLI command (`app seed`).

## 6. Repository Integration

Each domain repository may optionally implement a **Seeder interface** to ensure its own default data:

```go
type Seeder interface {
	EnsureDefaults(ctx context.Context) error
}
```

### MongoDB Example

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

### SQL Example

```go
func (r *SQLConfigRepo) EnsureDefaults(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
	INSERT INTO configs (id, value)
	VALUES ('default_currency', 'USD')
	ON CONFLICT (id) DO NOTHING`)
	return err
}
```

### Bootstrap Orchestration

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

## 7. Deployment Strategy

* **Self-contained:** all seeds are compiled into the service binary
* **Automatic execution:** runs at startup or via CI/CD init containers
* **Immutable design:** modifying seeds requires code changes and rebuilds
* **Cross-backend support:** same logical seeds for MongoDB and SQL
* **Scalable:** can evolve to file-based or hybrid seeding using `embed.FS`

## 8. Alternatives Considered

| Approach                             | Pros                                                       | Cons                                                                 |
| ------------------------------------ | ---------------------------------------------------------- | -------------------------------------------------------------------- |
| **Filesystem-based JSON seeds**      | Human-readable, editable                                   | Requires mounted volumes; fragile in containers                      |
| **Embedded file seeds (`embed.FS`)** | Self-contained, reproducible                               | Requires rebuilds to modify data; runtime validation only            |
| **Code-based seeds (chosen)**        | Versioned, auditable, strongly integrated with domain code | Requires code release for data changes; less readable for large sets |

The chosen solution ensures **deterministic seeding**, **auditability**, and **consistency** across all services while fitting seamlessly into immutable container workflows.

## Conclusion

This seeding strategy provides a unified, code-driven approach to initializing default data in distributed Go microservices.
It integrates with existing repository abstractions, guarantees idempotence and auditability, and supports future evolution toward file-based or hybrid models without disrupting service autonomy.
