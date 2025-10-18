package mongo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pulap/pulap/services/estate/internal/estate"
)

// EstateMongoRepo implements the EstateRepo interface using MongoDB.
// MongoDB is ideal for aggregates since each aggregate can be stored as a single document.
type EstateMongoRepo struct {
	collection *mongo.Collection
}

// NewEstateMongoRepo creates a new MongoDB repository for Estate aggregates.
func NewEstateMongoRepo(db *mongo.Database) *EstateMongoRepo {
	return &EstateMongoRepo{
		collection: db.Collection("estates"),
	}
}

// Create creates a new Estate aggregate in MongoDB.
// The entire aggregate (root + children) is stored as a single document.
func (r *EstateMongoRepo) Create(ctx context.Context, aggregate *estate.Estate) error {
	if aggregate == nil {
		return fmt.Errorf("aggregate cannot be nil")
	}

	aggregate.EnsureID()
	aggregate.BeforeCreate()

	_, err := r.collection.InsertOne(ctx, aggregate)
	if err != nil {
		return fmt.Errorf("could not create Estate aggregate: %w", err)
	}

	return nil
}

// Get retrieves a complete Estate aggregate by ID from MongoDB.
// Returns the aggregate root with all its child entities loaded.
func (r *EstateMongoRepo) Get(ctx context.Context, id uuid.UUID) (*estate.Estate, error) {
	var aggregate estate.Estate

	filter := bson.M{"_id": id.String()}
	err := r.collection.FindOne(ctx, filter).Decode(&aggregate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Estate aggregate with ID %s not found", id.String())
		}
		return nil, fmt.Errorf("could not get Estate aggregate: %w", err)
	}

	return &aggregate, nil
}

// Save performs a unit-of-work save operation on the Estate aggregate.
// In MongoDB, this is straightforward since the entire aggregate is replaced as one document.
func (r *EstateMongoRepo) Save(ctx context.Context, aggregate *estate.Estate) error {
	if aggregate == nil {
		return fmt.Errorf("aggregate cannot be nil")
	}

	aggregate.BeforeUpdate()

	filter := bson.M{"_id": aggregate.GetID().String()}
	opts := options.Replace().SetUpsert(false)

	result, err := r.collection.ReplaceOne(ctx, filter, aggregate, opts)
	if err != nil {
		return fmt.Errorf("could not save Estate aggregate: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("Estate aggregate with ID %s not found for update", aggregate.GetID().String())
	}

	return nil
}

// Delete removes the entire Estate aggregate from MongoDB.
// This automatically removes all child entities since they're part of the same document.
func (r *EstateMongoRepo) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id.String()}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("could not delete Estate aggregate: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("Estate aggregate with ID %s not found for deletion", id.String())
	}

	return nil
}

// Estate retrieves all Estate aggregates from MongoDB.
// Each document contains the complete aggregate (root + children).
func (r *EstateMongoRepo) Estate(ctx context.Context) ([]*estate.Estate, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not estate Estate aggregates: %w", err)
	}
	defer cursor.Close(ctx)

	var aggregates []*estate.Estate

	for cursor.Next(ctx) {
		var aggregate estate.Estate
		if err := cursor.Decode(&aggregate); err != nil {
			return nil, fmt.Errorf("could not decode Estate aggregate: %w", err)
		}
		aggregates = append(aggregates, &aggregate)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while estateing Estate aggregates: %w", err)
	}

	return aggregates, nil
}
