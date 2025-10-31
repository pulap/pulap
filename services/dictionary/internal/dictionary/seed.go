package dictionary

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Seed represents a versioned database seed operation.
type Seed struct {
	ID          string
	Description string
	Run         func(ctx context.Context, db *mongo.Database) error
}

// SeedRecord tracks when a seed was applied.
type SeedRecord struct {
	ID          string    `bson:"_id"`
	Application string    `bson:"application"`
	Description string    `bson:"description"`
	AppliedAt   time.Time `bson:"applied_at"`
}

// SeedTracker tracks which seeds have been applied to prevent duplicate runs.
type SeedTracker interface {
	HasRun(ctx context.Context, id string) (bool, error)
	MarkRun(ctx context.Context, record SeedRecord) error
}

// MongoSeedTracker implements SeedTracker using MongoDB.
type MongoSeedTracker struct {
	collection *mongo.Collection
}

// NewMongoSeedTracker creates a new MongoDB-backed seed tracker.
func NewMongoSeedTracker(db *mongo.Database) *MongoSeedTracker {
	return &MongoSeedTracker{
		collection: db.Collection("_seeds"),
	}
}

// HasRun checks if a seed with the given ID has already been applied.
func (t *MongoSeedTracker) HasRun(ctx context.Context, id string) (bool, error) {
	var record SeedRecord
	err := t.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("failed to check seed status: %w", err)
	}
	return true, nil
}

// MarkRun records that a seed has been applied.
func (t *MongoSeedTracker) MarkRun(ctx context.Context, record SeedRecord) error {
	_, err := t.collection.InsertOne(ctx, record)
	if err != nil {
		return fmt.Errorf("failed to mark seed as run: %w", err)
	}
	return nil
}

// ApplySeeds executes all seeds that have not been run yet.
func ApplySeeds(ctx context.Context, db *mongo.Database, tracker SeedTracker, seeds []Seed, application string) error {
	for _, s := range seeds {
		ok, err := tracker.HasRun(ctx, s.ID)
		if err != nil {
			return fmt.Errorf("failed to check if seed %s has run: %w", s.ID, err)
		}

		if ok {
			// Seed already applied, skip
			continue
		}

		// Run the seed
		if err := s.Run(ctx, db); err != nil {
			return fmt.Errorf("seed %s failed: %w", s.ID, err)
		}

		// Mark as completed
		if err := tracker.MarkRun(ctx, SeedRecord{
			ID:          s.ID,
			Application: application,
			Description: s.Description,
			AppliedAt:   time.Now(),
		}); err != nil {
			return fmt.Errorf("failed to mark seed %s as run: %w", s.ID, err)
		}
	}
	return nil
}

// UpsertHelper is a convenience function for upserting documents in seeds.
func UpsertHelper(ctx context.Context, collection *mongo.Collection, filter, update interface{}) error {
	_, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$setOnInsert": update},
		options.Update().SetUpsert(true),
	)
	return err
}
