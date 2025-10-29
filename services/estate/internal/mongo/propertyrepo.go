package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pulap/pulap/services/estate/internal/config"
	"github.com/pulap/pulap/services/estate/internal/estate"
)

// PropertyRepo implements the estate.Repo interface using MongoDB.
// MongoDB is ideal for aggregates since each aggregate can be stored as a single document.
type PropertyRepo struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
	xparams    config.XParams
}

// NewPropertyRepo creates a new MongoDB repository for Property aggregates.
func NewPropertyRepo(xparams config.XParams) *PropertyRepo {
	return &PropertyRepo{
		xparams: xparams,
	}
}

// Start connects to MongoDB and initializes the collection.
func (r *PropertyRepo) Start(ctx context.Context) error {
	appCfg := r.xparams.Cfg()

	// Set default MongoDB configuration if not provided
	connString := appCfg.Database.MongoURL
	if connString == "" {
		connString = "mongodb://localhost:27017"
	}

	dbName := appCfg.Database.MongoDatabase
	if dbName == "" {
		dbName = "estate"
	}

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(connString).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("cannot connect to MongoDB: %w", err)
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("cannot ping MongoDB: %w", err)
	}

	r.client = client
	r.db = client.Database(dbName)
	r.collection = r.db.Collection("properties")

	r.xparams.Log().Infof("Connected to MongoDB: %s, database: %s", connString, dbName)
	return nil
}

// Stop closes the MongoDB connection.
func (r *PropertyRepo) Stop(ctx context.Context) error {
	if r.client != nil {
		if err := r.client.Disconnect(ctx); err != nil {
			return fmt.Errorf("cannot disconnect from MongoDB: %w", err)
		}
		r.xparams.Log().Info("Disconnected from MongoDB")
	}
	return nil
}

// Create creates a new Property aggregate in MongoDB.
// The entire aggregate is stored as a single document.
func (r *PropertyRepo) Create(ctx context.Context, property *estate.Property) error {
	if property == nil {
		return fmt.Errorf("property cannot be nil")
	}

	property.EnsureID()
	property.BeforeCreate()

	_, err := r.collection.InsertOne(ctx, property)
	if err != nil {
		return fmt.Errorf("could not create Property aggregate: %w", err)
	}

	return nil
}

// Get retrieves a complete Property aggregate by ID from MongoDB.
func (r *PropertyRepo) Get(ctx context.Context, id uuid.UUID) (*estate.Property, error) {
	var property estate.Property

	filter := bson.M{"_id": id.String()}
	err := r.collection.FindOne(ctx, filter).Decode(&property)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Property aggregate with ID %s not found", id.String())
		}
		return nil, fmt.Errorf("could not get Property aggregate: %w", err)
	}

	return &property, nil
}

// Save performs a unit-of-work save operation on the Property aggregate.
// In MongoDB, this is straightforward since the entire aggregate is replaced as one document.
func (r *PropertyRepo) Save(ctx context.Context, property *estate.Property) error {
	if property == nil {
		return fmt.Errorf("property cannot be nil")
	}

	property.BeforeUpdate()

	filter := bson.M{"_id": property.GetID().String()}
	opts := options.Replace().SetUpsert(false)

	result, err := r.collection.ReplaceOne(ctx, filter, property, opts)
	if err != nil {
		return fmt.Errorf("could not save Property aggregate: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("Property aggregate with ID %s not found for update", property.GetID().String())
	}

	return nil
}

// Delete removes the entire Property aggregate from MongoDB.
func (r *PropertyRepo) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id.String()}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("could not delete Property aggregate: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("Property aggregate with ID %s not found for deletion", id.String())
	}

	return nil
}

// List retrieves all Property aggregates from MongoDB.
func (r *PropertyRepo) List(ctx context.Context) ([]*estate.Property, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not list Property aggregates: %w", err)
	}
	defer cursor.Close(ctx)

	var properties []*estate.Property

	for cursor.Next(ctx) {
		var property estate.Property
		if err := cursor.Decode(&property); err != nil {
			return nil, fmt.Errorf("could not decode Property aggregate: %w", err)
		}
		properties = append(properties, &property)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while listing Property aggregates: %w", err)
	}

	return properties, nil
}

// ListByOwner retrieves all properties for a specific owner.
func (r *PropertyRepo) ListByOwner(ctx context.Context, ownerID string) ([]*estate.Property, error) {
	filter := bson.M{"owner_id": ownerID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not list properties by owner: %w", err)
	}
	defer cursor.Close(ctx)

	var properties []*estate.Property

	for cursor.Next(ctx) {
		var property estate.Property
		if err := cursor.Decode(&property); err != nil {
			return nil, fmt.Errorf("could not decode Property aggregate: %w", err)
		}
		properties = append(properties, &property)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while listing properties by owner: %w", err)
	}

	return properties, nil
}

// ListByStatus retrieves all properties with a specific status.
func (r *PropertyRepo) ListByStatus(ctx context.Context, status string) ([]*estate.Property, error) {
	filter := bson.M{"status": status}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not list properties by status: %w", err)
	}
	defer cursor.Close(ctx)

	var properties []*estate.Property

	for cursor.Next(ctx) {
		var property estate.Property
		if err := cursor.Decode(&property); err != nil {
			return nil, fmt.Errorf("could not decode Property aggregate: %w", err)
		}
		properties = append(properties, &property)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while listing properties by status: %w", err)
	}

	return properties, nil
}
