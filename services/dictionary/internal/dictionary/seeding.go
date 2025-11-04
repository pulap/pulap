package dictionary

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDictionarySeeds returns all seeds for the Dictionary service.
func GetDictionarySeeds() []Seed {
	return []Seed{
		{
			ID:          "2025-11-04_real_estate_dictionary",
			Description: "Load real estate dictionary (excluding geographic data, media metadata included)",
			Run:         seedRealEstateDictionary,
		},
	}
}

// seedRealEstateDictionary creates all sets and options for the real estate dictionary.
func seedRealEstateDictionary(ctx context.Context, db *mongo.Database) error {
	setsCollection := db.Collection("sets")
	optionsCollection := db.Collection("options")

	// Track set IDs for option creation
	setIDMap := make(map[string]string)

	// Create all sets
	// Create sets
	setID_estatecategory_en := uuid.New().String()
	setIDMap["estate_category:en"] = setID_estatecategory_en
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_category", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatecategory_en,
		"name":        "estate_category",
		"locale":      "en",
		"label":       "Estate Category",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatecategory_es := uuid.New().String()
	setIDMap["estate_category:es"] = setID_estatecategory_es
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_category", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatecategory_es,
		"name":        "estate_category",
		"locale":      "es",
		"label":       "Estate Category",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatecategory_pl := uuid.New().String()
	setIDMap["estate_category:pl"] = setID_estatecategory_pl
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_category", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatecategory_pl,
		"name":        "estate_category",
		"locale":      "pl",
		"label":       "Estate Category",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatetype_en := uuid.New().String()
	setIDMap["estate_type:en"] = setID_estatetype_en
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_type", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatetype_en,
		"name":        "estate_type",
		"locale":      "en",
		"label":       "Estate Type",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatetype_es := uuid.New().String()
	setIDMap["estate_type:es"] = setID_estatetype_es
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_type", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatetype_es,
		"name":        "estate_type",
		"locale":      "es",
		"label":       "Estate Type",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatetype_pl := uuid.New().String()
	setIDMap["estate_type:pl"] = setID_estatetype_pl
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_type", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatetype_pl,
		"name":        "estate_type",
		"locale":      "pl",
		"label":       "Estate Type",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatesubtype_en := uuid.New().String()
	setIDMap["estate_subtype:en"] = setID_estatesubtype_en
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_subtype", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatesubtype_en,
		"name":        "estate_subtype",
		"locale":      "en",
		"label":       "Estate Subtype",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatesubtype_es := uuid.New().String()
	setIDMap["estate_subtype:es"] = setID_estatesubtype_es
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_subtype", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatesubtype_es,
		"name":        "estate_subtype",
		"locale":      "es",
		"label":       "Estate Subtype",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatesubtype_pl := uuid.New().String()
	setIDMap["estate_subtype:pl"] = setID_estatesubtype_pl
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_subtype", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatesubtype_pl,
		"name":        "estate_subtype",
		"locale":      "pl",
		"label":       "Estate Subtype",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatestatus_en := uuid.New().String()
	setIDMap["estate_status:en"] = setID_estatestatus_en
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_status", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatestatus_en,
		"name":        "estate_status",
		"locale":      "en",
		"label":       "Estate Status",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatestatus_es := uuid.New().String()
	setIDMap["estate_status:es"] = setID_estatestatus_es
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_status", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatestatus_es,
		"name":        "estate_status",
		"locale":      "es",
		"label":       "Estate Status",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_estatestatus_pl := uuid.New().String()
	setIDMap["estate_status:pl"] = setID_estatestatus_pl
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "estate_status", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_estatestatus_pl,
		"name":        "estate_status",
		"locale":      "pl",
		"label":       "Estate Status",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_pricetype_en := uuid.New().String()
	setIDMap["price_type:en"] = setID_pricetype_en
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "price_type", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_pricetype_en,
		"name":        "price_type",
		"locale":      "en",
		"label":       "Price Type",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_pricetype_es := uuid.New().String()
	setIDMap["price_type:es"] = setID_pricetype_es
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "price_type", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_pricetype_es,
		"name":        "price_type",
		"locale":      "es",
		"label":       "Price Type",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_pricetype_pl := uuid.New().String()
	setIDMap["price_type:pl"] = setID_pricetype_pl
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "price_type", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_pricetype_pl,
		"name":        "price_type",
		"locale":      "pl",
		"label":       "Price Type",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_condition_en := uuid.New().String()
	setIDMap["condition:en"] = setID_condition_en
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "condition", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_condition_en,
		"name":        "condition",
		"locale":      "en",
		"label":       "Condition",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_condition_es := uuid.New().String()
	setIDMap["condition:es"] = setID_condition_es
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "condition", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_condition_es,
		"name":        "condition",
		"locale":      "es",
		"label":       "Condition",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	setID_condition_pl := uuid.New().String()
	setIDMap["condition:pl"] = setID_condition_pl
	_, _ = setsCollection.UpdateOne(ctx, bson.M{"name": "condition", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         setID_condition_pl,
		"name":        "condition",
		"locale":      "pl",
		"label":       "Condition",
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	// Track option IDs for parent relationships
	optionIDMap := make(map[string]string)

	// Create options without parents
	optID_1 := uuid.New().String()
	optionIDMap["estate_category:residential:en"] = optID_1
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:en"], "key": "residential", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_1,
		"set_id":      setIDMap["estate_category:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "res",
		"key":         "residential",
		"label":       "Residential",
		"description": "",
		"value":       "Residential",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_2 := uuid.New().String()
	optionIDMap["estate_category:residential:es"] = optID_2
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:es"], "key": "residential", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_2,
		"set_id":      setIDMap["estate_category:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "res",
		"key":         "residential",
		"label":       "Residencial",
		"description": "",
		"value":       "Residential",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_3 := uuid.New().String()
	optionIDMap["estate_category:residential:pl"] = optID_3
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:pl"], "key": "residential", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_3,
		"set_id":      setIDMap["estate_category:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "res",
		"key":         "residential",
		"label":       "Mieszkaniowe",
		"description": "",
		"value":       "Residential",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_4 := uuid.New().String()
	optionIDMap["estate_category:commercial:en"] = optID_4
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:en"], "key": "commercial", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_4,
		"set_id":      setIDMap["estate_category:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "com",
		"key":         "commercial",
		"label":       "Commercial",
		"description": "",
		"value":       "Commercial",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_5 := uuid.New().String()
	optionIDMap["estate_category:commercial:es"] = optID_5
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:es"], "key": "commercial", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_5,
		"set_id":      setIDMap["estate_category:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "com",
		"key":         "commercial",
		"label":       "Comercial",
		"description": "",
		"value":       "Commercial",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_6 := uuid.New().String()
	optionIDMap["estate_category:commercial:pl"] = optID_6
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:pl"], "key": "commercial", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_6,
		"set_id":      setIDMap["estate_category:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "com",
		"key":         "commercial",
		"label":       "Komercyjne",
		"description": "",
		"value":       "Commercial",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_7 := uuid.New().String()
	optionIDMap["estate_category:land:en"] = optID_7
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:en"], "key": "land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_7,
		"set_id":      setIDMap["estate_category:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "land",
		"key":         "land",
		"label":       "Land",
		"description": "",
		"value":       "Land",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_8 := uuid.New().String()
	optionIDMap["estate_category:land:es"] = optID_8
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:es"], "key": "land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_8,
		"set_id":      setIDMap["estate_category:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "land",
		"key":         "land",
		"label":       "Terreno",
		"description": "",
		"value":       "Land",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_9 := uuid.New().String()
	optionIDMap["estate_category:land:pl"] = optID_9
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:pl"], "key": "land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_9,
		"set_id":      setIDMap["estate_category:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "land",
		"key":         "land",
		"label":       "Grunt",
		"description": "",
		"value":       "Land",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_10 := uuid.New().String()
	optionIDMap["estate_category:agricultural:en"] = optID_10
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:en"], "key": "agricultural", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_10,
		"set_id":      setIDMap["estate_category:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "agr",
		"key":         "agricultural",
		"label":       "Agricultural",
		"description": "",
		"value":       "Agricultural",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_11 := uuid.New().String()
	optionIDMap["estate_category:agricultural:es"] = optID_11
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:es"], "key": "agricultural", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_11,
		"set_id":      setIDMap["estate_category:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "agr",
		"key":         "agricultural",
		"label":       "Agropecuario",
		"description": "",
		"value":       "Agricultural",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_12 := uuid.New().String()
	optionIDMap["estate_category:agricultural:pl"] = optID_12
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:pl"], "key": "agricultural", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_12,
		"set_id":      setIDMap["estate_category:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "agr",
		"key":         "agricultural",
		"label":       "Rolnicze",
		"description": "",
		"value":       "Agricultural",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_13 := uuid.New().String()
	optionIDMap["estate_category:mixed_use:en"] = optID_13
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:en"], "key": "mixed_use", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_13,
		"set_id":      setIDMap["estate_category:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "mix",
		"key":         "mixed_use",
		"label":       "Mixed-use",
		"description": "",
		"value":       "Mixed-use",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_14 := uuid.New().String()
	optionIDMap["estate_category:mixed_use:es"] = optID_14
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:es"], "key": "mixed_use", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_14,
		"set_id":      setIDMap["estate_category:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "mix",
		"key":         "mixed_use",
		"label":       "Uso mixto",
		"description": "",
		"value":       "Mixed-use",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_15 := uuid.New().String()
	optionIDMap["estate_category:mixed_use:pl"] = optID_15
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:pl"], "key": "mixed_use", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_15,
		"set_id":      setIDMap["estate_category:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "mix",
		"key":         "mixed_use",
		"label":       "Mieszane",
		"description": "",
		"value":       "Mixed-use",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_16 := uuid.New().String()
	optionIDMap["estate_category:special_purpose:en"] = optID_16
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:en"], "key": "special_purpose", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_16,
		"set_id":      setIDMap["estate_category:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "spc",
		"key":         "special_purpose",
		"label":       "Special Purpose",
		"description": "",
		"value":       "Special Purpose",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_17 := uuid.New().String()
	optionIDMap["estate_category:special_purpose:es"] = optID_17
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:es"], "key": "special_purpose", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_17,
		"set_id":      setIDMap["estate_category:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "spc",
		"key":         "special_purpose",
		"label":       "Uso especial",
		"description": "",
		"value":       "Special Purpose",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_18 := uuid.New().String()
	optionIDMap["estate_category:special_purpose:pl"] = optID_18
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_category:pl"], "key": "special_purpose", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_18,
		"set_id":      setIDMap["estate_category:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "spc",
		"key":         "special_purpose",
		"label":       "Obiekty specjalne",
		"description": "",
		"value":       "Special Purpose",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_19 := uuid.New().String()
	optionIDMap["estate_status:available:en"] = optID_19
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:en"], "key": "available", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_19,
		"set_id":      setIDMap["estate_status:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "ava",
		"key":         "available",
		"label":       "Available",
		"description": "",
		"value":       "Available",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_20 := uuid.New().String()
	optionIDMap["estate_status:available:es"] = optID_20
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:es"], "key": "available", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_20,
		"set_id":      setIDMap["estate_status:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "ava",
		"key":         "available",
		"label":       "Disponible",
		"description": "",
		"value":       "Available",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_21 := uuid.New().String()
	optionIDMap["estate_status:available:pl"] = optID_21
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:pl"], "key": "available", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_21,
		"set_id":      setIDMap["estate_status:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "ava",
		"key":         "available",
		"label":       "Dostępne",
		"description": "",
		"value":       "Available",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_22 := uuid.New().String()
	optionIDMap["estate_status:sold:en"] = optID_22
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:en"], "key": "sold", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_22,
		"set_id":      setIDMap["estate_status:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "sol",
		"key":         "sold",
		"label":       "Sold",
		"description": "",
		"value":       "Sold",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_23 := uuid.New().String()
	optionIDMap["estate_status:sold:es"] = optID_23
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:es"], "key": "sold", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_23,
		"set_id":      setIDMap["estate_status:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "sol",
		"key":         "sold",
		"label":       "Vendido",
		"description": "",
		"value":       "Sold",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_24 := uuid.New().String()
	optionIDMap["estate_status:sold:pl"] = optID_24
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:pl"], "key": "sold", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_24,
		"set_id":      setIDMap["estate_status:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "sol",
		"key":         "sold",
		"label":       "Sprzedane",
		"description": "",
		"value":       "Sold",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_25 := uuid.New().String()
	optionIDMap["estate_status:rented:en"] = optID_25
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:en"], "key": "rented", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_25,
		"set_id":      setIDMap["estate_status:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "ren",
		"key":         "rented",
		"label":       "Rented",
		"description": "",
		"value":       "Rented",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_26 := uuid.New().String()
	optionIDMap["estate_status:rented:es"] = optID_26
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:es"], "key": "rented", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_26,
		"set_id":      setIDMap["estate_status:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "ren",
		"key":         "rented",
		"label":       "Alquilado",
		"description": "",
		"value":       "Rented",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_27 := uuid.New().String()
	optionIDMap["estate_status:rented:pl"] = optID_27
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:pl"], "key": "rented", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_27,
		"set_id":      setIDMap["estate_status:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "ren",
		"key":         "rented",
		"label":       "Wynajęte",
		"description": "",
		"value":       "Rented",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_28 := uuid.New().String()
	optionIDMap["estate_status:reserved:en"] = optID_28
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:en"], "key": "reserved", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_28,
		"set_id":      setIDMap["estate_status:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "res",
		"key":         "reserved",
		"label":       "Reserved",
		"description": "",
		"value":       "Reserved",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_29 := uuid.New().String()
	optionIDMap["estate_status:reserved:es"] = optID_29
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:es"], "key": "reserved", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_29,
		"set_id":      setIDMap["estate_status:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "res",
		"key":         "reserved",
		"label":       "Reservado",
		"description": "",
		"value":       "Reserved",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_30 := uuid.New().String()
	optionIDMap["estate_status:reserved:pl"] = optID_30
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:pl"], "key": "reserved", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_30,
		"set_id":      setIDMap["estate_status:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "res",
		"key":         "reserved",
		"label":       "Zarezerwowane",
		"description": "",
		"value":       "Reserved",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_31 := uuid.New().String()
	optionIDMap["estate_status:draft:en"] = optID_31
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:en"], "key": "draft", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_31,
		"set_id":      setIDMap["estate_status:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "dra",
		"key":         "draft",
		"label":       "Draft",
		"description": "",
		"value":       "Draft",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_32 := uuid.New().String()
	optionIDMap["estate_status:draft:es"] = optID_32
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:es"], "key": "draft", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_32,
		"set_id":      setIDMap["estate_status:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "dra",
		"key":         "draft",
		"label":       "Borrador",
		"description": "",
		"value":       "Draft",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_33 := uuid.New().String()
	optionIDMap["estate_status:draft:pl"] = optID_33
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:pl"], "key": "draft", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_33,
		"set_id":      setIDMap["estate_status:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "dra",
		"key":         "draft",
		"label":       "Szkic",
		"description": "",
		"value":       "Draft",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_34 := uuid.New().String()
	optionIDMap["estate_status:inactive:en"] = optID_34
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:en"], "key": "inactive", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_34,
		"set_id":      setIDMap["estate_status:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "ina",
		"key":         "inactive",
		"label":       "Inactive",
		"description": "",
		"value":       "Inactive",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_35 := uuid.New().String()
	optionIDMap["estate_status:inactive:es"] = optID_35
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:es"], "key": "inactive", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_35,
		"set_id":      setIDMap["estate_status:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "ina",
		"key":         "inactive",
		"label":       "Inactivo",
		"description": "",
		"value":       "Inactive",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_36 := uuid.New().String()
	optionIDMap["estate_status:inactive:pl"] = optID_36
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_status:pl"], "key": "inactive", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_36,
		"set_id":      setIDMap["estate_status:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "ina",
		"key":         "inactive",
		"label":       "Nieaktywne",
		"description": "",
		"value":       "Inactive",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_37 := uuid.New().String()
	optionIDMap["price_type:sale:en"] = optID_37
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:en"], "key": "sale", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_37,
		"set_id":      setIDMap["price_type:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "sale",
		"key":         "sale",
		"label":       "Sale",
		"description": "",
		"value":       "Sale",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_38 := uuid.New().String()
	optionIDMap["price_type:sale:es"] = optID_38
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:es"], "key": "sale", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_38,
		"set_id":      setIDMap["price_type:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "sale",
		"key":         "sale",
		"label":       "Venta",
		"description": "",
		"value":       "Sale",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_39 := uuid.New().String()
	optionIDMap["price_type:sale:pl"] = optID_39
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:pl"], "key": "sale", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_39,
		"set_id":      setIDMap["price_type:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "sale",
		"key":         "sale",
		"label":       "Sprzedaż",
		"description": "",
		"value":       "Sale",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_40 := uuid.New().String()
	optionIDMap["price_type:rent_monthly:en"] = optID_40
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:en"], "key": "rent_monthly", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_40,
		"set_id":      setIDMap["price_type:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "rent_mon",
		"key":         "rent_monthly",
		"label":       "Rent (monthly)",
		"description": "",
		"value":       "Rent (monthly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_41 := uuid.New().String()
	optionIDMap["price_type:rent_monthly:es"] = optID_41
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:es"], "key": "rent_monthly", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_41,
		"set_id":      setIDMap["price_type:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "rent_mon",
		"key":         "rent_monthly",
		"label":       "Alquiler (mensual)",
		"description": "",
		"value":       "Rent (monthly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_42 := uuid.New().String()
	optionIDMap["price_type:rent_monthly:pl"] = optID_42
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:pl"], "key": "rent_monthly", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_42,
		"set_id":      setIDMap["price_type:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "rent_mon",
		"key":         "rent_monthly",
		"label":       "Wynajem (miesięcznie)",
		"description": "",
		"value":       "Rent (monthly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_43 := uuid.New().String()
	optionIDMap["price_type:rent_weekly:en"] = optID_43
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:en"], "key": "rent_weekly", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_43,
		"set_id":      setIDMap["price_type:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "rent_wee",
		"key":         "rent_weekly",
		"label":       "Rent (weekly)",
		"description": "",
		"value":       "Rent (weekly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_44 := uuid.New().String()
	optionIDMap["price_type:rent_weekly:es"] = optID_44
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:es"], "key": "rent_weekly", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_44,
		"set_id":      setIDMap["price_type:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "rent_wee",
		"key":         "rent_weekly",
		"label":       "Alquiler (semanal)",
		"description": "",
		"value":       "Rent (weekly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_45 := uuid.New().String()
	optionIDMap["price_type:rent_weekly:pl"] = optID_45
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:pl"], "key": "rent_weekly", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_45,
		"set_id":      setIDMap["price_type:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "rent_wee",
		"key":         "rent_weekly",
		"label":       "Wynajem (tygodniowo)",
		"description": "",
		"value":       "Rent (weekly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_46 := uuid.New().String()
	optionIDMap["price_type:rent_daily:en"] = optID_46
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:en"], "key": "rent_daily", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_46,
		"set_id":      setIDMap["price_type:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "rent_dai",
		"key":         "rent_daily",
		"label":       "Rent (daily)",
		"description": "",
		"value":       "Rent (daily)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_47 := uuid.New().String()
	optionIDMap["price_type:rent_daily:es"] = optID_47
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:es"], "key": "rent_daily", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_47,
		"set_id":      setIDMap["price_type:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "rent_dai",
		"key":         "rent_daily",
		"label":       "Alquiler (diario)",
		"description": "",
		"value":       "Rent (daily)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_48 := uuid.New().String()
	optionIDMap["price_type:rent_daily:pl"] = optID_48
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:pl"], "key": "rent_daily", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_48,
		"set_id":      setIDMap["price_type:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "rent_dai",
		"key":         "rent_daily",
		"label":       "Wynajem (dziennie)",
		"description": "",
		"value":       "Rent (daily)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_49 := uuid.New().String()
	optionIDMap["price_type:rent_yearly:en"] = optID_49
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:en"], "key": "rent_yearly", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_49,
		"set_id":      setIDMap["price_type:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "rent_yea",
		"key":         "rent_yearly",
		"label":       "Rent (yearly)",
		"description": "",
		"value":       "Rent (yearly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_50 := uuid.New().String()
	optionIDMap["price_type:rent_yearly:es"] = optID_50
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:es"], "key": "rent_yearly", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_50,
		"set_id":      setIDMap["price_type:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "rent_yea",
		"key":         "rent_yearly",
		"label":       "Alquiler (anual)",
		"description": "",
		"value":       "Rent (yearly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_51 := uuid.New().String()
	optionIDMap["price_type:rent_yearly:pl"] = optID_51
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["price_type:pl"], "key": "rent_yearly", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_51,
		"set_id":      setIDMap["price_type:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "rent_yea",
		"key":         "rent_yearly",
		"label":       "Wynajem (rocznie)",
		"description": "",
		"value":       "Rent (yearly)",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_52 := uuid.New().String()
	optionIDMap["condition:new:en"] = optID_52
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:en"], "key": "new", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_52,
		"set_id":      setIDMap["condition:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "new",
		"key":         "new",
		"label":       "New",
		"description": "",
		"value":       "New",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_53 := uuid.New().String()
	optionIDMap["condition:new:es"] = optID_53
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:es"], "key": "new", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_53,
		"set_id":      setIDMap["condition:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "new",
		"key":         "new",
		"label":       "Nuevo",
		"description": "",
		"value":       "New",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_54 := uuid.New().String()
	optionIDMap["condition:new:pl"] = optID_54
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:pl"], "key": "new", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_54,
		"set_id":      setIDMap["condition:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "new",
		"key":         "new",
		"label":       "Nowy",
		"description": "",
		"value":       "New",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_55 := uuid.New().String()
	optionIDMap["condition:excellent:en"] = optID_55
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:en"], "key": "excellent", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_55,
		"set_id":      setIDMap["condition:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "excellen",
		"key":         "excellent",
		"label":       "Excellent",
		"description": "",
		"value":       "Excellent",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_56 := uuid.New().String()
	optionIDMap["condition:excellent:es"] = optID_56
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:es"], "key": "excellent", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_56,
		"set_id":      setIDMap["condition:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "excellen",
		"key":         "excellent",
		"label":       "Excelente",
		"description": "",
		"value":       "Excellent",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_57 := uuid.New().String()
	optionIDMap["condition:excellent:pl"] = optID_57
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:pl"], "key": "excellent", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_57,
		"set_id":      setIDMap["condition:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "excellen",
		"key":         "excellent",
		"label":       "Doskonały",
		"description": "",
		"value":       "Excellent",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_58 := uuid.New().String()
	optionIDMap["condition:good:en"] = optID_58
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:en"], "key": "good", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_58,
		"set_id":      setIDMap["condition:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "good",
		"key":         "good",
		"label":       "Good",
		"description": "",
		"value":       "Good",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_59 := uuid.New().String()
	optionIDMap["condition:good:es"] = optID_59
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:es"], "key": "good", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_59,
		"set_id":      setIDMap["condition:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "good",
		"key":         "good",
		"label":       "Bueno",
		"description": "",
		"value":       "Good",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_60 := uuid.New().String()
	optionIDMap["condition:good:pl"] = optID_60
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:pl"], "key": "good", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_60,
		"set_id":      setIDMap["condition:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "good",
		"key":         "good",
		"label":       "Dobry",
		"description": "",
		"value":       "Good",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_61 := uuid.New().String()
	optionIDMap["condition:fair:en"] = optID_61
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:en"], "key": "fair", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_61,
		"set_id":      setIDMap["condition:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "fair",
		"key":         "fair",
		"label":       "Fair",
		"description": "",
		"value":       "Fair",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_62 := uuid.New().String()
	optionIDMap["condition:fair:es"] = optID_62
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:es"], "key": "fair", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_62,
		"set_id":      setIDMap["condition:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "fair",
		"key":         "fair",
		"label":       "Regular",
		"description": "",
		"value":       "Fair",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_63 := uuid.New().String()
	optionIDMap["condition:fair:pl"] = optID_63
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:pl"], "key": "fair", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_63,
		"set_id":      setIDMap["condition:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "fair",
		"key":         "fair",
		"label":       "Średni",
		"description": "",
		"value":       "Fair",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_64 := uuid.New().String()
	optionIDMap["condition:needs_work:en"] = optID_64
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:en"], "key": "needs_work", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_64,
		"set_id":      setIDMap["condition:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "needs_wo",
		"key":         "needs_work",
		"label":       "Needs work",
		"description": "",
		"value":       "Needs work",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_65 := uuid.New().String()
	optionIDMap["condition:needs_work:es"] = optID_65
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:es"], "key": "needs_work", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_65,
		"set_id":      setIDMap["condition:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "needs_wo",
		"key":         "needs_work",
		"label":       "Necesita trabajo",
		"description": "",
		"value":       "Needs work",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_66 := uuid.New().String()
	optionIDMap["condition:needs_work:pl"] = optID_66
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:pl"], "key": "needs_work", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_66,
		"set_id":      setIDMap["condition:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "needs_wo",
		"key":         "needs_work",
		"label":       "Wymaga prac",
		"description": "",
		"value":       "Needs work",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_67 := uuid.New().String()
	optionIDMap["condition:renovation:en"] = optID_67
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:en"], "key": "renovation", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_67,
		"set_id":      setIDMap["condition:en"],
		"parent_id":   nil,
		"locale":      "en",
		"short_code":  "renovati",
		"key":         "renovation",
		"label":       "Under renovation",
		"description": "",
		"value":       "Under renovation",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_68 := uuid.New().String()
	optionIDMap["condition:renovation:es"] = optID_68
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:es"], "key": "renovation", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_68,
		"set_id":      setIDMap["condition:es"],
		"parent_id":   nil,
		"locale":      "es",
		"short_code":  "renovati",
		"key":         "renovation",
		"label":       "En renovación",
		"description": "",
		"value":       "Under renovation",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	optID_69 := uuid.New().String()
	optionIDMap["condition:renovation:pl"] = optID_69
	_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["condition:pl"], "key": "renovation", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
		"_id":         optID_69,
		"set_id":      setIDMap["condition:pl"],
		"parent_id":   nil,
		"locale":      "pl",
		"short_code":  "renovati",
		"key":         "renovation",
		"label":       "W remoncie",
		"description": "",
		"value":       "Under renovation",
		"order":       0,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}}, options.Update().SetUpsert(true))

	// Create options with parents
	optID_70 := uuid.New().String()
	optionIDMap["estate_type:house:en"] = optID_70
	if parentID_70, ok := optionIDMap["estate_category:residential:en"]; ok {
		parentID_70_ptr := parentID_70
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "house", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_70,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_70_ptr,
			"locale":      "en",
			"short_code":  "house",
			"key":         "house",
			"label":       "House",
			"description": "",
			"value":       "House",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_71 := uuid.New().String()
	optionIDMap["estate_type:house:es"] = optID_71
	if parentID_71, ok := optionIDMap["estate_category:residential:es"]; ok {
		parentID_71_ptr := parentID_71
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "house", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_71,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_71_ptr,
			"locale":      "es",
			"short_code":  "house",
			"key":         "house",
			"label":       "Casa",
			"description": "",
			"value":       "House",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_72 := uuid.New().String()
	optionIDMap["estate_type:house:pl"] = optID_72
	if parentID_72, ok := optionIDMap["estate_category:residential:pl"]; ok {
		parentID_72_ptr := parentID_72
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "house", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_72,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_72_ptr,
			"locale":      "pl",
			"short_code":  "house",
			"key":         "house",
			"label":       "Dom",
			"description": "",
			"value":       "House",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_73 := uuid.New().String()
	optionIDMap["estate_type:apartment:en"] = optID_73
	if parentID_73, ok := optionIDMap["estate_category:residential:en"]; ok {
		parentID_73_ptr := parentID_73
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "apartment", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_73,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_73_ptr,
			"locale":      "en",
			"short_code":  "apt",
			"key":         "apartment",
			"label":       "Apartment",
			"description": "",
			"value":       "Apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_74 := uuid.New().String()
	optionIDMap["estate_type:apartment:es"] = optID_74
	if parentID_74, ok := optionIDMap["estate_category:residential:es"]; ok {
		parentID_74_ptr := parentID_74
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "apartment", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_74,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_74_ptr,
			"locale":      "es",
			"short_code":  "apt",
			"key":         "apartment",
			"label":       "Apartamento",
			"description": "",
			"value":       "Apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_75 := uuid.New().String()
	optionIDMap["estate_type:apartment:pl"] = optID_75
	if parentID_75, ok := optionIDMap["estate_category:residential:pl"]; ok {
		parentID_75_ptr := parentID_75
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "apartment", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_75,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_75_ptr,
			"locale":      "pl",
			"short_code":  "apt",
			"key":         "apartment",
			"label":       "Apartament",
			"description": "",
			"value":       "Apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_76 := uuid.New().String()
	optionIDMap["estate_type:multi_unit:en"] = optID_76
	if parentID_76, ok := optionIDMap["estate_category:residential:en"]; ok {
		parentID_76_ptr := parentID_76
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "multi_unit", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_76,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_76_ptr,
			"locale":      "en",
			"short_code":  "muf",
			"key":         "multi_unit",
			"label":       "Multi-unit",
			"description": "",
			"value":       "Multi-unit",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_77 := uuid.New().String()
	optionIDMap["estate_type:multi_unit:es"] = optID_77
	if parentID_77, ok := optionIDMap["estate_category:residential:es"]; ok {
		parentID_77_ptr := parentID_77
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "multi_unit", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_77,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_77_ptr,
			"locale":      "es",
			"short_code":  "muf",
			"key":         "multi_unit",
			"label":       "Multifamiliar",
			"description": "",
			"value":       "Multi-unit",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_78 := uuid.New().String()
	optionIDMap["estate_type:multi_unit:pl"] = optID_78
	if parentID_78, ok := optionIDMap["estate_category:residential:pl"]; ok {
		parentID_78_ptr := parentID_78
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "multi_unit", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_78,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_78_ptr,
			"locale":      "pl",
			"short_code":  "muf",
			"key":         "multi_unit",
			"label":       "Wielorodzinny",
			"description": "",
			"value":       "Multi-unit",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_79 := uuid.New().String()
	optionIDMap["estate_type:mobile_modular:en"] = optID_79
	if parentID_79, ok := optionIDMap["estate_category:residential:en"]; ok {
		parentID_79_ptr := parentID_79
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "mobile_modular", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_79,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_79_ptr,
			"locale":      "en",
			"short_code":  "mobmod",
			"key":         "mobile_modular",
			"label":       "Mobile/Modular",
			"description": "",
			"value":       "Mobile/Modular",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_80 := uuid.New().String()
	optionIDMap["estate_type:mobile_modular:es"] = optID_80
	if parentID_80, ok := optionIDMap["estate_category:residential:es"]; ok {
		parentID_80_ptr := parentID_80
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "mobile_modular", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_80,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_80_ptr,
			"locale":      "es",
			"short_code":  "mobmod",
			"key":         "mobile_modular",
			"label":       "Móvil/Modular",
			"description": "",
			"value":       "Mobile/Modular",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_81 := uuid.New().String()
	optionIDMap["estate_type:mobile_modular:pl"] = optID_81
	if parentID_81, ok := optionIDMap["estate_category:residential:pl"]; ok {
		parentID_81_ptr := parentID_81
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "mobile_modular", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_81,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_81_ptr,
			"locale":      "pl",
			"short_code":  "mobmod",
			"key":         "mobile_modular",
			"label":       "Mobilny/Modułowy",
			"description": "",
			"value":       "Mobile/Modular",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_82 := uuid.New().String()
	optionIDMap["estate_type:other_res:en"] = optID_82
	if parentID_82, ok := optionIDMap["estate_category:residential:en"]; ok {
		parentID_82_ptr := parentID_82
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "other_res", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_82,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_82_ptr,
			"locale":      "en",
			"short_code":  "resoth",
			"key":         "other_res",
			"label":       "Other (residential)",
			"description": "",
			"value":       "Other (residential)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_83 := uuid.New().String()
	optionIDMap["estate_type:other_res:es"] = optID_83
	if parentID_83, ok := optionIDMap["estate_category:residential:es"]; ok {
		parentID_83_ptr := parentID_83
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "other_res", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_83,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_83_ptr,
			"locale":      "es",
			"short_code":  "resoth",
			"key":         "other_res",
			"label":       "Otros (residencial)",
			"description": "",
			"value":       "Other (residential)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_84 := uuid.New().String()
	optionIDMap["estate_type:other_res:pl"] = optID_84
	if parentID_84, ok := optionIDMap["estate_category:residential:pl"]; ok {
		parentID_84_ptr := parentID_84
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "other_res", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_84,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_84_ptr,
			"locale":      "pl",
			"short_code":  "resoth",
			"key":         "other_res",
			"label":       "Inne (mieszkaniowe)",
			"description": "",
			"value":       "Other (residential)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_85 := uuid.New().String()
	optionIDMap["estate_type:office:en"] = optID_85
	if parentID_85, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_85_ptr := parentID_85
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "office", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_85,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_85_ptr,
			"locale":      "en",
			"short_code":  "off",
			"key":         "office",
			"label":       "Office",
			"description": "",
			"value":       "Office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_86 := uuid.New().String()
	optionIDMap["estate_type:office:es"] = optID_86
	if parentID_86, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_86_ptr := parentID_86
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "office", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_86,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_86_ptr,
			"locale":      "es",
			"short_code":  "off",
			"key":         "office",
			"label":       "Oficina",
			"description": "",
			"value":       "Office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_87 := uuid.New().String()
	optionIDMap["estate_type:office:pl"] = optID_87
	if parentID_87, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_87_ptr := parentID_87
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "office", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_87,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_87_ptr,
			"locale":      "pl",
			"short_code":  "off",
			"key":         "office",
			"label":       "Biuro",
			"description": "",
			"value":       "Office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_88 := uuid.New().String()
	optionIDMap["estate_type:retail:en"] = optID_88
	if parentID_88, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_88_ptr := parentID_88
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "retail", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_88,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_88_ptr,
			"locale":      "en",
			"short_code":  "rtl",
			"key":         "retail",
			"label":       "Retail",
			"description": "",
			"value":       "Retail",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_89 := uuid.New().String()
	optionIDMap["estate_type:retail:es"] = optID_89
	if parentID_89, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_89_ptr := parentID_89
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "retail", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_89,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_89_ptr,
			"locale":      "es",
			"short_code":  "rtl",
			"key":         "retail",
			"label":       "Retail / Comercio",
			"description": "",
			"value":       "Retail",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_90 := uuid.New().String()
	optionIDMap["estate_type:retail:pl"] = optID_90
	if parentID_90, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_90_ptr := parentID_90
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "retail", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_90,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_90_ptr,
			"locale":      "pl",
			"short_code":  "rtl",
			"key":         "retail",
			"label":       "Handel detaliczny",
			"description": "",
			"value":       "Retail",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_91 := uuid.New().String()
	optionIDMap["estate_type:hospitality:en"] = optID_91
	if parentID_91, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_91_ptr := parentID_91
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "hospitality", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_91,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_91_ptr,
			"locale":      "en",
			"short_code":  "hosp",
			"key":         "hospitality",
			"label":       "Hospitality",
			"description": "",
			"value":       "Hospitality",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_92 := uuid.New().String()
	optionIDMap["estate_type:hospitality:es"] = optID_92
	if parentID_92, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_92_ptr := parentID_92
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "hospitality", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_92,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_92_ptr,
			"locale":      "es",
			"short_code":  "hosp",
			"key":         "hospitality",
			"label":       "Hotelería",
			"description": "",
			"value":       "Hospitality",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_93 := uuid.New().String()
	optionIDMap["estate_type:hospitality:pl"] = optID_93
	if parentID_93, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_93_ptr := parentID_93
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "hospitality", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_93,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_93_ptr,
			"locale":      "pl",
			"short_code":  "hosp",
			"key":         "hospitality",
			"label":       "Hotelarstwo",
			"description": "",
			"value":       "Hospitality",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_94 := uuid.New().String()
	optionIDMap["estate_type:food_beverage:en"] = optID_94
	if parentID_94, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_94_ptr := parentID_94
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "food_beverage", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_94,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_94_ptr,
			"locale":      "en",
			"short_code":  "fnb",
			"key":         "food_beverage",
			"label":       "Food & Beverage",
			"description": "",
			"value":       "Food & Beverage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_95 := uuid.New().String()
	optionIDMap["estate_type:food_beverage:es"] = optID_95
	if parentID_95, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_95_ptr := parentID_95
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "food_beverage", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_95,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_95_ptr,
			"locale":      "es",
			"short_code":  "fnb",
			"key":         "food_beverage",
			"label":       "Gastronomía",
			"description": "",
			"value":       "Food & Beverage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_96 := uuid.New().String()
	optionIDMap["estate_type:food_beverage:pl"] = optID_96
	if parentID_96, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_96_ptr := parentID_96
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "food_beverage", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_96,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_96_ptr,
			"locale":      "pl",
			"short_code":  "fnb",
			"key":         "food_beverage",
			"label":       "Gastronomia",
			"description": "",
			"value":       "Food & Beverage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_97 := uuid.New().String()
	optionIDMap["estate_type:medical:en"] = optID_97
	if parentID_97, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_97_ptr := parentID_97
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "medical", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_97,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_97_ptr,
			"locale":      "en",
			"short_code":  "med",
			"key":         "medical",
			"label":       "Medical",
			"description": "",
			"value":       "Medical",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_98 := uuid.New().String()
	optionIDMap["estate_type:medical:es"] = optID_98
	if parentID_98, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_98_ptr := parentID_98
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "medical", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_98,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_98_ptr,
			"locale":      "es",
			"short_code":  "med",
			"key":         "medical",
			"label":       "Médico",
			"description": "",
			"value":       "Medical",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_99 := uuid.New().String()
	optionIDMap["estate_type:medical:pl"] = optID_99
	if parentID_99, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_99_ptr := parentID_99
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "medical", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_99,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_99_ptr,
			"locale":      "pl",
			"short_code":  "med",
			"key":         "medical",
			"label":       "Medyczne",
			"description": "",
			"value":       "Medical",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_100 := uuid.New().String()
	optionIDMap["estate_type:industrial:en"] = optID_100
	if parentID_100, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_100_ptr := parentID_100
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "industrial", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_100,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_100_ptr,
			"locale":      "en",
			"short_code":  "ind",
			"key":         "industrial",
			"label":       "Industrial",
			"description": "",
			"value":       "Industrial",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_101 := uuid.New().String()
	optionIDMap["estate_type:industrial:es"] = optID_101
	if parentID_101, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_101_ptr := parentID_101
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "industrial", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_101,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_101_ptr,
			"locale":      "es",
			"short_code":  "ind",
			"key":         "industrial",
			"label":       "Industrial",
			"description": "",
			"value":       "Industrial",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_102 := uuid.New().String()
	optionIDMap["estate_type:industrial:pl"] = optID_102
	if parentID_102, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_102_ptr := parentID_102
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "industrial", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_102,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_102_ptr,
			"locale":      "pl",
			"short_code":  "ind",
			"key":         "industrial",
			"label":       "Przemysłowe",
			"description": "",
			"value":       "Industrial",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_103 := uuid.New().String()
	optionIDMap["estate_type:special_com:en"] = optID_103
	if parentID_103, ok := optionIDMap["estate_category:commercial:en"]; ok {
		parentID_103_ptr := parentID_103
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "special_com", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_103,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_103_ptr,
			"locale":      "en",
			"short_code":  "comsp",
			"key":         "special_com",
			"label":       "Special (commercial)",
			"description": "",
			"value":       "Special (commercial)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_104 := uuid.New().String()
	optionIDMap["estate_type:special_com:es"] = optID_104
	if parentID_104, ok := optionIDMap["estate_category:commercial:es"]; ok {
		parentID_104_ptr := parentID_104
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "special_com", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_104,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_104_ptr,
			"locale":      "es",
			"short_code":  "comsp",
			"key":         "special_com",
			"label":       "Especial (comercial)",
			"description": "",
			"value":       "Special (commercial)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_105 := uuid.New().String()
	optionIDMap["estate_type:special_com:pl"] = optID_105
	if parentID_105, ok := optionIDMap["estate_category:commercial:pl"]; ok {
		parentID_105_ptr := parentID_105
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "special_com", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_105,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_105_ptr,
			"locale":      "pl",
			"short_code":  "comsp",
			"key":         "special_com",
			"label":       "Specjalne (komercyjne)",
			"description": "",
			"value":       "Special (commercial)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_106 := uuid.New().String()
	optionIDMap["estate_type:urban_land:en"] = optID_106
	if parentID_106, ok := optionIDMap["estate_category:land:en"]; ok {
		parentID_106_ptr := parentID_106
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "urban_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_106,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_106_ptr,
			"locale":      "en",
			"short_code":  "urb",
			"key":         "urban_land",
			"label":       "Urban",
			"description": "",
			"value":       "Urban",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_107 := uuid.New().String()
	optionIDMap["estate_type:urban_land:es"] = optID_107
	if parentID_107, ok := optionIDMap["estate_category:land:es"]; ok {
		parentID_107_ptr := parentID_107
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "urban_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_107,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_107_ptr,
			"locale":      "es",
			"short_code":  "urb",
			"key":         "urban_land",
			"label":       "Urbano",
			"description": "",
			"value":       "Urban",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_108 := uuid.New().String()
	optionIDMap["estate_type:urban_land:pl"] = optID_108
	if parentID_108, ok := optionIDMap["estate_category:land:pl"]; ok {
		parentID_108_ptr := parentID_108
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "urban_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_108,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_108_ptr,
			"locale":      "pl",
			"short_code":  "urb",
			"key":         "urban_land",
			"label":       "Miejski",
			"description": "",
			"value":       "Urban",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_109 := uuid.New().String()
	optionIDMap["estate_type:rural_land:en"] = optID_109
	if parentID_109, ok := optionIDMap["estate_category:land:en"]; ok {
		parentID_109_ptr := parentID_109
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "rural_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_109,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_109_ptr,
			"locale":      "en",
			"short_code":  "rur",
			"key":         "rural_land",
			"label":       "Rural",
			"description": "",
			"value":       "Rural",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_110 := uuid.New().String()
	optionIDMap["estate_type:rural_land:es"] = optID_110
	if parentID_110, ok := optionIDMap["estate_category:land:es"]; ok {
		parentID_110_ptr := parentID_110
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "rural_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_110,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_110_ptr,
			"locale":      "es",
			"short_code":  "rur",
			"key":         "rural_land",
			"label":       "Rural",
			"description": "",
			"value":       "Rural",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_111 := uuid.New().String()
	optionIDMap["estate_type:rural_land:pl"] = optID_111
	if parentID_111, ok := optionIDMap["estate_category:land:pl"]; ok {
		parentID_111_ptr := parentID_111
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "rural_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_111,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_111_ptr,
			"locale":      "pl",
			"short_code":  "rur",
			"key":         "rural_land",
			"label":       "Wiejski",
			"description": "",
			"value":       "Rural",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_112 := uuid.New().String()
	optionIDMap["estate_type:waterfront_land:en"] = optID_112
	if parentID_112, ok := optionIDMap["estate_category:land:en"]; ok {
		parentID_112_ptr := parentID_112
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "waterfront_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_112,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_112_ptr,
			"locale":      "en",
			"short_code":  "wfr",
			"key":         "waterfront_land",
			"label":       "Waterfront",
			"description": "",
			"value":       "Waterfront",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_113 := uuid.New().String()
	optionIDMap["estate_type:waterfront_land:es"] = optID_113
	if parentID_113, ok := optionIDMap["estate_category:land:es"]; ok {
		parentID_113_ptr := parentID_113
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "waterfront_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_113,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_113_ptr,
			"locale":      "es",
			"short_code":  "wfr",
			"key":         "waterfront_land",
			"label":       "Frente al agua",
			"description": "",
			"value":       "Waterfront",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_114 := uuid.New().String()
	optionIDMap["estate_type:waterfront_land:pl"] = optID_114
	if parentID_114, ok := optionIDMap["estate_category:land:pl"]; ok {
		parentID_114_ptr := parentID_114
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "waterfront_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_114,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_114_ptr,
			"locale":      "pl",
			"short_code":  "wfr",
			"key":         "waterfront_land",
			"label":       "Nabrzeżny",
			"description": "",
			"value":       "Waterfront",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_115 := uuid.New().String()
	optionIDMap["estate_type:special_land:en"] = optID_115
	if parentID_115, ok := optionIDMap["estate_category:land:en"]; ok {
		parentID_115_ptr := parentID_115
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "special_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_115,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_115_ptr,
			"locale":      "en",
			"short_code":  "lspec",
			"key":         "special_land",
			"label":       "Special",
			"description": "",
			"value":       "Special",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_116 := uuid.New().String()
	optionIDMap["estate_type:special_land:es"] = optID_116
	if parentID_116, ok := optionIDMap["estate_category:land:es"]; ok {
		parentID_116_ptr := parentID_116
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "special_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_116,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_116_ptr,
			"locale":      "es",
			"short_code":  "lspec",
			"key":         "special_land",
			"label":       "Especial",
			"description": "",
			"value":       "Special",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_117 := uuid.New().String()
	optionIDMap["estate_type:special_land:pl"] = optID_117
	if parentID_117, ok := optionIDMap["estate_category:land:pl"]; ok {
		parentID_117_ptr := parentID_117
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "special_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_117,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_117_ptr,
			"locale":      "pl",
			"short_code":  "lspec",
			"key":         "special_land",
			"label":       "Specjalny",
			"description": "",
			"value":       "Special",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_118 := uuid.New().String()
	optionIDMap["estate_type:farm:en"] = optID_118
	if parentID_118, ok := optionIDMap["estate_category:agricultural:en"]; ok {
		parentID_118_ptr := parentID_118
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "farm", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_118,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_118_ptr,
			"locale":      "en",
			"short_code":  "farm",
			"key":         "farm",
			"label":       "Farm",
			"description": "",
			"value":       "Farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_119 := uuid.New().String()
	optionIDMap["estate_type:farm:es"] = optID_119
	if parentID_119, ok := optionIDMap["estate_category:agricultural:es"]; ok {
		parentID_119_ptr := parentID_119
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "farm", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_119,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_119_ptr,
			"locale":      "es",
			"short_code":  "farm",
			"key":         "farm",
			"label":       "Granja",
			"description": "",
			"value":       "Farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_120 := uuid.New().String()
	optionIDMap["estate_type:farm:pl"] = optID_120
	if parentID_120, ok := optionIDMap["estate_category:agricultural:pl"]; ok {
		parentID_120_ptr := parentID_120
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "farm", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_120,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_120_ptr,
			"locale":      "pl",
			"short_code":  "farm",
			"key":         "farm",
			"label":       "Gospodarstwo",
			"description": "",
			"value":       "Farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_121 := uuid.New().String()
	optionIDMap["estate_type:ranch:en"] = optID_121
	if parentID_121, ok := optionIDMap["estate_category:agricultural:en"]; ok {
		parentID_121_ptr := parentID_121
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "ranch", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_121,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_121_ptr,
			"locale":      "en",
			"short_code":  "ranch",
			"key":         "ranch",
			"label":       "Ranch",
			"description": "",
			"value":       "Ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_122 := uuid.New().String()
	optionIDMap["estate_type:ranch:es"] = optID_122
	if parentID_122, ok := optionIDMap["estate_category:agricultural:es"]; ok {
		parentID_122_ptr := parentID_122
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "ranch", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_122,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_122_ptr,
			"locale":      "es",
			"short_code":  "ranch",
			"key":         "ranch",
			"label":       "Estancia / Rancho",
			"description": "",
			"value":       "Ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_123 := uuid.New().String()
	optionIDMap["estate_type:ranch:pl"] = optID_123
	if parentID_123, ok := optionIDMap["estate_category:agricultural:pl"]; ok {
		parentID_123_ptr := parentID_123
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "ranch", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_123,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_123_ptr,
			"locale":      "pl",
			"short_code":  "ranch",
			"key":         "ranch",
			"label":       "Ranczo",
			"description": "",
			"value":       "Ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_124 := uuid.New().String()
	optionIDMap["estate_type:agri_specialty:en"] = optID_124
	if parentID_124, ok := optionIDMap["estate_category:agricultural:en"]; ok {
		parentID_124_ptr := parentID_124
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "agri_specialty", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_124,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_124_ptr,
			"locale":      "en",
			"short_code":  "agsp",
			"key":         "agri_specialty",
			"label":       "Specialty",
			"description": "",
			"value":       "Specialty",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_125 := uuid.New().String()
	optionIDMap["estate_type:agri_specialty:es"] = optID_125
	if parentID_125, ok := optionIDMap["estate_category:agricultural:es"]; ok {
		parentID_125_ptr := parentID_125
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "agri_specialty", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_125,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_125_ptr,
			"locale":      "es",
			"short_code":  "agsp",
			"key":         "agri_specialty",
			"label":       "Especialidad",
			"description": "",
			"value":       "Specialty",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_126 := uuid.New().String()
	optionIDMap["estate_type:agri_specialty:pl"] = optID_126
	if parentID_126, ok := optionIDMap["estate_category:agricultural:pl"]; ok {
		parentID_126_ptr := parentID_126
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "agri_specialty", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_126,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_126_ptr,
			"locale":      "pl",
			"short_code":  "agsp",
			"key":         "agri_specialty",
			"label":       "Specjalistyczne",
			"description": "",
			"value":       "Specialty",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_127 := uuid.New().String()
	optionIDMap["estate_type:live_work:en"] = optID_127
	if parentID_127, ok := optionIDMap["estate_category:mixed_use:en"]; ok {
		parentID_127_ptr := parentID_127
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "live_work", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_127,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_127_ptr,
			"locale":      "en",
			"short_code":  "lw",
			"key":         "live_work",
			"label":       "Live/work",
			"description": "",
			"value":       "Live/work",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_128 := uuid.New().String()
	optionIDMap["estate_type:live_work:es"] = optID_128
	if parentID_128, ok := optionIDMap["estate_category:mixed_use:es"]; ok {
		parentID_128_ptr := parentID_128
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "live_work", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_128,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_128_ptr,
			"locale":      "es",
			"short_code":  "lw",
			"key":         "live_work",
			"label":       "Vivienda + Trabajo",
			"description": "",
			"value":       "Live/work",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_129 := uuid.New().String()
	optionIDMap["estate_type:live_work:pl"] = optID_129
	if parentID_129, ok := optionIDMap["estate_category:mixed_use:pl"]; ok {
		parentID_129_ptr := parentID_129
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "live_work", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_129,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_129_ptr,
			"locale":      "pl",
			"short_code":  "lw",
			"key":         "live_work",
			"label":       "Mieszkanie + Praca",
			"description": "",
			"value":       "Live/work",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_130 := uuid.New().String()
	optionIDMap["estate_type:mixed_building:en"] = optID_130
	if parentID_130, ok := optionIDMap["estate_category:mixed_use:en"]; ok {
		parentID_130_ptr := parentID_130
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "mixed_building", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_130,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_130_ptr,
			"locale":      "en",
			"short_code":  "mixbld",
			"key":         "mixed_building",
			"label":       "Mixed building",
			"description": "",
			"value":       "Mixed building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_131 := uuid.New().String()
	optionIDMap["estate_type:mixed_building:es"] = optID_131
	if parentID_131, ok := optionIDMap["estate_category:mixed_use:es"]; ok {
		parentID_131_ptr := parentID_131
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "mixed_building", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_131,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_131_ptr,
			"locale":      "es",
			"short_code":  "mixbld",
			"key":         "mixed_building",
			"label":       "Edificio mixto",
			"description": "",
			"value":       "Mixed building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_132 := uuid.New().String()
	optionIDMap["estate_type:mixed_building:pl"] = optID_132
	if parentID_132, ok := optionIDMap["estate_category:mixed_use:pl"]; ok {
		parentID_132_ptr := parentID_132
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "mixed_building", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_132,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_132_ptr,
			"locale":      "pl",
			"short_code":  "mixbld",
			"key":         "mixed_building",
			"label":       "Budynek mieszany",
			"description": "",
			"value":       "Mixed building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_133 := uuid.New().String()
	optionIDMap["estate_type:transportation:en"] = optID_133
	if parentID_133, ok := optionIDMap["estate_category:special_purpose:en"]; ok {
		parentID_133_ptr := parentID_133
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "transportation", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_133,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_133_ptr,
			"locale":      "en",
			"short_code":  "transp",
			"key":         "transportation",
			"label":       "Transportation",
			"description": "",
			"value":       "Transportation",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_134 := uuid.New().String()
	optionIDMap["estate_type:transportation:es"] = optID_134
	if parentID_134, ok := optionIDMap["estate_category:special_purpose:es"]; ok {
		parentID_134_ptr := parentID_134
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "transportation", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_134,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_134_ptr,
			"locale":      "es",
			"short_code":  "transp",
			"key":         "transportation",
			"label":       "Transporte",
			"description": "",
			"value":       "Transportation",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_135 := uuid.New().String()
	optionIDMap["estate_type:transportation:pl"] = optID_135
	if parentID_135, ok := optionIDMap["estate_category:special_purpose:pl"]; ok {
		parentID_135_ptr := parentID_135
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "transportation", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_135,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_135_ptr,
			"locale":      "pl",
			"short_code":  "transp",
			"key":         "transportation",
			"label":       "Transport",
			"description": "",
			"value":       "Transportation",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_136 := uuid.New().String()
	optionIDMap["estate_type:utilities:en"] = optID_136
	if parentID_136, ok := optionIDMap["estate_category:special_purpose:en"]; ok {
		parentID_136_ptr := parentID_136
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "utilities", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_136,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_136_ptr,
			"locale":      "en",
			"short_code":  "util",
			"key":         "utilities",
			"label":       "Utility / Infrastructure",
			"description": "",
			"value":       "Utility / Infrastructure",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_137 := uuid.New().String()
	optionIDMap["estate_type:utilities:es"] = optID_137
	if parentID_137, ok := optionIDMap["estate_category:special_purpose:es"]; ok {
		parentID_137_ptr := parentID_137
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "utilities", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_137,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_137_ptr,
			"locale":      "es",
			"short_code":  "util",
			"key":         "utilities",
			"label":       "Servicios / Infraestructura",
			"description": "",
			"value":       "Utility / Infrastructure",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_138 := uuid.New().String()
	optionIDMap["estate_type:utilities:pl"] = optID_138
	if parentID_138, ok := optionIDMap["estate_category:special_purpose:pl"]; ok {
		parentID_138_ptr := parentID_138
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "utilities", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_138,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_138_ptr,
			"locale":      "pl",
			"short_code":  "util",
			"key":         "utilities",
			"label":       "Usługi / Infrastruktura",
			"description": "",
			"value":       "Utility / Infrastructure",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_139 := uuid.New().String()
	optionIDMap["estate_type:institutional:en"] = optID_139
	if parentID_139, ok := optionIDMap["estate_category:special_purpose:en"]; ok {
		parentID_139_ptr := parentID_139
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "institutional", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_139,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_139_ptr,
			"locale":      "en",
			"short_code":  "inst",
			"key":         "institutional",
			"label":       "Institutional",
			"description": "",
			"value":       "Institutional",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_140 := uuid.New().String()
	optionIDMap["estate_type:institutional:es"] = optID_140
	if parentID_140, ok := optionIDMap["estate_category:special_purpose:es"]; ok {
		parentID_140_ptr := parentID_140
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "institutional", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_140,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_140_ptr,
			"locale":      "es",
			"short_code":  "inst",
			"key":         "institutional",
			"label":       "Institucional",
			"description": "",
			"value":       "Institutional",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_141 := uuid.New().String()
	optionIDMap["estate_type:institutional:pl"] = optID_141
	if parentID_141, ok := optionIDMap["estate_category:special_purpose:pl"]; ok {
		parentID_141_ptr := parentID_141
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "institutional", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_141,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_141_ptr,
			"locale":      "pl",
			"short_code":  "inst",
			"key":         "institutional",
			"label":       "Instytucjonalne",
			"description": "",
			"value":       "Institutional",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_142 := uuid.New().String()
	optionIDMap["estate_type:recreational:en"] = optID_142
	if parentID_142, ok := optionIDMap["estate_category:special_purpose:en"]; ok {
		parentID_142_ptr := parentID_142
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:en"], "key": "recreational", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_142,
			"set_id":      setIDMap["estate_type:en"],
			"parent_id":   &parentID_142_ptr,
			"locale":      "en",
			"short_code":  "rec",
			"key":         "recreational",
			"label":       "Recreational",
			"description": "",
			"value":       "Recreational",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_143 := uuid.New().String()
	optionIDMap["estate_type:recreational:es"] = optID_143
	if parentID_143, ok := optionIDMap["estate_category:special_purpose:es"]; ok {
		parentID_143_ptr := parentID_143
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:es"], "key": "recreational", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_143,
			"set_id":      setIDMap["estate_type:es"],
			"parent_id":   &parentID_143_ptr,
			"locale":      "es",
			"short_code":  "rec",
			"key":         "recreational",
			"label":       "Recreativo",
			"description": "",
			"value":       "Recreational",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_144 := uuid.New().String()
	optionIDMap["estate_type:recreational:pl"] = optID_144
	if parentID_144, ok := optionIDMap["estate_category:special_purpose:pl"]; ok {
		parentID_144_ptr := parentID_144
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_type:pl"], "key": "recreational", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_144,
			"set_id":      setIDMap["estate_type:pl"],
			"parent_id":   &parentID_144_ptr,
			"locale":      "pl",
			"short_code":  "rec",
			"key":         "recreational",
			"label":       "Rekreacyjne",
			"description": "",
			"value":       "Recreational",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_145 := uuid.New().String()
	optionIDMap["estate_subtype:detached_house:en"] = optID_145
	if parentID_145, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_145_ptr := parentID_145
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "detached_house", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_145,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_145_ptr,
			"locale":      "en",
			"short_code":  "detached",
			"key":         "detached_house",
			"label":       "Detached house",
			"description": "",
			"value":       "Detached house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_146 := uuid.New().String()
	optionIDMap["estate_subtype:detached_house:es"] = optID_146
	if parentID_146, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_146_ptr := parentID_146
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "detached_house", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_146,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_146_ptr,
			"locale":      "es",
			"short_code":  "detached",
			"key":         "detached_house",
			"label":       "Casa independiente",
			"description": "",
			"value":       "Detached house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_147 := uuid.New().String()
	optionIDMap["estate_subtype:detached_house:pl"] = optID_147
	if parentID_147, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_147_ptr := parentID_147
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "detached_house", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_147,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_147_ptr,
			"locale":      "pl",
			"short_code":  "detached",
			"key":         "detached_house",
			"label":       "Dom wolnostojący",
			"description": "",
			"value":       "Detached house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_148 := uuid.New().String()
	optionIDMap["estate_subtype:semi_detached:en"] = optID_148
	if parentID_148, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_148_ptr := parentID_148
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "semi_detached", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_148,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_148_ptr,
			"locale":      "en",
			"short_code":  "semi_det",
			"key":         "semi_detached",
			"label":       "Semi-detached house",
			"description": "",
			"value":       "Semi-detached house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_149 := uuid.New().String()
	optionIDMap["estate_subtype:semi_detached:es"] = optID_149
	if parentID_149, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_149_ptr := parentID_149
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "semi_detached", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_149,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_149_ptr,
			"locale":      "es",
			"short_code":  "semi_det",
			"key":         "semi_detached",
			"label":       "Casa pareada",
			"description": "",
			"value":       "Semi-detached house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_150 := uuid.New().String()
	optionIDMap["estate_subtype:semi_detached:pl"] = optID_150
	if parentID_150, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_150_ptr := parentID_150
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "semi_detached", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_150,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_150_ptr,
			"locale":      "pl",
			"short_code":  "semi_det",
			"key":         "semi_detached",
			"label":       "Bliźniak",
			"description": "",
			"value":       "Semi-detached house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_151 := uuid.New().String()
	optionIDMap["estate_subtype:bungalow:en"] = optID_151
	if parentID_151, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_151_ptr := parentID_151
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "bungalow", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_151,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_151_ptr,
			"locale":      "en",
			"short_code":  "bungalow",
			"key":         "bungalow",
			"label":       "Bungalow",
			"description": "",
			"value":       "Bungalow",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_152 := uuid.New().String()
	optionIDMap["estate_subtype:bungalow:es"] = optID_152
	if parentID_152, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_152_ptr := parentID_152
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "bungalow", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_152,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_152_ptr,
			"locale":      "es",
			"short_code":  "bungalow",
			"key":         "bungalow",
			"label":       "Bungaló",
			"description": "",
			"value":       "Bungalow",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_153 := uuid.New().String()
	optionIDMap["estate_subtype:bungalow:pl"] = optID_153
	if parentID_153, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_153_ptr := parentID_153
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "bungalow", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_153,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_153_ptr,
			"locale":      "pl",
			"short_code":  "bungalow",
			"key":         "bungalow",
			"label":       "Bungalow",
			"description": "",
			"value":       "Bungalow",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_154 := uuid.New().String()
	optionIDMap["estate_subtype:ranch_style:en"] = optID_154
	if parentID_154, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_154_ptr := parentID_154
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "ranch_style", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_154,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_154_ptr,
			"locale":      "en",
			"short_code":  "ranch_st",
			"key":         "ranch_style",
			"label":       "Ranch-style home",
			"description": "",
			"value":       "Ranch-style home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_155 := uuid.New().String()
	optionIDMap["estate_subtype:ranch_style:es"] = optID_155
	if parentID_155, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_155_ptr := parentID_155
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "ranch_style", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_155,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_155_ptr,
			"locale":      "es",
			"short_code":  "ranch_st",
			"key":         "ranch_style",
			"label":       "Casa estilo rancho",
			"description": "",
			"value":       "Ranch-style home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_156 := uuid.New().String()
	optionIDMap["estate_subtype:ranch_style:pl"] = optID_156
	if parentID_156, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_156_ptr := parentID_156
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "ranch_style", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_156,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_156_ptr,
			"locale":      "pl",
			"short_code":  "ranch_st",
			"key":         "ranch_style",
			"label":       "Dom typu ranch",
			"description": "",
			"value":       "Ranch-style home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_157 := uuid.New().String()
	optionIDMap["estate_subtype:cottage:en"] = optID_157
	if parentID_157, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_157_ptr := parentID_157
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "cottage", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_157,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_157_ptr,
			"locale":      "en",
			"short_code":  "cottage",
			"key":         "cottage",
			"label":       "Cottage",
			"description": "",
			"value":       "Cottage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_158 := uuid.New().String()
	optionIDMap["estate_subtype:cottage:es"] = optID_158
	if parentID_158, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_158_ptr := parentID_158
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "cottage", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_158,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_158_ptr,
			"locale":      "es",
			"short_code":  "cottage",
			"key":         "cottage",
			"label":       "Cabaña rural",
			"description": "",
			"value":       "Cottage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_159 := uuid.New().String()
	optionIDMap["estate_subtype:cottage:pl"] = optID_159
	if parentID_159, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_159_ptr := parentID_159
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "cottage", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_159,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_159_ptr,
			"locale":      "pl",
			"short_code":  "cottage",
			"key":         "cottage",
			"label":       "Domek wiejski",
			"description": "",
			"value":       "Cottage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_160 := uuid.New().String()
	optionIDMap["estate_subtype:chalet:en"] = optID_160
	if parentID_160, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_160_ptr := parentID_160
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "chalet", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_160,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_160_ptr,
			"locale":      "en",
			"short_code":  "chalet",
			"key":         "chalet",
			"label":       "Chalet",
			"description": "",
			"value":       "Chalet",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_161 := uuid.New().String()
	optionIDMap["estate_subtype:chalet:es"] = optID_161
	if parentID_161, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_161_ptr := parentID_161
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "chalet", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_161,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_161_ptr,
			"locale":      "es",
			"short_code":  "chalet",
			"key":         "chalet",
			"label":       "Chalet",
			"description": "",
			"value":       "Chalet",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_162 := uuid.New().String()
	optionIDMap["estate_subtype:chalet:pl"] = optID_162
	if parentID_162, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_162_ptr := parentID_162
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "chalet", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_162,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_162_ptr,
			"locale":      "pl",
			"short_code":  "chalet",
			"key":         "chalet",
			"label":       "Domek górski",
			"description": "",
			"value":       "Chalet",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_163 := uuid.New().String()
	optionIDMap["estate_subtype:cabin:en"] = optID_163
	if parentID_163, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_163_ptr := parentID_163
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "cabin", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_163,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_163_ptr,
			"locale":      "en",
			"short_code":  "cabin",
			"key":         "cabin",
			"label":       "Cabin",
			"description": "",
			"value":       "Cabin",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_164 := uuid.New().String()
	optionIDMap["estate_subtype:cabin:es"] = optID_164
	if parentID_164, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_164_ptr := parentID_164
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "cabin", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_164,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_164_ptr,
			"locale":      "es",
			"short_code":  "cabin",
			"key":         "cabin",
			"label":       "Cabaña",
			"description": "",
			"value":       "Cabin",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_165 := uuid.New().String()
	optionIDMap["estate_subtype:cabin:pl"] = optID_165
	if parentID_165, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_165_ptr := parentID_165
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "cabin", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_165,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_165_ptr,
			"locale":      "pl",
			"short_code":  "cabin",
			"key":         "cabin",
			"label":       "Chatka",
			"description": "",
			"value":       "Cabin",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_166 := uuid.New().String()
	optionIDMap["estate_subtype:eco_home:en"] = optID_166
	if parentID_166, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_166_ptr := parentID_166
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "eco_home", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_166,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_166_ptr,
			"locale":      "en",
			"short_code":  "eco_home",
			"key":         "eco_home",
			"label":       "Eco-home",
			"description": "",
			"value":       "Eco-home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_167 := uuid.New().String()
	optionIDMap["estate_subtype:eco_home:es"] = optID_167
	if parentID_167, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_167_ptr := parentID_167
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "eco_home", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_167,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_167_ptr,
			"locale":      "es",
			"short_code":  "eco_home",
			"key":         "eco_home",
			"label":       "Casa ecológica",
			"description": "",
			"value":       "Eco-home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_168 := uuid.New().String()
	optionIDMap["estate_subtype:eco_home:pl"] = optID_168
	if parentID_168, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_168_ptr := parentID_168
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "eco_home", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_168,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_168_ptr,
			"locale":      "pl",
			"short_code":  "eco_home",
			"key":         "eco_home",
			"label":       "Dom ekologiczny",
			"description": "",
			"value":       "Eco-home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_169 := uuid.New().String()
	optionIDMap["estate_subtype:smart_home:en"] = optID_169
	if parentID_169, ok := optionIDMap["estate_type:house:en"]; ok {
		parentID_169_ptr := parentID_169
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "smart_home", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_169,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_169_ptr,
			"locale":      "en",
			"short_code":  "smart_ho",
			"key":         "smart_home",
			"label":       "Smart home",
			"description": "",
			"value":       "Smart home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_170 := uuid.New().String()
	optionIDMap["estate_subtype:smart_home:es"] = optID_170
	if parentID_170, ok := optionIDMap["estate_type:house:es"]; ok {
		parentID_170_ptr := parentID_170
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "smart_home", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_170,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_170_ptr,
			"locale":      "es",
			"short_code":  "smart_ho",
			"key":         "smart_home",
			"label":       "Casa inteligente",
			"description": "",
			"value":       "Smart home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_171 := uuid.New().String()
	optionIDMap["estate_subtype:smart_home:pl"] = optID_171
	if parentID_171, ok := optionIDMap["estate_type:house:pl"]; ok {
		parentID_171_ptr := parentID_171
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "smart_home", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_171,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_171_ptr,
			"locale":      "pl",
			"short_code":  "smart_ho",
			"key":         "smart_home",
			"label":       "Inteligentny dom",
			"description": "",
			"value":       "Smart home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_172 := uuid.New().String()
	optionIDMap["estate_subtype:studio:en"] = optID_172
	if parentID_172, ok := optionIDMap["estate_type:apartment:en"]; ok {
		parentID_172_ptr := parentID_172
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "studio", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_172,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_172_ptr,
			"locale":      "en",
			"short_code":  "studio",
			"key":         "studio",
			"label":       "Studio apartment",
			"description": "",
			"value":       "Studio apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_173 := uuid.New().String()
	optionIDMap["estate_subtype:studio:es"] = optID_173
	if parentID_173, ok := optionIDMap["estate_type:apartment:es"]; ok {
		parentID_173_ptr := parentID_173
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "studio", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_173,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_173_ptr,
			"locale":      "es",
			"short_code":  "studio",
			"key":         "studio",
			"label":       "Estudio",
			"description": "",
			"value":       "Studio apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_174 := uuid.New().String()
	optionIDMap["estate_subtype:studio:pl"] = optID_174
	if parentID_174, ok := optionIDMap["estate_type:apartment:pl"]; ok {
		parentID_174_ptr := parentID_174
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "studio", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_174,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_174_ptr,
			"locale":      "pl",
			"short_code":  "studio",
			"key":         "studio",
			"label":       "Kawalerka",
			"description": "",
			"value":       "Studio apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_175 := uuid.New().String()
	optionIDMap["estate_subtype:basement_apt:en"] = optID_175
	if parentID_175, ok := optionIDMap["estate_type:apartment:en"]; ok {
		parentID_175_ptr := parentID_175
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "basement_apt", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_175,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_175_ptr,
			"locale":      "en",
			"short_code":  "basement",
			"key":         "basement_apt",
			"label":       "Basement apartment",
			"description": "",
			"value":       "Basement apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_176 := uuid.New().String()
	optionIDMap["estate_subtype:basement_apt:es"] = optID_176
	if parentID_176, ok := optionIDMap["estate_type:apartment:es"]; ok {
		parentID_176_ptr := parentID_176
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "basement_apt", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_176,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_176_ptr,
			"locale":      "es",
			"short_code":  "basement",
			"key":         "basement_apt",
			"label":       "Departamento en sótano",
			"description": "",
			"value":       "Basement apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_177 := uuid.New().String()
	optionIDMap["estate_subtype:basement_apt:pl"] = optID_177
	if parentID_177, ok := optionIDMap["estate_type:apartment:pl"]; ok {
		parentID_177_ptr := parentID_177
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "basement_apt", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_177,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_177_ptr,
			"locale":      "pl",
			"short_code":  "basement",
			"key":         "basement_apt",
			"label":       "Mieszkanie w piwnicy",
			"description": "",
			"value":       "Basement apartment",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_178 := uuid.New().String()
	optionIDMap["estate_subtype:penthouse:en"] = optID_178
	if parentID_178, ok := optionIDMap["estate_type:apartment:en"]; ok {
		parentID_178_ptr := parentID_178
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "penthouse", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_178,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_178_ptr,
			"locale":      "en",
			"short_code":  "penthous",
			"key":         "penthouse",
			"label":       "Penthouse",
			"description": "",
			"value":       "Penthouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_179 := uuid.New().String()
	optionIDMap["estate_subtype:penthouse:es"] = optID_179
	if parentID_179, ok := optionIDMap["estate_type:apartment:es"]; ok {
		parentID_179_ptr := parentID_179
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "penthouse", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_179,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_179_ptr,
			"locale":      "es",
			"short_code":  "penthous",
			"key":         "penthouse",
			"label":       "Ático",
			"description": "",
			"value":       "Penthouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_180 := uuid.New().String()
	optionIDMap["estate_subtype:penthouse:pl"] = optID_180
	if parentID_180, ok := optionIDMap["estate_type:apartment:pl"]; ok {
		parentID_180_ptr := parentID_180
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "penthouse", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_180,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_180_ptr,
			"locale":      "pl",
			"short_code":  "penthous",
			"key":         "penthouse",
			"label":       "Penthouse",
			"description": "",
			"value":       "Penthouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_181 := uuid.New().String()
	optionIDMap["estate_subtype:loft:en"] = optID_181
	if parentID_181, ok := optionIDMap["estate_type:apartment:en"]; ok {
		parentID_181_ptr := parentID_181
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "loft", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_181,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_181_ptr,
			"locale":      "en",
			"short_code":  "loft",
			"key":         "loft",
			"label":       "Loft",
			"description": "",
			"value":       "Loft",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_182 := uuid.New().String()
	optionIDMap["estate_subtype:loft:es"] = optID_182
	if parentID_182, ok := optionIDMap["estate_type:apartment:es"]; ok {
		parentID_182_ptr := parentID_182
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "loft", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_182,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_182_ptr,
			"locale":      "es",
			"short_code":  "loft",
			"key":         "loft",
			"label":       "Loft",
			"description": "",
			"value":       "Loft",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_183 := uuid.New().String()
	optionIDMap["estate_subtype:loft:pl"] = optID_183
	if parentID_183, ok := optionIDMap["estate_type:apartment:pl"]; ok {
		parentID_183_ptr := parentID_183
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "loft", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_183,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_183_ptr,
			"locale":      "pl",
			"short_code":  "loft",
			"key":         "loft",
			"label":       "Loft",
			"description": "",
			"value":       "Loft",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_184 := uuid.New().String()
	optionIDMap["estate_subtype:condo:en"] = optID_184
	if parentID_184, ok := optionIDMap["estate_type:apartment:en"]; ok {
		parentID_184_ptr := parentID_184
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "condo", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_184,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_184_ptr,
			"locale":      "en",
			"short_code":  "condo",
			"key":         "condo",
			"label":       "Condominium",
			"description": "",
			"value":       "Condominium",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_185 := uuid.New().String()
	optionIDMap["estate_subtype:condo:es"] = optID_185
	if parentID_185, ok := optionIDMap["estate_type:apartment:es"]; ok {
		parentID_185_ptr := parentID_185
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "condo", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_185,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_185_ptr,
			"locale":      "es",
			"short_code":  "condo",
			"key":         "condo",
			"label":       "Condominio",
			"description": "",
			"value":       "Condominium",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_186 := uuid.New().String()
	optionIDMap["estate_subtype:condo:pl"] = optID_186
	if parentID_186, ok := optionIDMap["estate_type:apartment:pl"]; ok {
		parentID_186_ptr := parentID_186
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "condo", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_186,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_186_ptr,
			"locale":      "pl",
			"short_code":  "condo",
			"key":         "condo",
			"label":       "Mieszkanie własnościowe",
			"description": "",
			"value":       "Condominium",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_187 := uuid.New().String()
	optionIDMap["estate_subtype:coop:en"] = optID_187
	if parentID_187, ok := optionIDMap["estate_type:apartment:en"]; ok {
		parentID_187_ptr := parentID_187
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "coop", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_187,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_187_ptr,
			"locale":      "en",
			"short_code":  "coop",
			"key":         "coop",
			"label":       "Co-op unit",
			"description": "",
			"value":       "Co-op unit",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_188 := uuid.New().String()
	optionIDMap["estate_subtype:coop:es"] = optID_188
	if parentID_188, ok := optionIDMap["estate_type:apartment:es"]; ok {
		parentID_188_ptr := parentID_188
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "coop", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_188,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_188_ptr,
			"locale":      "es",
			"short_code":  "coop",
			"key":         "coop",
			"label":       "Unidad cooperativa",
			"description": "",
			"value":       "Co-op unit",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_189 := uuid.New().String()
	optionIDMap["estate_subtype:coop:pl"] = optID_189
	if parentID_189, ok := optionIDMap["estate_type:apartment:pl"]; ok {
		parentID_189_ptr := parentID_189
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "coop", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_189,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_189_ptr,
			"locale":      "pl",
			"short_code":  "coop",
			"key":         "coop",
			"label":       "Mieszkanie spółdzielcze",
			"description": "",
			"value":       "Co-op unit",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_190 := uuid.New().String()
	optionIDMap["estate_subtype:duplex:en"] = optID_190
	if parentID_190, ok := optionIDMap["estate_type:multi_unit:en"]; ok {
		parentID_190_ptr := parentID_190
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "duplex", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_190,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_190_ptr,
			"locale":      "en",
			"short_code":  "duplex",
			"key":         "duplex",
			"label":       "Duplex",
			"description": "",
			"value":       "Duplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_191 := uuid.New().String()
	optionIDMap["estate_subtype:duplex:es"] = optID_191
	if parentID_191, ok := optionIDMap["estate_type:multi_unit:es"]; ok {
		parentID_191_ptr := parentID_191
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "duplex", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_191,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_191_ptr,
			"locale":      "es",
			"short_code":  "duplex",
			"key":         "duplex",
			"label":       "Dúplex",
			"description": "",
			"value":       "Duplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_192 := uuid.New().String()
	optionIDMap["estate_subtype:duplex:pl"] = optID_192
	if parentID_192, ok := optionIDMap["estate_type:multi_unit:pl"]; ok {
		parentID_192_ptr := parentID_192
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "duplex", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_192,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_192_ptr,
			"locale":      "pl",
			"short_code":  "duplex",
			"key":         "duplex",
			"label":       "Bliźniak/dwupoziomowe",
			"description": "",
			"value":       "Duplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_193 := uuid.New().String()
	optionIDMap["estate_subtype:triplex:en"] = optID_193
	if parentID_193, ok := optionIDMap["estate_type:multi_unit:en"]; ok {
		parentID_193_ptr := parentID_193
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "triplex", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_193,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_193_ptr,
			"locale":      "en",
			"short_code":  "triplex",
			"key":         "triplex",
			"label":       "Triplex",
			"description": "",
			"value":       "Triplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_194 := uuid.New().String()
	optionIDMap["estate_subtype:triplex:es"] = optID_194
	if parentID_194, ok := optionIDMap["estate_type:multi_unit:es"]; ok {
		parentID_194_ptr := parentID_194
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "triplex", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_194,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_194_ptr,
			"locale":      "es",
			"short_code":  "triplex",
			"key":         "triplex",
			"label":       "Tríplex",
			"description": "",
			"value":       "Triplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_195 := uuid.New().String()
	optionIDMap["estate_subtype:triplex:pl"] = optID_195
	if parentID_195, ok := optionIDMap["estate_type:multi_unit:pl"]; ok {
		parentID_195_ptr := parentID_195
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "triplex", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_195,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_195_ptr,
			"locale":      "pl",
			"short_code":  "triplex",
			"key":         "triplex",
			"label":       "Trójpoziomowe",
			"description": "",
			"value":       "Triplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_196 := uuid.New().String()
	optionIDMap["estate_subtype:fourplex:en"] = optID_196
	if parentID_196, ok := optionIDMap["estate_type:multi_unit:en"]; ok {
		parentID_196_ptr := parentID_196
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "fourplex", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_196,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_196_ptr,
			"locale":      "en",
			"short_code":  "fourplex",
			"key":         "fourplex",
			"label":       "Fourplex",
			"description": "",
			"value":       "Fourplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_197 := uuid.New().String()
	optionIDMap["estate_subtype:fourplex:es"] = optID_197
	if parentID_197, ok := optionIDMap["estate_type:multi_unit:es"]; ok {
		parentID_197_ptr := parentID_197
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "fourplex", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_197,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_197_ptr,
			"locale":      "es",
			"short_code":  "fourplex",
			"key":         "fourplex",
			"label":       "Cuádruplex",
			"description": "",
			"value":       "Fourplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_198 := uuid.New().String()
	optionIDMap["estate_subtype:fourplex:pl"] = optID_198
	if parentID_198, ok := optionIDMap["estate_type:multi_unit:pl"]; ok {
		parentID_198_ptr := parentID_198
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "fourplex", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_198,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_198_ptr,
			"locale":      "pl",
			"short_code":  "fourplex",
			"key":         "fourplex",
			"label":       "Czterolokalowe",
			"description": "",
			"value":       "Fourplex",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_199 := uuid.New().String()
	optionIDMap["estate_subtype:townhouse:en"] = optID_199
	if parentID_199, ok := optionIDMap["estate_type:multi_unit:en"]; ok {
		parentID_199_ptr := parentID_199
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "townhouse", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_199,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_199_ptr,
			"locale":      "en",
			"short_code":  "townhous",
			"key":         "townhouse",
			"label":       "Townhouse",
			"description": "",
			"value":       "Townhouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_200 := uuid.New().String()
	optionIDMap["estate_subtype:townhouse:es"] = optID_200
	if parentID_200, ok := optionIDMap["estate_type:multi_unit:es"]; ok {
		parentID_200_ptr := parentID_200
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "townhouse", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_200,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_200_ptr,
			"locale":      "es",
			"short_code":  "townhous",
			"key":         "townhouse",
			"label":       "Casa en hilera",
			"description": "",
			"value":       "Townhouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_201 := uuid.New().String()
	optionIDMap["estate_subtype:townhouse:pl"] = optID_201
	if parentID_201, ok := optionIDMap["estate_type:multi_unit:pl"]; ok {
		parentID_201_ptr := parentID_201
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "townhouse", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_201,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_201_ptr,
			"locale":      "pl",
			"short_code":  "townhous",
			"key":         "townhouse",
			"label":       "Szeregowiec",
			"description": "",
			"value":       "Townhouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_202 := uuid.New().String()
	optionIDMap["estate_subtype:row_house:en"] = optID_202
	if parentID_202, ok := optionIDMap["estate_type:multi_unit:en"]; ok {
		parentID_202_ptr := parentID_202
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "row_house", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_202,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_202_ptr,
			"locale":      "en",
			"short_code":  "row_hous",
			"key":         "row_house",
			"label":       "Row house",
			"description": "",
			"value":       "Row house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_203 := uuid.New().String()
	optionIDMap["estate_subtype:row_house:es"] = optID_203
	if parentID_203, ok := optionIDMap["estate_type:multi_unit:es"]; ok {
		parentID_203_ptr := parentID_203
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "row_house", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_203,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_203_ptr,
			"locale":      "es",
			"short_code":  "row_hous",
			"key":         "row_house",
			"label":       "Casa adosada",
			"description": "",
			"value":       "Row house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_204 := uuid.New().String()
	optionIDMap["estate_subtype:row_house:pl"] = optID_204
	if parentID_204, ok := optionIDMap["estate_type:multi_unit:pl"]; ok {
		parentID_204_ptr := parentID_204
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "row_house", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_204,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_204_ptr,
			"locale":      "pl",
			"short_code":  "row_hous",
			"key":         "row_house",
			"label":       "Dom szeregowy",
			"description": "",
			"value":       "Row house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_205 := uuid.New().String()
	optionIDMap["estate_subtype:mobile_home:en"] = optID_205
	if parentID_205, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_205_ptr := parentID_205
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "mobile_home", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_205,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_205_ptr,
			"locale":      "en",
			"short_code":  "mobile_h",
			"key":         "mobile_home",
			"label":       "Mobile home",
			"description": "",
			"value":       "Mobile home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_206 := uuid.New().String()
	optionIDMap["estate_subtype:mobile_home:es"] = optID_206
	if parentID_206, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_206_ptr := parentID_206
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "mobile_home", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_206,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_206_ptr,
			"locale":      "es",
			"short_code":  "mobile_h",
			"key":         "mobile_home",
			"label":       "Casa móvil",
			"description": "",
			"value":       "Mobile home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_207 := uuid.New().String()
	optionIDMap["estate_subtype:mobile_home:pl"] = optID_207
	if parentID_207, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_207_ptr := parentID_207
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "mobile_home", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_207,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_207_ptr,
			"locale":      "pl",
			"short_code":  "mobile_h",
			"key":         "mobile_home",
			"label":       "Dom mobilny",
			"description": "",
			"value":       "Mobile home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_208 := uuid.New().String()
	optionIDMap["estate_subtype:modular_home:en"] = optID_208
	if parentID_208, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_208_ptr := parentID_208
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "modular_home", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_208,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_208_ptr,
			"locale":      "en",
			"short_code":  "modular_",
			"key":         "modular_home",
			"label":       "Modular home",
			"description": "",
			"value":       "Modular home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_209 := uuid.New().String()
	optionIDMap["estate_subtype:modular_home:es"] = optID_209
	if parentID_209, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_209_ptr := parentID_209
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "modular_home", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_209,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_209_ptr,
			"locale":      "es",
			"short_code":  "modular_",
			"key":         "modular_home",
			"label":       "Casa modular",
			"description": "",
			"value":       "Modular home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_210 := uuid.New().String()
	optionIDMap["estate_subtype:modular_home:pl"] = optID_210
	if parentID_210, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_210_ptr := parentID_210
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "modular_home", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_210,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_210_ptr,
			"locale":      "pl",
			"short_code":  "modular_",
			"key":         "modular_home",
			"label":       "Dom modułowy",
			"description": "",
			"value":       "Modular home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_211 := uuid.New().String()
	optionIDMap["estate_subtype:park_model:en"] = optID_211
	if parentID_211, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_211_ptr := parentID_211
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "park_model", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_211,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_211_ptr,
			"locale":      "en",
			"short_code":  "park_mod",
			"key":         "park_model",
			"label":       "Park model home",
			"description": "",
			"value":       "Park model home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_212 := uuid.New().String()
	optionIDMap["estate_subtype:park_model:es"] = optID_212
	if parentID_212, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_212_ptr := parentID_212
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "park_model", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_212,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_212_ptr,
			"locale":      "es",
			"short_code":  "park_mod",
			"key":         "park_model",
			"label":       "Casa de parque",
			"description": "",
			"value":       "Park model home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_213 := uuid.New().String()
	optionIDMap["estate_subtype:park_model:pl"] = optID_213
	if parentID_213, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_213_ptr := parentID_213
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "park_model", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_213,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_213_ptr,
			"locale":      "pl",
			"short_code":  "park_mod",
			"key":         "park_model",
			"label":       "Domek kempingowy",
			"description": "",
			"value":       "Park model home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_214 := uuid.New().String()
	optionIDMap["estate_subtype:tiny_house:en"] = optID_214
	if parentID_214, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_214_ptr := parentID_214
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "tiny_house", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_214,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_214_ptr,
			"locale":      "en",
			"short_code":  "tiny_hou",
			"key":         "tiny_house",
			"label":       "Tiny house",
			"description": "",
			"value":       "Tiny house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_215 := uuid.New().String()
	optionIDMap["estate_subtype:tiny_house:es"] = optID_215
	if parentID_215, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_215_ptr := parentID_215
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "tiny_house", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_215,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_215_ptr,
			"locale":      "es",
			"short_code":  "tiny_hou",
			"key":         "tiny_house",
			"label":       "Mini casa",
			"description": "",
			"value":       "Tiny house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_216 := uuid.New().String()
	optionIDMap["estate_subtype:tiny_house:pl"] = optID_216
	if parentID_216, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_216_ptr := parentID_216
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "tiny_house", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_216,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_216_ptr,
			"locale":      "pl",
			"short_code":  "tiny_hou",
			"key":         "tiny_house",
			"label":       "Tiny house",
			"description": "",
			"value":       "Tiny house",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_217 := uuid.New().String()
	optionIDMap["estate_subtype:floating_home:en"] = optID_217
	if parentID_217, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_217_ptr := parentID_217
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "floating_home", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_217,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_217_ptr,
			"locale":      "en",
			"short_code":  "floating",
			"key":         "floating_home",
			"label":       "Floating home",
			"description": "",
			"value":       "Floating home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_218 := uuid.New().String()
	optionIDMap["estate_subtype:floating_home:es"] = optID_218
	if parentID_218, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_218_ptr := parentID_218
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "floating_home", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_218,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_218_ptr,
			"locale":      "es",
			"short_code":  "floating",
			"key":         "floating_home",
			"label":       "Casa flotante",
			"description": "",
			"value":       "Floating home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_219 := uuid.New().String()
	optionIDMap["estate_subtype:floating_home:pl"] = optID_219
	if parentID_219, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_219_ptr := parentID_219
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "floating_home", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_219,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_219_ptr,
			"locale":      "pl",
			"short_code":  "floating",
			"key":         "floating_home",
			"label":       "Dom pływający",
			"description": "",
			"value":       "Floating home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_220 := uuid.New().String()
	optionIDMap["estate_subtype:houseboat:en"] = optID_220
	if parentID_220, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_220_ptr := parentID_220
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "houseboat", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_220,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_220_ptr,
			"locale":      "en",
			"short_code":  "houseboa",
			"key":         "houseboat",
			"label":       "Houseboat",
			"description": "",
			"value":       "Houseboat",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_221 := uuid.New().String()
	optionIDMap["estate_subtype:houseboat:es"] = optID_221
	if parentID_221, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_221_ptr := parentID_221
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "houseboat", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_221,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_221_ptr,
			"locale":      "es",
			"short_code":  "houseboa",
			"key":         "houseboat",
			"label":       "Casa barco",
			"description": "",
			"value":       "Houseboat",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_222 := uuid.New().String()
	optionIDMap["estate_subtype:houseboat:pl"] = optID_222
	if parentID_222, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_222_ptr := parentID_222
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "houseboat", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_222,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_222_ptr,
			"locale":      "pl",
			"short_code":  "houseboa",
			"key":         "houseboat",
			"label":       "Łódź mieszkalna",
			"description": "",
			"value":       "Houseboat",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_223 := uuid.New().String()
	optionIDMap["estate_subtype:yurt:en"] = optID_223
	if parentID_223, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_223_ptr := parentID_223
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "yurt", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_223,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_223_ptr,
			"locale":      "en",
			"short_code":  "yurt",
			"key":         "yurt",
			"label":       "Yurt",
			"description": "",
			"value":       "Yurt",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_224 := uuid.New().String()
	optionIDMap["estate_subtype:yurt:es"] = optID_224
	if parentID_224, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_224_ptr := parentID_224
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "yurt", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_224,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_224_ptr,
			"locale":      "es",
			"short_code":  "yurt",
			"key":         "yurt",
			"label":       "Yurta",
			"description": "",
			"value":       "Yurt",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_225 := uuid.New().String()
	optionIDMap["estate_subtype:yurt:pl"] = optID_225
	if parentID_225, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_225_ptr := parentID_225
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "yurt", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_225,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_225_ptr,
			"locale":      "pl",
			"short_code":  "yurt",
			"key":         "yurt",
			"label":       "Jurta",
			"description": "",
			"value":       "Yurt",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_226 := uuid.New().String()
	optionIDMap["estate_subtype:treehouse:en"] = optID_226
	if parentID_226, ok := optionIDMap["estate_type:mobile_modular:en"]; ok {
		parentID_226_ptr := parentID_226
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "treehouse", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_226,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_226_ptr,
			"locale":      "en",
			"short_code":  "treehous",
			"key":         "treehouse",
			"label":       "Treehouse",
			"description": "",
			"value":       "Treehouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_227 := uuid.New().String()
	optionIDMap["estate_subtype:treehouse:es"] = optID_227
	if parentID_227, ok := optionIDMap["estate_type:mobile_modular:es"]; ok {
		parentID_227_ptr := parentID_227
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "treehouse", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_227,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_227_ptr,
			"locale":      "es",
			"short_code":  "treehous",
			"key":         "treehouse",
			"label":       "Casa del árbol",
			"description": "",
			"value":       "Treehouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_228 := uuid.New().String()
	optionIDMap["estate_subtype:treehouse:pl"] = optID_228
	if parentID_228, ok := optionIDMap["estate_type:mobile_modular:pl"]; ok {
		parentID_228_ptr := parentID_228
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "treehouse", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_228,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_228_ptr,
			"locale":      "pl",
			"short_code":  "treehous",
			"key":         "treehouse",
			"label":       "Domek na drzewie",
			"description": "",
			"value":       "Treehouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_229 := uuid.New().String()
	optionIDMap["estate_subtype:office_bldg:en"] = optID_229
	if parentID_229, ok := optionIDMap["estate_type:office:en"]; ok {
		parentID_229_ptr := parentID_229
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "office_bldg", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_229,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_229_ptr,
			"locale":      "en",
			"short_code":  "office_b",
			"key":         "office_bldg",
			"label":       "Office building",
			"description": "",
			"value":       "Office building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_230 := uuid.New().String()
	optionIDMap["estate_subtype:office_bldg:es"] = optID_230
	if parentID_230, ok := optionIDMap["estate_type:office:es"]; ok {
		parentID_230_ptr := parentID_230
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "office_bldg", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_230,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_230_ptr,
			"locale":      "es",
			"short_code":  "office_b",
			"key":         "office_bldg",
			"label":       "Edificio de oficinas",
			"description": "",
			"value":       "Office building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_231 := uuid.New().String()
	optionIDMap["estate_subtype:office_bldg:pl"] = optID_231
	if parentID_231, ok := optionIDMap["estate_type:office:pl"]; ok {
		parentID_231_ptr := parentID_231
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "office_bldg", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_231,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_231_ptr,
			"locale":      "pl",
			"short_code":  "office_b",
			"key":         "office_bldg",
			"label":       "Biurowiec",
			"description": "",
			"value":       "Office building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_232 := uuid.New().String()
	optionIDMap["estate_subtype:exec_suite:en"] = optID_232
	if parentID_232, ok := optionIDMap["estate_type:office:en"]; ok {
		parentID_232_ptr := parentID_232
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "exec_suite", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_232,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_232_ptr,
			"locale":      "en",
			"short_code":  "exec_sui",
			"key":         "exec_suite",
			"label":       "Executive suite",
			"description": "",
			"value":       "Executive suite",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_233 := uuid.New().String()
	optionIDMap["estate_subtype:exec_suite:es"] = optID_233
	if parentID_233, ok := optionIDMap["estate_type:office:es"]; ok {
		parentID_233_ptr := parentID_233
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "exec_suite", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_233,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_233_ptr,
			"locale":      "es",
			"short_code":  "exec_sui",
			"key":         "exec_suite",
			"label":       "Oficina ejecutiva",
			"description": "",
			"value":       "Executive suite",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_234 := uuid.New().String()
	optionIDMap["estate_subtype:exec_suite:pl"] = optID_234
	if parentID_234, ok := optionIDMap["estate_type:office:pl"]; ok {
		parentID_234_ptr := parentID_234
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "exec_suite", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_234,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_234_ptr,
			"locale":      "pl",
			"short_code":  "exec_sui",
			"key":         "exec_suite",
			"label":       "Biuro serwisowane",
			"description": "",
			"value":       "Executive suite",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_235 := uuid.New().String()
	optionIDMap["estate_subtype:coworking:en"] = optID_235
	if parentID_235, ok := optionIDMap["estate_type:office:en"]; ok {
		parentID_235_ptr := parentID_235
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "coworking", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_235,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_235_ptr,
			"locale":      "en",
			"short_code":  "coworkin",
			"key":         "coworking",
			"label":       "Co-working space",
			"description": "",
			"value":       "Co-working space",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_236 := uuid.New().String()
	optionIDMap["estate_subtype:coworking:es"] = optID_236
	if parentID_236, ok := optionIDMap["estate_type:office:es"]; ok {
		parentID_236_ptr := parentID_236
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "coworking", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_236,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_236_ptr,
			"locale":      "es",
			"short_code":  "coworkin",
			"key":         "coworking",
			"label":       "Espacio de co-working",
			"description": "",
			"value":       "Co-working space",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_237 := uuid.New().String()
	optionIDMap["estate_subtype:coworking:pl"] = optID_237
	if parentID_237, ok := optionIDMap["estate_type:office:pl"]; ok {
		parentID_237_ptr := parentID_237
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "coworking", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_237,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_237_ptr,
			"locale":      "pl",
			"short_code":  "coworkin",
			"key":         "coworking",
			"label":       "Coworking",
			"description": "",
			"value":       "Co-working space",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_238 := uuid.New().String()
	optionIDMap["estate_subtype:retail_store:en"] = optID_238
	if parentID_238, ok := optionIDMap["estate_type:retail:en"]; ok {
		parentID_238_ptr := parentID_238
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "retail_store", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_238,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_238_ptr,
			"locale":      "en",
			"short_code":  "retail_s",
			"key":         "retail_store",
			"label":       "Retail store",
			"description": "",
			"value":       "Retail store",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_239 := uuid.New().String()
	optionIDMap["estate_subtype:retail_store:es"] = optID_239
	if parentID_239, ok := optionIDMap["estate_type:retail:es"]; ok {
		parentID_239_ptr := parentID_239
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "retail_store", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_239,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_239_ptr,
			"locale":      "es",
			"short_code":  "retail_s",
			"key":         "retail_store",
			"label":       "Tienda",
			"description": "",
			"value":       "Retail store",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_240 := uuid.New().String()
	optionIDMap["estate_subtype:retail_store:pl"] = optID_240
	if parentID_240, ok := optionIDMap["estate_type:retail:pl"]; ok {
		parentID_240_ptr := parentID_240
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "retail_store", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_240,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_240_ptr,
			"locale":      "pl",
			"short_code":  "retail_s",
			"key":         "retail_store",
			"label":       "Sklep detaliczny",
			"description": "",
			"value":       "Retail store",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_241 := uuid.New().String()
	optionIDMap["estate_subtype:shopping_mall:en"] = optID_241
	if parentID_241, ok := optionIDMap["estate_type:retail:en"]; ok {
		parentID_241_ptr := parentID_241
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "shopping_mall", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_241,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_241_ptr,
			"locale":      "en",
			"short_code":  "shopping",
			"key":         "shopping_mall",
			"label":       "Shopping mall",
			"description": "",
			"value":       "Shopping mall",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_242 := uuid.New().String()
	optionIDMap["estate_subtype:shopping_mall:es"] = optID_242
	if parentID_242, ok := optionIDMap["estate_type:retail:es"]; ok {
		parentID_242_ptr := parentID_242
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "shopping_mall", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_242,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_242_ptr,
			"locale":      "es",
			"short_code":  "shopping",
			"key":         "shopping_mall",
			"label":       "Centro comercial",
			"description": "",
			"value":       "Shopping mall",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_243 := uuid.New().String()
	optionIDMap["estate_subtype:shopping_mall:pl"] = optID_243
	if parentID_243, ok := optionIDMap["estate_type:retail:pl"]; ok {
		parentID_243_ptr := parentID_243
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "shopping_mall", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_243,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_243_ptr,
			"locale":      "pl",
			"short_code":  "shopping",
			"key":         "shopping_mall",
			"label":       "Centrum handlowe",
			"description": "",
			"value":       "Shopping mall",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_244 := uuid.New().String()
	optionIDMap["estate_subtype:strip_mall:en"] = optID_244
	if parentID_244, ok := optionIDMap["estate_type:retail:en"]; ok {
		parentID_244_ptr := parentID_244
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "strip_mall", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_244,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_244_ptr,
			"locale":      "en",
			"short_code":  "strip_ma",
			"key":         "strip_mall",
			"label":       "Strip mall",
			"description": "",
			"value":       "Strip mall",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_245 := uuid.New().String()
	optionIDMap["estate_subtype:strip_mall:es"] = optID_245
	if parentID_245, ok := optionIDMap["estate_type:retail:es"]; ok {
		parentID_245_ptr := parentID_245
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "strip_mall", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_245,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_245_ptr,
			"locale":      "es",
			"short_code":  "strip_ma",
			"key":         "strip_mall",
			"label":       "Strip mall",
			"description": "",
			"value":       "Strip mall",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_246 := uuid.New().String()
	optionIDMap["estate_subtype:strip_mall:pl"] = optID_246
	if parentID_246, ok := optionIDMap["estate_type:retail:pl"]; ok {
		parentID_246_ptr := parentID_246
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "strip_mall", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_246,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_246_ptr,
			"locale":      "pl",
			"short_code":  "strip_ma",
			"key":         "strip_mall",
			"label":       "Park handlowy",
			"description": "",
			"value":       "Strip mall",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_247 := uuid.New().String()
	optionIDMap["estate_subtype:showroom:en"] = optID_247
	if parentID_247, ok := optionIDMap["estate_type:retail:en"]; ok {
		parentID_247_ptr := parentID_247
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "showroom", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_247,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_247_ptr,
			"locale":      "en",
			"short_code":  "showroom",
			"key":         "showroom",
			"label":       "Showroom",
			"description": "",
			"value":       "Showroom",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_248 := uuid.New().String()
	optionIDMap["estate_subtype:showroom:es"] = optID_248
	if parentID_248, ok := optionIDMap["estate_type:retail:es"]; ok {
		parentID_248_ptr := parentID_248
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "showroom", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_248,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_248_ptr,
			"locale":      "es",
			"short_code":  "showroom",
			"key":         "showroom",
			"label":       "Showroom",
			"description": "",
			"value":       "Showroom",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_249 := uuid.New().String()
	optionIDMap["estate_subtype:showroom:pl"] = optID_249
	if parentID_249, ok := optionIDMap["estate_type:retail:pl"]; ok {
		parentID_249_ptr := parentID_249
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "showroom", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_249,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_249_ptr,
			"locale":      "pl",
			"short_code":  "showroom",
			"key":         "showroom",
			"label":       "Salon ekspozycyjny",
			"description": "",
			"value":       "Showroom",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_250 := uuid.New().String()
	optionIDMap["estate_subtype:hotel:en"] = optID_250
	if parentID_250, ok := optionIDMap["estate_type:hospitality:en"]; ok {
		parentID_250_ptr := parentID_250
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "hotel", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_250,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_250_ptr,
			"locale":      "en",
			"short_code":  "hotel",
			"key":         "hotel",
			"label":       "Hotel",
			"description": "",
			"value":       "Hotel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_251 := uuid.New().String()
	optionIDMap["estate_subtype:hotel:es"] = optID_251
	if parentID_251, ok := optionIDMap["estate_type:hospitality:es"]; ok {
		parentID_251_ptr := parentID_251
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "hotel", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_251,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_251_ptr,
			"locale":      "es",
			"short_code":  "hotel",
			"key":         "hotel",
			"label":       "Hotel",
			"description": "",
			"value":       "Hotel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_252 := uuid.New().String()
	optionIDMap["estate_subtype:hotel:pl"] = optID_252
	if parentID_252, ok := optionIDMap["estate_type:hospitality:pl"]; ok {
		parentID_252_ptr := parentID_252
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "hotel", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_252,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_252_ptr,
			"locale":      "pl",
			"short_code":  "hotel",
			"key":         "hotel",
			"label":       "Hotel",
			"description": "",
			"value":       "Hotel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_253 := uuid.New().String()
	optionIDMap["estate_subtype:motel:en"] = optID_253
	if parentID_253, ok := optionIDMap["estate_type:hospitality:en"]; ok {
		parentID_253_ptr := parentID_253
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "motel", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_253,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_253_ptr,
			"locale":      "en",
			"short_code":  "motel",
			"key":         "motel",
			"label":       "Motel",
			"description": "",
			"value":       "Motel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_254 := uuid.New().String()
	optionIDMap["estate_subtype:motel:es"] = optID_254
	if parentID_254, ok := optionIDMap["estate_type:hospitality:es"]; ok {
		parentID_254_ptr := parentID_254
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "motel", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_254,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_254_ptr,
			"locale":      "es",
			"short_code":  "motel",
			"key":         "motel",
			"label":       "Motel",
			"description": "",
			"value":       "Motel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_255 := uuid.New().String()
	optionIDMap["estate_subtype:motel:pl"] = optID_255
	if parentID_255, ok := optionIDMap["estate_type:hospitality:pl"]; ok {
		parentID_255_ptr := parentID_255
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "motel", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_255,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_255_ptr,
			"locale":      "pl",
			"short_code":  "motel",
			"key":         "motel",
			"label":       "Motel",
			"description": "",
			"value":       "Motel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_256 := uuid.New().String()
	optionIDMap["estate_subtype:boutique_hotel:en"] = optID_256
	if parentID_256, ok := optionIDMap["estate_type:hospitality:en"]; ok {
		parentID_256_ptr := parentID_256
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "boutique_hotel", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_256,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_256_ptr,
			"locale":      "en",
			"short_code":  "boutique",
			"key":         "boutique_hotel",
			"label":       "Boutique hotel",
			"description": "",
			"value":       "Boutique hotel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_257 := uuid.New().String()
	optionIDMap["estate_subtype:boutique_hotel:es"] = optID_257
	if parentID_257, ok := optionIDMap["estate_type:hospitality:es"]; ok {
		parentID_257_ptr := parentID_257
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "boutique_hotel", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_257,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_257_ptr,
			"locale":      "es",
			"short_code":  "boutique",
			"key":         "boutique_hotel",
			"label":       "Hotel boutique",
			"description": "",
			"value":       "Boutique hotel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_258 := uuid.New().String()
	optionIDMap["estate_subtype:boutique_hotel:pl"] = optID_258
	if parentID_258, ok := optionIDMap["estate_type:hospitality:pl"]; ok {
		parentID_258_ptr := parentID_258
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "boutique_hotel", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_258,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_258_ptr,
			"locale":      "pl",
			"short_code":  "boutique",
			"key":         "boutique_hotel",
			"label":       "Hotel butikowy",
			"description": "",
			"value":       "Boutique hotel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_259 := uuid.New().String()
	optionIDMap["estate_subtype:bnb:en"] = optID_259
	if parentID_259, ok := optionIDMap["estate_type:hospitality:en"]; ok {
		parentID_259_ptr := parentID_259
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "bnb", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_259,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_259_ptr,
			"locale":      "en",
			"short_code":  "bnb",
			"key":         "bnb",
			"label":       "Bed and breakfast",
			"description": "",
			"value":       "Bed and breakfast",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_260 := uuid.New().String()
	optionIDMap["estate_subtype:bnb:es"] = optID_260
	if parentID_260, ok := optionIDMap["estate_type:hospitality:es"]; ok {
		parentID_260_ptr := parentID_260
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "bnb", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_260,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_260_ptr,
			"locale":      "es",
			"short_code":  "bnb",
			"key":         "bnb",
			"label":       "Bed and breakfast",
			"description": "",
			"value":       "Bed and breakfast",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_261 := uuid.New().String()
	optionIDMap["estate_subtype:bnb:pl"] = optID_261
	if parentID_261, ok := optionIDMap["estate_type:hospitality:pl"]; ok {
		parentID_261_ptr := parentID_261
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "bnb", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_261,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_261_ptr,
			"locale":      "pl",
			"short_code":  "bnb",
			"key":         "bnb",
			"label":       "Pensjonat B&B",
			"description": "",
			"value":       "Bed and breakfast",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_262 := uuid.New().String()
	optionIDMap["estate_subtype:hostel:en"] = optID_262
	if parentID_262, ok := optionIDMap["estate_type:hospitality:en"]; ok {
		parentID_262_ptr := parentID_262
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "hostel", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_262,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_262_ptr,
			"locale":      "en",
			"short_code":  "hostel",
			"key":         "hostel",
			"label":       "Hostel",
			"description": "",
			"value":       "Hostel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_263 := uuid.New().String()
	optionIDMap["estate_subtype:hostel:es"] = optID_263
	if parentID_263, ok := optionIDMap["estate_type:hospitality:es"]; ok {
		parentID_263_ptr := parentID_263
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "hostel", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_263,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_263_ptr,
			"locale":      "es",
			"short_code":  "hostel",
			"key":         "hostel",
			"label":       "Hostel",
			"description": "",
			"value":       "Hostel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_264 := uuid.New().String()
	optionIDMap["estate_subtype:hostel:pl"] = optID_264
	if parentID_264, ok := optionIDMap["estate_type:hospitality:pl"]; ok {
		parentID_264_ptr := parentID_264
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "hostel", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_264,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_264_ptr,
			"locale":      "pl",
			"short_code":  "hostel",
			"key":         "hostel",
			"label":       "Hostel",
			"description": "",
			"value":       "Hostel",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_265 := uuid.New().String()
	optionIDMap["estate_subtype:restaurant:en"] = optID_265
	if parentID_265, ok := optionIDMap["estate_type:food_beverage:en"]; ok {
		parentID_265_ptr := parentID_265
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "restaurant", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_265,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_265_ptr,
			"locale":      "en",
			"short_code":  "restaura",
			"key":         "restaurant",
			"label":       "Restaurant",
			"description": "",
			"value":       "Restaurant",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_266 := uuid.New().String()
	optionIDMap["estate_subtype:restaurant:es"] = optID_266
	if parentID_266, ok := optionIDMap["estate_type:food_beverage:es"]; ok {
		parentID_266_ptr := parentID_266
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "restaurant", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_266,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_266_ptr,
			"locale":      "es",
			"short_code":  "restaura",
			"key":         "restaurant",
			"label":       "Restaurante",
			"description": "",
			"value":       "Restaurant",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_267 := uuid.New().String()
	optionIDMap["estate_subtype:restaurant:pl"] = optID_267
	if parentID_267, ok := optionIDMap["estate_type:food_beverage:pl"]; ok {
		parentID_267_ptr := parentID_267
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "restaurant", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_267,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_267_ptr,
			"locale":      "pl",
			"short_code":  "restaura",
			"key":         "restaurant",
			"label":       "Restauracja",
			"description": "",
			"value":       "Restaurant",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_268 := uuid.New().String()
	optionIDMap["estate_subtype:cafe:en"] = optID_268
	if parentID_268, ok := optionIDMap["estate_type:food_beverage:en"]; ok {
		parentID_268_ptr := parentID_268
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "cafe", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_268,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_268_ptr,
			"locale":      "en",
			"short_code":  "cafe",
			"key":         "cafe",
			"label":       "Café",
			"description": "",
			"value":       "Café",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_269 := uuid.New().String()
	optionIDMap["estate_subtype:cafe:es"] = optID_269
	if parentID_269, ok := optionIDMap["estate_type:food_beverage:es"]; ok {
		parentID_269_ptr := parentID_269
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "cafe", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_269,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_269_ptr,
			"locale":      "es",
			"short_code":  "cafe",
			"key":         "cafe",
			"label":       "Cafetería",
			"description": "",
			"value":       "Café",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_270 := uuid.New().String()
	optionIDMap["estate_subtype:cafe:pl"] = optID_270
	if parentID_270, ok := optionIDMap["estate_type:food_beverage:pl"]; ok {
		parentID_270_ptr := parentID_270
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "cafe", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_270,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_270_ptr,
			"locale":      "pl",
			"short_code":  "cafe",
			"key":         "cafe",
			"label":       "Kawiarnia",
			"description": "",
			"value":       "Café",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_271 := uuid.New().String()
	optionIDMap["estate_subtype:bar_pub:en"] = optID_271
	if parentID_271, ok := optionIDMap["estate_type:food_beverage:en"]; ok {
		parentID_271_ptr := parentID_271
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "bar_pub", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_271,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_271_ptr,
			"locale":      "en",
			"short_code":  "bar_pub",
			"key":         "bar_pub",
			"label":       "Bar / Pub",
			"description": "",
			"value":       "Bar / Pub",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_272 := uuid.New().String()
	optionIDMap["estate_subtype:bar_pub:es"] = optID_272
	if parentID_272, ok := optionIDMap["estate_type:food_beverage:es"]; ok {
		parentID_272_ptr := parentID_272
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "bar_pub", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_272,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_272_ptr,
			"locale":      "es",
			"short_code":  "bar_pub",
			"key":         "bar_pub",
			"label":       "Bar / Pub",
			"description": "",
			"value":       "Bar / Pub",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_273 := uuid.New().String()
	optionIDMap["estate_subtype:bar_pub:pl"] = optID_273
	if parentID_273, ok := optionIDMap["estate_type:food_beverage:pl"]; ok {
		parentID_273_ptr := parentID_273
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "bar_pub", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_273,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_273_ptr,
			"locale":      "pl",
			"short_code":  "bar_pub",
			"key":         "bar_pub",
			"label":       "Bar / Pub",
			"description": "",
			"value":       "Bar / Pub",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_274 := uuid.New().String()
	optionIDMap["estate_subtype:nightclub:en"] = optID_274
	if parentID_274, ok := optionIDMap["estate_type:food_beverage:en"]; ok {
		parentID_274_ptr := parentID_274
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "nightclub", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_274,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_274_ptr,
			"locale":      "en",
			"short_code":  "nightclu",
			"key":         "nightclub",
			"label":       "Nightclub",
			"description": "",
			"value":       "Nightclub",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_275 := uuid.New().String()
	optionIDMap["estate_subtype:nightclub:es"] = optID_275
	if parentID_275, ok := optionIDMap["estate_type:food_beverage:es"]; ok {
		parentID_275_ptr := parentID_275
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "nightclub", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_275,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_275_ptr,
			"locale":      "es",
			"short_code":  "nightclu",
			"key":         "nightclub",
			"label":       "Discoteca",
			"description": "",
			"value":       "Nightclub",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_276 := uuid.New().String()
	optionIDMap["estate_subtype:nightclub:pl"] = optID_276
	if parentID_276, ok := optionIDMap["estate_type:food_beverage:pl"]; ok {
		parentID_276_ptr := parentID_276
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "nightclub", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_276,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_276_ptr,
			"locale":      "pl",
			"short_code":  "nightclu",
			"key":         "nightclub",
			"label":       "Klub nocny",
			"description": "",
			"value":       "Nightclub",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_277 := uuid.New().String()
	optionIDMap["estate_subtype:medical_office:en"] = optID_277
	if parentID_277, ok := optionIDMap["estate_type:medical:en"]; ok {
		parentID_277_ptr := parentID_277
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "medical_office", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_277,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_277_ptr,
			"locale":      "en",
			"short_code":  "medical_",
			"key":         "medical_office",
			"label":       "Medical office",
			"description": "",
			"value":       "Medical office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_278 := uuid.New().String()
	optionIDMap["estate_subtype:medical_office:es"] = optID_278
	if parentID_278, ok := optionIDMap["estate_type:medical:es"]; ok {
		parentID_278_ptr := parentID_278
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "medical_office", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_278,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_278_ptr,
			"locale":      "es",
			"short_code":  "medical_",
			"key":         "medical_office",
			"label":       "Consultorio médico",
			"description": "",
			"value":       "Medical office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_279 := uuid.New().String()
	optionIDMap["estate_subtype:medical_office:pl"] = optID_279
	if parentID_279, ok := optionIDMap["estate_type:medical:pl"]; ok {
		parentID_279_ptr := parentID_279
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "medical_office", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_279,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_279_ptr,
			"locale":      "pl",
			"short_code":  "medical_",
			"key":         "medical_office",
			"label":       "Przychodnia",
			"description": "",
			"value":       "Medical office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_280 := uuid.New().String()
	optionIDMap["estate_subtype:dental_clinic:en"] = optID_280
	if parentID_280, ok := optionIDMap["estate_type:medical:en"]; ok {
		parentID_280_ptr := parentID_280
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "dental_clinic", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_280,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_280_ptr,
			"locale":      "en",
			"short_code":  "dental_c",
			"key":         "dental_clinic",
			"label":       "Dental clinic",
			"description": "",
			"value":       "Dental clinic",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_281 := uuid.New().String()
	optionIDMap["estate_subtype:dental_clinic:es"] = optID_281
	if parentID_281, ok := optionIDMap["estate_type:medical:es"]; ok {
		parentID_281_ptr := parentID_281
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "dental_clinic", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_281,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_281_ptr,
			"locale":      "es",
			"short_code":  "dental_c",
			"key":         "dental_clinic",
			"label":       "Clínica dental",
			"description": "",
			"value":       "Dental clinic",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_282 := uuid.New().String()
	optionIDMap["estate_subtype:dental_clinic:pl"] = optID_282
	if parentID_282, ok := optionIDMap["estate_type:medical:pl"]; ok {
		parentID_282_ptr := parentID_282
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "dental_clinic", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_282,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_282_ptr,
			"locale":      "pl",
			"short_code":  "dental_c",
			"key":         "dental_clinic",
			"label":       "Klinika stomatologiczna",
			"description": "",
			"value":       "Dental clinic",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_283 := uuid.New().String()
	optionIDMap["estate_subtype:vet_clinic:en"] = optID_283
	if parentID_283, ok := optionIDMap["estate_type:medical:en"]; ok {
		parentID_283_ptr := parentID_283
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "vet_clinic", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_283,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_283_ptr,
			"locale":      "en",
			"short_code":  "vet_clin",
			"key":         "vet_clinic",
			"label":       "Veterinary clinic",
			"description": "",
			"value":       "Veterinary clinic",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_284 := uuid.New().String()
	optionIDMap["estate_subtype:vet_clinic:es"] = optID_284
	if parentID_284, ok := optionIDMap["estate_type:medical:es"]; ok {
		parentID_284_ptr := parentID_284
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "vet_clinic", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_284,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_284_ptr,
			"locale":      "es",
			"short_code":  "vet_clin",
			"key":         "vet_clinic",
			"label":       "Clínica veterinaria",
			"description": "",
			"value":       "Veterinary clinic",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_285 := uuid.New().String()
	optionIDMap["estate_subtype:vet_clinic:pl"] = optID_285
	if parentID_285, ok := optionIDMap["estate_type:medical:pl"]; ok {
		parentID_285_ptr := parentID_285
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "vet_clinic", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_285,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_285_ptr,
			"locale":      "pl",
			"short_code":  "vet_clin",
			"key":         "vet_clinic",
			"label":       "Przychodnia weterynaryjna",
			"description": "",
			"value":       "Veterinary clinic",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_286 := uuid.New().String()
	optionIDMap["estate_subtype:wellness_center:en"] = optID_286
	if parentID_286, ok := optionIDMap["estate_type:medical:en"]; ok {
		parentID_286_ptr := parentID_286
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "wellness_center", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_286,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_286_ptr,
			"locale":      "en",
			"short_code":  "wellness",
			"key":         "wellness_center",
			"label":       "Wellness center / Spa",
			"description": "",
			"value":       "Wellness center / Spa",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_287 := uuid.New().String()
	optionIDMap["estate_subtype:wellness_center:es"] = optID_287
	if parentID_287, ok := optionIDMap["estate_type:medical:es"]; ok {
		parentID_287_ptr := parentID_287
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "wellness_center", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_287,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_287_ptr,
			"locale":      "es",
			"short_code":  "wellness",
			"key":         "wellness_center",
			"label":       "Centro de bienestar / Spa",
			"description": "",
			"value":       "Wellness center / Spa",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_288 := uuid.New().String()
	optionIDMap["estate_subtype:wellness_center:pl"] = optID_288
	if parentID_288, ok := optionIDMap["estate_type:medical:pl"]; ok {
		parentID_288_ptr := parentID_288
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "wellness_center", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_288,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_288_ptr,
			"locale":      "pl",
			"short_code":  "wellness",
			"key":         "wellness_center",
			"label":       "Spa / Centrum wellness",
			"description": "",
			"value":       "Wellness center / Spa",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_289 := uuid.New().String()
	optionIDMap["estate_subtype:warehouse:en"] = optID_289
	if parentID_289, ok := optionIDMap["estate_type:industrial:en"]; ok {
		parentID_289_ptr := parentID_289
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "warehouse", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_289,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_289_ptr,
			"locale":      "en",
			"short_code":  "warehous",
			"key":         "warehouse",
			"label":       "Warehouse",
			"description": "",
			"value":       "Warehouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_290 := uuid.New().String()
	optionIDMap["estate_subtype:warehouse:es"] = optID_290
	if parentID_290, ok := optionIDMap["estate_type:industrial:es"]; ok {
		parentID_290_ptr := parentID_290
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "warehouse", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_290,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_290_ptr,
			"locale":      "es",
			"short_code":  "warehous",
			"key":         "warehouse",
			"label":       "Depósito / Almacén",
			"description": "",
			"value":       "Warehouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_291 := uuid.New().String()
	optionIDMap["estate_subtype:warehouse:pl"] = optID_291
	if parentID_291, ok := optionIDMap["estate_type:industrial:pl"]; ok {
		parentID_291_ptr := parentID_291
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "warehouse", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_291,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_291_ptr,
			"locale":      "pl",
			"short_code":  "warehous",
			"key":         "warehouse",
			"label":       "Magazyn",
			"description": "",
			"value":       "Warehouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_292 := uuid.New().String()
	optionIDMap["estate_subtype:factory:en"] = optID_292
	if parentID_292, ok := optionIDMap["estate_type:industrial:en"]; ok {
		parentID_292_ptr := parentID_292
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "factory", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_292,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_292_ptr,
			"locale":      "en",
			"short_code":  "factory",
			"key":         "factory",
			"label":       "Factory",
			"description": "",
			"value":       "Factory",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_293 := uuid.New().String()
	optionIDMap["estate_subtype:factory:es"] = optID_293
	if parentID_293, ok := optionIDMap["estate_type:industrial:es"]; ok {
		parentID_293_ptr := parentID_293
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "factory", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_293,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_293_ptr,
			"locale":      "es",
			"short_code":  "factory",
			"key":         "factory",
			"label":       "Fábrica",
			"description": "",
			"value":       "Factory",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_294 := uuid.New().String()
	optionIDMap["estate_subtype:factory:pl"] = optID_294
	if parentID_294, ok := optionIDMap["estate_type:industrial:pl"]; ok {
		parentID_294_ptr := parentID_294
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "factory", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_294,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_294_ptr,
			"locale":      "pl",
			"short_code":  "factory",
			"key":         "factory",
			"label":       "Fabryka",
			"description": "",
			"value":       "Factory",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_295 := uuid.New().String()
	optionIDMap["estate_subtype:cold_storage:en"] = optID_295
	if parentID_295, ok := optionIDMap["estate_type:industrial:en"]; ok {
		parentID_295_ptr := parentID_295
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "cold_storage", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_295,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_295_ptr,
			"locale":      "en",
			"short_code":  "cold_sto",
			"key":         "cold_storage",
			"label":       "Cold storage",
			"description": "",
			"value":       "Cold storage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_296 := uuid.New().String()
	optionIDMap["estate_subtype:cold_storage:es"] = optID_296
	if parentID_296, ok := optionIDMap["estate_type:industrial:es"]; ok {
		parentID_296_ptr := parentID_296
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "cold_storage", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_296,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_296_ptr,
			"locale":      "es",
			"short_code":  "cold_sto",
			"key":         "cold_storage",
			"label":       "Cámara frigorífica",
			"description": "",
			"value":       "Cold storage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_297 := uuid.New().String()
	optionIDMap["estate_subtype:cold_storage:pl"] = optID_297
	if parentID_297, ok := optionIDMap["estate_type:industrial:pl"]; ok {
		parentID_297_ptr := parentID_297
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "cold_storage", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_297,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_297_ptr,
			"locale":      "pl",
			"short_code":  "cold_sto",
			"key":         "cold_storage",
			"label":       "Chłodnia",
			"description": "",
			"value":       "Cold storage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_298 := uuid.New().String()
	optionIDMap["estate_subtype:distribution_center:en"] = optID_298
	if parentID_298, ok := optionIDMap["estate_type:industrial:en"]; ok {
		parentID_298_ptr := parentID_298
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "distribution_center", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_298,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_298_ptr,
			"locale":      "en",
			"short_code":  "distribu",
			"key":         "distribution_center",
			"label":       "Distribution center",
			"description": "",
			"value":       "Distribution center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_299 := uuid.New().String()
	optionIDMap["estate_subtype:distribution_center:es"] = optID_299
	if parentID_299, ok := optionIDMap["estate_type:industrial:es"]; ok {
		parentID_299_ptr := parentID_299
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "distribution_center", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_299,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_299_ptr,
			"locale":      "es",
			"short_code":  "distribu",
			"key":         "distribution_center",
			"label":       "Centro de distribución",
			"description": "",
			"value":       "Distribution center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_300 := uuid.New().String()
	optionIDMap["estate_subtype:distribution_center:pl"] = optID_300
	if parentID_300, ok := optionIDMap["estate_type:industrial:pl"]; ok {
		parentID_300_ptr := parentID_300
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "distribution_center", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_300,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_300_ptr,
			"locale":      "pl",
			"short_code":  "distribu",
			"key":         "distribution_center",
			"label":       "Centrum dystrybucyjne",
			"description": "",
			"value":       "Distribution center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_301 := uuid.New().String()
	optionIDMap["estate_subtype:workshop:en"] = optID_301
	if parentID_301, ok := optionIDMap["estate_type:industrial:en"]; ok {
		parentID_301_ptr := parentID_301
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "workshop", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_301,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_301_ptr,
			"locale":      "en",
			"short_code":  "workshop",
			"key":         "workshop",
			"label":       "Workshop",
			"description": "",
			"value":       "Workshop",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_302 := uuid.New().String()
	optionIDMap["estate_subtype:workshop:es"] = optID_302
	if parentID_302, ok := optionIDMap["estate_type:industrial:es"]; ok {
		parentID_302_ptr := parentID_302
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "workshop", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_302,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_302_ptr,
			"locale":      "es",
			"short_code":  "workshop",
			"key":         "workshop",
			"label":       "Taller",
			"description": "",
			"value":       "Workshop",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_303 := uuid.New().String()
	optionIDMap["estate_subtype:workshop:pl"] = optID_303
	if parentID_303, ok := optionIDMap["estate_type:industrial:pl"]; ok {
		parentID_303_ptr := parentID_303
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "workshop", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_303,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_303_ptr,
			"locale":      "pl",
			"short_code":  "workshop",
			"key":         "workshop",
			"label":       "Warsztat",
			"description": "",
			"value":       "Workshop",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_304 := uuid.New().String()
	optionIDMap["estate_subtype:data_center:en"] = optID_304
	if parentID_304, ok := optionIDMap["estate_type:industrial:en"]; ok {
		parentID_304_ptr := parentID_304
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "data_center", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_304,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_304_ptr,
			"locale":      "en",
			"short_code":  "data_cen",
			"key":         "data_center",
			"label":       "Data center",
			"description": "",
			"value":       "Data center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_305 := uuid.New().String()
	optionIDMap["estate_subtype:data_center:es"] = optID_305
	if parentID_305, ok := optionIDMap["estate_type:industrial:es"]; ok {
		parentID_305_ptr := parentID_305
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "data_center", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_305,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_305_ptr,
			"locale":      "es",
			"short_code":  "data_cen",
			"key":         "data_center",
			"label":       "Data center",
			"description": "",
			"value":       "Data center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_306 := uuid.New().String()
	optionIDMap["estate_subtype:data_center:pl"] = optID_306
	if parentID_306, ok := optionIDMap["estate_type:industrial:pl"]; ok {
		parentID_306_ptr := parentID_306
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "data_center", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_306,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_306_ptr,
			"locale":      "pl",
			"short_code":  "data_cen",
			"key":         "data_center",
			"label":       "Centrum danych",
			"description": "",
			"value":       "Data center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_307 := uuid.New().String()
	optionIDMap["estate_subtype:gym:en"] = optID_307
	if parentID_307, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_307_ptr := parentID_307
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "gym", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_307,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_307_ptr,
			"locale":      "en",
			"short_code":  "gym",
			"key":         "gym",
			"label":       "Gym / Fitness",
			"description": "",
			"value":       "Gym / Fitness",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_308 := uuid.New().String()
	optionIDMap["estate_subtype:gym:es"] = optID_308
	if parentID_308, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_308_ptr := parentID_308
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "gym", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_308,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_308_ptr,
			"locale":      "es",
			"short_code":  "gym",
			"key":         "gym",
			"label":       "Gimnasio",
			"description": "",
			"value":       "Gym / Fitness",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_309 := uuid.New().String()
	optionIDMap["estate_subtype:gym:pl"] = optID_309
	if parentID_309, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_309_ptr := parentID_309
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "gym", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_309,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_309_ptr,
			"locale":      "pl",
			"short_code":  "gym",
			"key":         "gym",
			"label":       "Siłownia",
			"description": "",
			"value":       "Gym / Fitness",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_310 := uuid.New().String()
	optionIDMap["estate_subtype:studio:en"] = optID_310
	if parentID_310, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_310_ptr := parentID_310
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "studio", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_310,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_310_ptr,
			"locale":      "en",
			"short_code":  "studio",
			"key":         "studio",
			"label":       "Studio (yoga/dance)",
			"description": "",
			"value":       "Studio (yoga/dance)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_311 := uuid.New().String()
	optionIDMap["estate_subtype:studio:es"] = optID_311
	if parentID_311, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_311_ptr := parentID_311
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "studio", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_311,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_311_ptr,
			"locale":      "es",
			"short_code":  "studio",
			"key":         "studio",
			"label":       "Estudio (yoga/danza)",
			"description": "",
			"value":       "Studio (yoga/dance)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_312 := uuid.New().String()
	optionIDMap["estate_subtype:studio:pl"] = optID_312
	if parentID_312, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_312_ptr := parentID_312
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "studio", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_312,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_312_ptr,
			"locale":      "pl",
			"short_code":  "studio",
			"key":         "studio",
			"label":       "Studio (joga/taniec)",
			"description": "",
			"value":       "Studio (yoga/dance)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_313 := uuid.New().String()
	optionIDMap["estate_subtype:bank_branch:en"] = optID_313
	if parentID_313, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_313_ptr := parentID_313
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "bank_branch", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_313,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_313_ptr,
			"locale":      "en",
			"short_code":  "bank_bra",
			"key":         "bank_branch",
			"label":       "Bank branch",
			"description": "",
			"value":       "Bank branch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_314 := uuid.New().String()
	optionIDMap["estate_subtype:bank_branch:es"] = optID_314
	if parentID_314, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_314_ptr := parentID_314
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "bank_branch", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_314,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_314_ptr,
			"locale":      "es",
			"short_code":  "bank_bra",
			"key":         "bank_branch",
			"label":       "Sucursal bancaria",
			"description": "",
			"value":       "Bank branch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_315 := uuid.New().String()
	optionIDMap["estate_subtype:bank_branch:pl"] = optID_315
	if parentID_315, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_315_ptr := parentID_315
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "bank_branch", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_315,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_315_ptr,
			"locale":      "pl",
			"short_code":  "bank_bra",
			"key":         "bank_branch",
			"label":       "Oddział banku",
			"description": "",
			"value":       "Bank branch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_316 := uuid.New().String()
	optionIDMap["estate_subtype:car_dealership:en"] = optID_316
	if parentID_316, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_316_ptr := parentID_316
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "car_dealership", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_316,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_316_ptr,
			"locale":      "en",
			"short_code":  "car_deal",
			"key":         "car_dealership",
			"label":       "Car dealership",
			"description": "",
			"value":       "Car dealership",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_317 := uuid.New().String()
	optionIDMap["estate_subtype:car_dealership:es"] = optID_317
	if parentID_317, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_317_ptr := parentID_317
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "car_dealership", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_317,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_317_ptr,
			"locale":      "es",
			"short_code":  "car_deal",
			"key":         "car_dealership",
			"label":       "Concesionario",
			"description": "",
			"value":       "Car dealership",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_318 := uuid.New().String()
	optionIDMap["estate_subtype:car_dealership:pl"] = optID_318
	if parentID_318, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_318_ptr := parentID_318
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "car_dealership", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_318,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_318_ptr,
			"locale":      "pl",
			"short_code":  "car_deal",
			"key":         "car_dealership",
			"label":       "Salon samochodowy",
			"description": "",
			"value":       "Car dealership",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_319 := uuid.New().String()
	optionIDMap["estate_subtype:funeral_home:en"] = optID_319
	if parentID_319, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_319_ptr := parentID_319
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "funeral_home", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_319,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_319_ptr,
			"locale":      "en",
			"short_code":  "funeral_",
			"key":         "funeral_home",
			"label":       "Funeral home",
			"description": "",
			"value":       "Funeral home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_320 := uuid.New().String()
	optionIDMap["estate_subtype:funeral_home:es"] = optID_320
	if parentID_320, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_320_ptr := parentID_320
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "funeral_home", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_320,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_320_ptr,
			"locale":      "es",
			"short_code":  "funeral_",
			"key":         "funeral_home",
			"label":       "Casa funeraria",
			"description": "",
			"value":       "Funeral home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_321 := uuid.New().String()
	optionIDMap["estate_subtype:funeral_home:pl"] = optID_321
	if parentID_321, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_321_ptr := parentID_321
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "funeral_home", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_321,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_321_ptr,
			"locale":      "pl",
			"short_code":  "funeral_",
			"key":         "funeral_home",
			"label":       "Dom pogrzebowy",
			"description": "",
			"value":       "Funeral home",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_322 := uuid.New().String()
	optionIDMap["estate_subtype:religious_facility:en"] = optID_322
	if parentID_322, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_322_ptr := parentID_322
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "religious_facility", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_322,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_322_ptr,
			"locale":      "en",
			"short_code":  "religiou",
			"key":         "religious_facility",
			"label":       "Religious facility",
			"description": "",
			"value":       "Religious facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_323 := uuid.New().String()
	optionIDMap["estate_subtype:religious_facility:es"] = optID_323
	if parentID_323, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_323_ptr := parentID_323
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "religious_facility", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_323,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_323_ptr,
			"locale":      "es",
			"short_code":  "religiou",
			"key":         "religious_facility",
			"label":       "Templo / Iglesia / Mezquita",
			"description": "",
			"value":       "Religious facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_324 := uuid.New().String()
	optionIDMap["estate_subtype:religious_facility:pl"] = optID_324
	if parentID_324, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_324_ptr := parentID_324
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "religious_facility", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_324,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_324_ptr,
			"locale":      "pl",
			"short_code":  "religiou",
			"key":         "religious_facility",
			"label":       "Obiekt religijny",
			"description": "",
			"value":       "Religious facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_325 := uuid.New().String()
	optionIDMap["estate_subtype:school:en"] = optID_325
	if parentID_325, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_325_ptr := parentID_325
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "school", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_325,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_325_ptr,
			"locale":      "en",
			"short_code":  "school",
			"key":         "school",
			"label":       "School building",
			"description": "",
			"value":       "School building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_326 := uuid.New().String()
	optionIDMap["estate_subtype:school:es"] = optID_326
	if parentID_326, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_326_ptr := parentID_326
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "school", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_326,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_326_ptr,
			"locale":      "es",
			"short_code":  "school",
			"key":         "school",
			"label":       "Edificio escolar",
			"description": "",
			"value":       "School building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_327 := uuid.New().String()
	optionIDMap["estate_subtype:school:pl"] = optID_327
	if parentID_327, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_327_ptr := parentID_327
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "school", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_327,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_327_ptr,
			"locale":      "pl",
			"short_code":  "school",
			"key":         "school",
			"label":       "Budynek szkolny",
			"description": "",
			"value":       "School building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_328 := uuid.New().String()
	optionIDMap["estate_subtype:government:en"] = optID_328
	if parentID_328, ok := optionIDMap["estate_type:special_com:en"]; ok {
		parentID_328_ptr := parentID_328
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "government", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_328,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_328_ptr,
			"locale":      "en",
			"short_code":  "governme",
			"key":         "government",
			"label":       "Government building",
			"description": "",
			"value":       "Government building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_329 := uuid.New().String()
	optionIDMap["estate_subtype:government:es"] = optID_329
	if parentID_329, ok := optionIDMap["estate_type:special_com:es"]; ok {
		parentID_329_ptr := parentID_329
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "government", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_329,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_329_ptr,
			"locale":      "es",
			"short_code":  "governme",
			"key":         "government",
			"label":       "Edificio gubernamental",
			"description": "",
			"value":       "Government building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_330 := uuid.New().String()
	optionIDMap["estate_subtype:government:pl"] = optID_330
	if parentID_330, ok := optionIDMap["estate_type:special_com:pl"]; ok {
		parentID_330_ptr := parentID_330
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "government", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_330,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_330_ptr,
			"locale":      "pl",
			"short_code":  "governme",
			"key":         "government",
			"label":       "Budynek rządowy",
			"description": "",
			"value":       "Government building",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_331 := uuid.New().String()
	optionIDMap["estate_subtype:residential_lot:en"] = optID_331
	if parentID_331, ok := optionIDMap["estate_type:urban_land:en"]; ok {
		parentID_331_ptr := parentID_331
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "residential_lot", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_331,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_331_ptr,
			"locale":      "en",
			"short_code":  "resident",
			"key":         "residential_lot",
			"label":       "Residential lot",
			"description": "",
			"value":       "Residential lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_332 := uuid.New().String()
	optionIDMap["estate_subtype:residential_lot:es"] = optID_332
	if parentID_332, ok := optionIDMap["estate_type:urban_land:es"]; ok {
		parentID_332_ptr := parentID_332
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "residential_lot", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_332,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_332_ptr,
			"locale":      "es",
			"short_code":  "resident",
			"key":         "residential_lot",
			"label":       "Lote residencial",
			"description": "",
			"value":       "Residential lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_333 := uuid.New().String()
	optionIDMap["estate_subtype:residential_lot:pl"] = optID_333
	if parentID_333, ok := optionIDMap["estate_type:urban_land:pl"]; ok {
		parentID_333_ptr := parentID_333
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "residential_lot", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_333,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_333_ptr,
			"locale":      "pl",
			"short_code":  "resident",
			"key":         "residential_lot",
			"label":       "Działka mieszkaniowa",
			"description": "",
			"value":       "Residential lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_334 := uuid.New().String()
	optionIDMap["estate_subtype:commercial_lot:en"] = optID_334
	if parentID_334, ok := optionIDMap["estate_type:urban_land:en"]; ok {
		parentID_334_ptr := parentID_334
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "commercial_lot", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_334,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_334_ptr,
			"locale":      "en",
			"short_code":  "commerci",
			"key":         "commercial_lot",
			"label":       "Commercial lot",
			"description": "",
			"value":       "Commercial lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_335 := uuid.New().String()
	optionIDMap["estate_subtype:commercial_lot:es"] = optID_335
	if parentID_335, ok := optionIDMap["estate_type:urban_land:es"]; ok {
		parentID_335_ptr := parentID_335
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "commercial_lot", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_335,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_335_ptr,
			"locale":      "es",
			"short_code":  "commerci",
			"key":         "commercial_lot",
			"label":       "Lote comercial",
			"description": "",
			"value":       "Commercial lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_336 := uuid.New().String()
	optionIDMap["estate_subtype:commercial_lot:pl"] = optID_336
	if parentID_336, ok := optionIDMap["estate_type:urban_land:pl"]; ok {
		parentID_336_ptr := parentID_336
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "commercial_lot", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_336,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_336_ptr,
			"locale":      "pl",
			"short_code":  "commerci",
			"key":         "commercial_lot",
			"label":       "Działka komercyjna",
			"description": "",
			"value":       "Commercial lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_337 := uuid.New().String()
	optionIDMap["estate_subtype:corner_lot:en"] = optID_337
	if parentID_337, ok := optionIDMap["estate_type:urban_land:en"]; ok {
		parentID_337_ptr := parentID_337
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "corner_lot", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_337,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_337_ptr,
			"locale":      "en",
			"short_code":  "corner_l",
			"key":         "corner_lot",
			"label":       "Corner lot",
			"description": "",
			"value":       "Corner lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_338 := uuid.New().String()
	optionIDMap["estate_subtype:corner_lot:es"] = optID_338
	if parentID_338, ok := optionIDMap["estate_type:urban_land:es"]; ok {
		parentID_338_ptr := parentID_338
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "corner_lot", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_338,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_338_ptr,
			"locale":      "es",
			"short_code":  "corner_l",
			"key":         "corner_lot",
			"label":       "Lote en esquina",
			"description": "",
			"value":       "Corner lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_339 := uuid.New().String()
	optionIDMap["estate_subtype:corner_lot:pl"] = optID_339
	if parentID_339, ok := optionIDMap["estate_type:urban_land:pl"]; ok {
		parentID_339_ptr := parentID_339
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "corner_lot", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_339,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_339_ptr,
			"locale":      "pl",
			"short_code":  "corner_l",
			"key":         "corner_lot",
			"label":       "Działka narożna",
			"description": "",
			"value":       "Corner lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_340 := uuid.New().String()
	optionIDMap["estate_subtype:infill_lot:en"] = optID_340
	if parentID_340, ok := optionIDMap["estate_type:urban_land:en"]; ok {
		parentID_340_ptr := parentID_340
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "infill_lot", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_340,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_340_ptr,
			"locale":      "en",
			"short_code":  "infill_l",
			"key":         "infill_lot",
			"label":       "Infill lot",
			"description": "",
			"value":       "Infill lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_341 := uuid.New().String()
	optionIDMap["estate_subtype:infill_lot:es"] = optID_341
	if parentID_341, ok := optionIDMap["estate_type:urban_land:es"]; ok {
		parentID_341_ptr := parentID_341
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "infill_lot", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_341,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_341_ptr,
			"locale":      "es",
			"short_code":  "infill_l",
			"key":         "infill_lot",
			"label":       "Lote intersticial",
			"description": "",
			"value":       "Infill lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_342 := uuid.New().String()
	optionIDMap["estate_subtype:infill_lot:pl"] = optID_342
	if parentID_342, ok := optionIDMap["estate_type:urban_land:pl"]; ok {
		parentID_342_ptr := parentID_342
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "infill_lot", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_342,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_342_ptr,
			"locale":      "pl",
			"short_code":  "infill_l",
			"key":         "infill_lot",
			"label":       "Działka uzupełniająca",
			"description": "",
			"value":       "Infill lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_343 := uuid.New().String()
	optionIDMap["estate_subtype:agri_land:en"] = optID_343
	if parentID_343, ok := optionIDMap["estate_type:rural_land:en"]; ok {
		parentID_343_ptr := parentID_343
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "agri_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_343,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_343_ptr,
			"locale":      "en",
			"short_code":  "agri_lan",
			"key":         "agri_land",
			"label":       "Agricultural land",
			"description": "",
			"value":       "Agricultural land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_344 := uuid.New().String()
	optionIDMap["estate_subtype:agri_land:es"] = optID_344
	if parentID_344, ok := optionIDMap["estate_type:rural_land:es"]; ok {
		parentID_344_ptr := parentID_344
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "agri_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_344,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_344_ptr,
			"locale":      "es",
			"short_code":  "agri_lan",
			"key":         "agri_land",
			"label":       "Tierra agrícola",
			"description": "",
			"value":       "Agricultural land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_345 := uuid.New().String()
	optionIDMap["estate_subtype:agri_land:pl"] = optID_345
	if parentID_345, ok := optionIDMap["estate_type:rural_land:pl"]; ok {
		parentID_345_ptr := parentID_345
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "agri_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_345,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_345_ptr,
			"locale":      "pl",
			"short_code":  "agri_lan",
			"key":         "agri_land",
			"label":       "Grunty rolne",
			"description": "",
			"value":       "Agricultural land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_346 := uuid.New().String()
	optionIDMap["estate_subtype:timberland:en"] = optID_346
	if parentID_346, ok := optionIDMap["estate_type:rural_land:en"]; ok {
		parentID_346_ptr := parentID_346
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "timberland", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_346,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_346_ptr,
			"locale":      "en",
			"short_code":  "timberla",
			"key":         "timberland",
			"label":       "Timberland / Forest land",
			"description": "",
			"value":       "Timberland / Forest land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_347 := uuid.New().String()
	optionIDMap["estate_subtype:timberland:es"] = optID_347
	if parentID_347, ok := optionIDMap["estate_type:rural_land:es"]; ok {
		parentID_347_ptr := parentID_347
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "timberland", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_347,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_347_ptr,
			"locale":      "es",
			"short_code":  "timberla",
			"key":         "timberland",
			"label":       "Bosque / Forestal",
			"description": "",
			"value":       "Timberland / Forest land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_348 := uuid.New().String()
	optionIDMap["estate_subtype:timberland:pl"] = optID_348
	if parentID_348, ok := optionIDMap["estate_type:rural_land:pl"]; ok {
		parentID_348_ptr := parentID_348
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "timberland", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_348,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_348_ptr,
			"locale":      "pl",
			"short_code":  "timberla",
			"key":         "timberland",
			"label":       "Teren leśny",
			"description": "",
			"value":       "Timberland / Forest land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_349 := uuid.New().String()
	optionIDMap["estate_subtype:grazing_land:en"] = optID_349
	if parentID_349, ok := optionIDMap["estate_type:rural_land:en"]; ok {
		parentID_349_ptr := parentID_349
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "grazing_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_349,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_349_ptr,
			"locale":      "en",
			"short_code":  "grazing_",
			"key":         "grazing_land",
			"label":       "Grazing land",
			"description": "",
			"value":       "Grazing land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_350 := uuid.New().String()
	optionIDMap["estate_subtype:grazing_land:es"] = optID_350
	if parentID_350, ok := optionIDMap["estate_type:rural_land:es"]; ok {
		parentID_350_ptr := parentID_350
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "grazing_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_350,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_350_ptr,
			"locale":      "es",
			"short_code":  "grazing_",
			"key":         "grazing_land",
			"label":       "Pasto / Ganadero",
			"description": "",
			"value":       "Grazing land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_351 := uuid.New().String()
	optionIDMap["estate_subtype:grazing_land:pl"] = optID_351
	if parentID_351, ok := optionIDMap["estate_type:rural_land:pl"]; ok {
		parentID_351_ptr := parentID_351
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "grazing_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_351,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_351_ptr,
			"locale":      "pl",
			"short_code":  "grazing_",
			"key":         "grazing_land",
			"label":       "Pastwiska",
			"description": "",
			"value":       "Grazing land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_352 := uuid.New().String()
	optionIDMap["estate_subtype:undeveloped_land:en"] = optID_352
	if parentID_352, ok := optionIDMap["estate_type:rural_land:en"]; ok {
		parentID_352_ptr := parentID_352
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "undeveloped_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_352,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_352_ptr,
			"locale":      "en",
			"short_code":  "undevelo",
			"key":         "undeveloped_land",
			"label":       "Undeveloped land",
			"description": "",
			"value":       "Undeveloped land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_353 := uuid.New().String()
	optionIDMap["estate_subtype:undeveloped_land:es"] = optID_353
	if parentID_353, ok := optionIDMap["estate_type:rural_land:es"]; ok {
		parentID_353_ptr := parentID_353
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "undeveloped_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_353,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_353_ptr,
			"locale":      "es",
			"short_code":  "undevelo",
			"key":         "undeveloped_land",
			"label":       "Tierra no urbanizada",
			"description": "",
			"value":       "Undeveloped land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_354 := uuid.New().String()
	optionIDMap["estate_subtype:undeveloped_land:pl"] = optID_354
	if parentID_354, ok := optionIDMap["estate_type:rural_land:pl"]; ok {
		parentID_354_ptr := parentID_354
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "undeveloped_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_354,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_354_ptr,
			"locale":      "pl",
			"short_code":  "undevelo",
			"key":         "undeveloped_land",
			"label":       "Teren niezabudowany",
			"description": "",
			"value":       "Undeveloped land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_355 := uuid.New().String()
	optionIDMap["estate_subtype:beachfront:en"] = optID_355
	if parentID_355, ok := optionIDMap["estate_type:waterfront_land:en"]; ok {
		parentID_355_ptr := parentID_355
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "beachfront", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_355,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_355_ptr,
			"locale":      "en",
			"short_code":  "beachfro",
			"key":         "beachfront",
			"label":       "Beachfront lot",
			"description": "",
			"value":       "Beachfront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_356 := uuid.New().String()
	optionIDMap["estate_subtype:beachfront:es"] = optID_356
	if parentID_356, ok := optionIDMap["estate_type:waterfront_land:es"]; ok {
		parentID_356_ptr := parentID_356
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "beachfront", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_356,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_356_ptr,
			"locale":      "es",
			"short_code":  "beachfro",
			"key":         "beachfront",
			"label":       "Lote frente a playa",
			"description": "",
			"value":       "Beachfront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_357 := uuid.New().String()
	optionIDMap["estate_subtype:beachfront:pl"] = optID_357
	if parentID_357, ok := optionIDMap["estate_type:waterfront_land:pl"]; ok {
		parentID_357_ptr := parentID_357
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "beachfront", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_357,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_357_ptr,
			"locale":      "pl",
			"short_code":  "beachfro",
			"key":         "beachfront",
			"label":       "Działka przy plaży",
			"description": "",
			"value":       "Beachfront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_358 := uuid.New().String()
	optionIDMap["estate_subtype:lakefront:en"] = optID_358
	if parentID_358, ok := optionIDMap["estate_type:waterfront_land:en"]; ok {
		parentID_358_ptr := parentID_358
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "lakefront", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_358,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_358_ptr,
			"locale":      "en",
			"short_code":  "lakefron",
			"key":         "lakefront",
			"label":       "Lakefront lot",
			"description": "",
			"value":       "Lakefront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_359 := uuid.New().String()
	optionIDMap["estate_subtype:lakefront:es"] = optID_359
	if parentID_359, ok := optionIDMap["estate_type:waterfront_land:es"]; ok {
		parentID_359_ptr := parentID_359
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "lakefront", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_359,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_359_ptr,
			"locale":      "es",
			"short_code":  "lakefron",
			"key":         "lakefront",
			"label":       "Lote frente a lago",
			"description": "",
			"value":       "Lakefront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_360 := uuid.New().String()
	optionIDMap["estate_subtype:lakefront:pl"] = optID_360
	if parentID_360, ok := optionIDMap["estate_type:waterfront_land:pl"]; ok {
		parentID_360_ptr := parentID_360
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "lakefront", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_360,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_360_ptr,
			"locale":      "pl",
			"short_code":  "lakefron",
			"key":         "lakefront",
			"label":       "Działka przy jeziorze",
			"description": "",
			"value":       "Lakefront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_361 := uuid.New().String()
	optionIDMap["estate_subtype:riverfront:en"] = optID_361
	if parentID_361, ok := optionIDMap["estate_type:waterfront_land:en"]; ok {
		parentID_361_ptr := parentID_361
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "riverfront", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_361,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_361_ptr,
			"locale":      "en",
			"short_code":  "riverfro",
			"key":         "riverfront",
			"label":       "Riverfront lot",
			"description": "",
			"value":       "Riverfront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_362 := uuid.New().String()
	optionIDMap["estate_subtype:riverfront:es"] = optID_362
	if parentID_362, ok := optionIDMap["estate_type:waterfront_land:es"]; ok {
		parentID_362_ptr := parentID_362
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "riverfront", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_362,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_362_ptr,
			"locale":      "es",
			"short_code":  "riverfro",
			"key":         "riverfront",
			"label":       "Lote frente a río",
			"description": "",
			"value":       "Riverfront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_363 := uuid.New().String()
	optionIDMap["estate_subtype:riverfront:pl"] = optID_363
	if parentID_363, ok := optionIDMap["estate_type:waterfront_land:pl"]; ok {
		parentID_363_ptr := parentID_363
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "riverfront", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_363,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_363_ptr,
			"locale":      "pl",
			"short_code":  "riverfro",
			"key":         "riverfront",
			"label":       "Działka nad rzeką",
			"description": "",
			"value":       "Riverfront lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_364 := uuid.New().String()
	optionIDMap["estate_subtype:mountain:en"] = optID_364
	if parentID_364, ok := optionIDMap["estate_type:special_land:en"]; ok {
		parentID_364_ptr := parentID_364
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "mountain", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_364,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_364_ptr,
			"locale":      "en",
			"short_code":  "mountain",
			"key":         "mountain",
			"label":       "Mountain land",
			"description": "",
			"value":       "Mountain land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_365 := uuid.New().String()
	optionIDMap["estate_subtype:mountain:es"] = optID_365
	if parentID_365, ok := optionIDMap["estate_type:special_land:es"]; ok {
		parentID_365_ptr := parentID_365
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "mountain", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_365,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_365_ptr,
			"locale":      "es",
			"short_code":  "mountain",
			"key":         "mountain",
			"label":       "Terreno de montaña",
			"description": "",
			"value":       "Mountain land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_366 := uuid.New().String()
	optionIDMap["estate_subtype:mountain:pl"] = optID_366
	if parentID_366, ok := optionIDMap["estate_type:special_land:pl"]; ok {
		parentID_366_ptr := parentID_366
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "mountain", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_366,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_366_ptr,
			"locale":      "pl",
			"short_code":  "mountain",
			"key":         "mountain",
			"label":       "Teren górski",
			"description": "",
			"value":       "Mountain land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_367 := uuid.New().String()
	optionIDMap["estate_subtype:desert:en"] = optID_367
	if parentID_367, ok := optionIDMap["estate_type:special_land:en"]; ok {
		parentID_367_ptr := parentID_367
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "desert", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_367,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_367_ptr,
			"locale":      "en",
			"short_code":  "desert",
			"key":         "desert",
			"label":       "Desert land",
			"description": "",
			"value":       "Desert land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_368 := uuid.New().String()
	optionIDMap["estate_subtype:desert:es"] = optID_368
	if parentID_368, ok := optionIDMap["estate_type:special_land:es"]; ok {
		parentID_368_ptr := parentID_368
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "desert", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_368,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_368_ptr,
			"locale":      "es",
			"short_code":  "desert",
			"key":         "desert",
			"label":       "Terreno desértico",
			"description": "",
			"value":       "Desert land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_369 := uuid.New().String()
	optionIDMap["estate_subtype:desert:pl"] = optID_369
	if parentID_369, ok := optionIDMap["estate_type:special_land:pl"]; ok {
		parentID_369_ptr := parentID_369
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "desert", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_369,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_369_ptr,
			"locale":      "pl",
			"short_code":  "desert",
			"key":         "desert",
			"label":       "Teren pustynny",
			"description": "",
			"value":       "Desert land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_370 := uuid.New().String()
	optionIDMap["estate_subtype:raw_land:en"] = optID_370
	if parentID_370, ok := optionIDMap["estate_type:special_land:en"]; ok {
		parentID_370_ptr := parentID_370
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "raw_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_370,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_370_ptr,
			"locale":      "en",
			"short_code":  "raw_land",
			"key":         "raw_land",
			"label":       "Raw land",
			"description": "",
			"value":       "Raw land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_371 := uuid.New().String()
	optionIDMap["estate_subtype:raw_land:es"] = optID_371
	if parentID_371, ok := optionIDMap["estate_type:special_land:es"]; ok {
		parentID_371_ptr := parentID_371
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "raw_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_371,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_371_ptr,
			"locale":      "es",
			"short_code":  "raw_land",
			"key":         "raw_land",
			"label":       "Terreno virgen",
			"description": "",
			"value":       "Raw land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_372 := uuid.New().String()
	optionIDMap["estate_subtype:raw_land:pl"] = optID_372
	if parentID_372, ok := optionIDMap["estate_type:special_land:pl"]; ok {
		parentID_372_ptr := parentID_372
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "raw_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_372,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_372_ptr,
			"locale":      "pl",
			"short_code":  "raw_land",
			"key":         "raw_land",
			"label":       "Surowy teren",
			"description": "",
			"value":       "Raw land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_373 := uuid.New().String()
	optionIDMap["estate_subtype:improved_land:en"] = optID_373
	if parentID_373, ok := optionIDMap["estate_type:special_land:en"]; ok {
		parentID_373_ptr := parentID_373
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "improved_land", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_373,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_373_ptr,
			"locale":      "en",
			"short_code":  "improved",
			"key":         "improved_land",
			"label":       "Improved land",
			"description": "",
			"value":       "Improved land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_374 := uuid.New().String()
	optionIDMap["estate_subtype:improved_land:es"] = optID_374
	if parentID_374, ok := optionIDMap["estate_type:special_land:es"]; ok {
		parentID_374_ptr := parentID_374
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "improved_land", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_374,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_374_ptr,
			"locale":      "es",
			"short_code":  "improved",
			"key":         "improved_land",
			"label":       "Terreno mejorado",
			"description": "",
			"value":       "Improved land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_375 := uuid.New().String()
	optionIDMap["estate_subtype:improved_land:pl"] = optID_375
	if parentID_375, ok := optionIDMap["estate_type:special_land:pl"]; ok {
		parentID_375_ptr := parentID_375
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "improved_land", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_375,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_375_ptr,
			"locale":      "pl",
			"short_code":  "improved",
			"key":         "improved_land",
			"label":       "Ulepszony teren",
			"description": "",
			"value":       "Improved land",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_376 := uuid.New().String()
	optionIDMap["estate_subtype:crop_farm:en"] = optID_376
	if parentID_376, ok := optionIDMap["estate_type:farm:en"]; ok {
		parentID_376_ptr := parentID_376
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "crop_farm", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_376,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_376_ptr,
			"locale":      "en",
			"short_code":  "crop_far",
			"key":         "crop_farm",
			"label":       "Crop farm",
			"description": "",
			"value":       "Crop farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_377 := uuid.New().String()
	optionIDMap["estate_subtype:crop_farm:es"] = optID_377
	if parentID_377, ok := optionIDMap["estate_type:farm:es"]; ok {
		parentID_377_ptr := parentID_377
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "crop_farm", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_377,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_377_ptr,
			"locale":      "es",
			"short_code":  "crop_far",
			"key":         "crop_farm",
			"label":       "Granja de cultivos",
			"description": "",
			"value":       "Crop farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_378 := uuid.New().String()
	optionIDMap["estate_subtype:crop_farm:pl"] = optID_378
	if parentID_378, ok := optionIDMap["estate_type:farm:pl"]; ok {
		parentID_378_ptr := parentID_378
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "crop_farm", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_378,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_378_ptr,
			"locale":      "pl",
			"short_code":  "crop_far",
			"key":         "crop_farm",
			"label":       "Gospodarstwo rolne (uprawy)",
			"description": "",
			"value":       "Crop farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_379 := uuid.New().String()
	optionIDMap["estate_subtype:mixed_farm:en"] = optID_379
	if parentID_379, ok := optionIDMap["estate_type:farm:en"]; ok {
		parentID_379_ptr := parentID_379
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "mixed_farm", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_379,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_379_ptr,
			"locale":      "en",
			"short_code":  "mixed_fa",
			"key":         "mixed_farm",
			"label":       "Mixed farm",
			"description": "",
			"value":       "Mixed farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_380 := uuid.New().String()
	optionIDMap["estate_subtype:mixed_farm:es"] = optID_380
	if parentID_380, ok := optionIDMap["estate_type:farm:es"]; ok {
		parentID_380_ptr := parentID_380
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "mixed_farm", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_380,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_380_ptr,
			"locale":      "es",
			"short_code":  "mixed_fa",
			"key":         "mixed_farm",
			"label":       "Granja mixta",
			"description": "",
			"value":       "Mixed farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_381 := uuid.New().String()
	optionIDMap["estate_subtype:mixed_farm:pl"] = optID_381
	if parentID_381, ok := optionIDMap["estate_type:farm:pl"]; ok {
		parentID_381_ptr := parentID_381
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "mixed_farm", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_381,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_381_ptr,
			"locale":      "pl",
			"short_code":  "mixed_fa",
			"key":         "mixed_farm",
			"label":       "Gospodarstwo mieszane",
			"description": "",
			"value":       "Mixed farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_382 := uuid.New().String()
	optionIDMap["estate_subtype:organic_farm:en"] = optID_382
	if parentID_382, ok := optionIDMap["estate_type:farm:en"]; ok {
		parentID_382_ptr := parentID_382
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "organic_farm", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_382,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_382_ptr,
			"locale":      "en",
			"short_code":  "organic_",
			"key":         "organic_farm",
			"label":       "Organic farm",
			"description": "",
			"value":       "Organic farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_383 := uuid.New().String()
	optionIDMap["estate_subtype:organic_farm:es"] = optID_383
	if parentID_383, ok := optionIDMap["estate_type:farm:es"]; ok {
		parentID_383_ptr := parentID_383
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "organic_farm", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_383,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_383_ptr,
			"locale":      "es",
			"short_code":  "organic_",
			"key":         "organic_farm",
			"label":       "Granja orgánica",
			"description": "",
			"value":       "Organic farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_384 := uuid.New().String()
	optionIDMap["estate_subtype:organic_farm:pl"] = optID_384
	if parentID_384, ok := optionIDMap["estate_type:farm:pl"]; ok {
		parentID_384_ptr := parentID_384
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "organic_farm", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_384,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_384_ptr,
			"locale":      "pl",
			"short_code":  "organic_",
			"key":         "organic_farm",
			"label":       "Gospodarstwo ekologiczne",
			"description": "",
			"value":       "Organic farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_385 := uuid.New().String()
	optionIDMap["estate_subtype:cattle_ranch:en"] = optID_385
	if parentID_385, ok := optionIDMap["estate_type:ranch:en"]; ok {
		parentID_385_ptr := parentID_385
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "cattle_ranch", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_385,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_385_ptr,
			"locale":      "en",
			"short_code":  "cattle_r",
			"key":         "cattle_ranch",
			"label":       "Cattle ranch",
			"description": "",
			"value":       "Cattle ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_386 := uuid.New().String()
	optionIDMap["estate_subtype:cattle_ranch:es"] = optID_386
	if parentID_386, ok := optionIDMap["estate_type:ranch:es"]; ok {
		parentID_386_ptr := parentID_386
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "cattle_ranch", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_386,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_386_ptr,
			"locale":      "es",
			"short_code":  "cattle_r",
			"key":         "cattle_ranch",
			"label":       "Estancia ganadera",
			"description": "",
			"value":       "Cattle ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_387 := uuid.New().String()
	optionIDMap["estate_subtype:cattle_ranch:pl"] = optID_387
	if parentID_387, ok := optionIDMap["estate_type:ranch:pl"]; ok {
		parentID_387_ptr := parentID_387
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "cattle_ranch", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_387,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_387_ptr,
			"locale":      "pl",
			"short_code":  "cattle_r",
			"key":         "cattle_ranch",
			"label":       "Ranczo bydła",
			"description": "",
			"value":       "Cattle ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_388 := uuid.New().String()
	optionIDMap["estate_subtype:horse_ranch:en"] = optID_388
	if parentID_388, ok := optionIDMap["estate_type:ranch:en"]; ok {
		parentID_388_ptr := parentID_388
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "horse_ranch", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_388,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_388_ptr,
			"locale":      "en",
			"short_code":  "horse_ra",
			"key":         "horse_ranch",
			"label":       "Horse ranch",
			"description": "",
			"value":       "Horse ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_389 := uuid.New().String()
	optionIDMap["estate_subtype:horse_ranch:es"] = optID_389
	if parentID_389, ok := optionIDMap["estate_type:ranch:es"]; ok {
		parentID_389_ptr := parentID_389
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "horse_ranch", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_389,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_389_ptr,
			"locale":      "es",
			"short_code":  "horse_ra",
			"key":         "horse_ranch",
			"label":       "Haras / Rancho de caballos",
			"description": "",
			"value":       "Horse ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_390 := uuid.New().String()
	optionIDMap["estate_subtype:horse_ranch:pl"] = optID_390
	if parentID_390, ok := optionIDMap["estate_type:ranch:pl"]; ok {
		parentID_390_ptr := parentID_390
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "horse_ranch", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_390,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_390_ptr,
			"locale":      "pl",
			"short_code":  "horse_ra",
			"key":         "horse_ranch",
			"label":       "Ranczo konne",
			"description": "",
			"value":       "Horse ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_391 := uuid.New().String()
	optionIDMap["estate_subtype:working_ranch:en"] = optID_391
	if parentID_391, ok := optionIDMap["estate_type:ranch:en"]; ok {
		parentID_391_ptr := parentID_391
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "working_ranch", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_391,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_391_ptr,
			"locale":      "en",
			"short_code":  "working_",
			"key":         "working_ranch",
			"label":       "Working ranch",
			"description": "",
			"value":       "Working ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_392 := uuid.New().String()
	optionIDMap["estate_subtype:working_ranch:es"] = optID_392
	if parentID_392, ok := optionIDMap["estate_type:ranch:es"]; ok {
		parentID_392_ptr := parentID_392
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "working_ranch", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_392,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_392_ptr,
			"locale":      "es",
			"short_code":  "working_",
			"key":         "working_ranch",
			"label":       "Rancho operativo",
			"description": "",
			"value":       "Working ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_393 := uuid.New().String()
	optionIDMap["estate_subtype:working_ranch:pl"] = optID_393
	if parentID_393, ok := optionIDMap["estate_type:ranch:pl"]; ok {
		parentID_393_ptr := parentID_393
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "working_ranch", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_393,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_393_ptr,
			"locale":      "pl",
			"short_code":  "working_",
			"key":         "working_ranch",
			"label":       "Ranczo produkcyjne",
			"description": "",
			"value":       "Working ranch",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_394 := uuid.New().String()
	optionIDMap["estate_subtype:orchard:en"] = optID_394
	if parentID_394, ok := optionIDMap["estate_type:agri_specialty:en"]; ok {
		parentID_394_ptr := parentID_394
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "orchard", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_394,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_394_ptr,
			"locale":      "en",
			"short_code":  "orchard",
			"key":         "orchard",
			"label":       "Orchard",
			"description": "",
			"value":       "Orchard",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_395 := uuid.New().String()
	optionIDMap["estate_subtype:orchard:es"] = optID_395
	if parentID_395, ok := optionIDMap["estate_type:agri_specialty:es"]; ok {
		parentID_395_ptr := parentID_395
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "orchard", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_395,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_395_ptr,
			"locale":      "es",
			"short_code":  "orchard",
			"key":         "orchard",
			"label":       "Huerto frutal",
			"description": "",
			"value":       "Orchard",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_396 := uuid.New().String()
	optionIDMap["estate_subtype:orchard:pl"] = optID_396
	if parentID_396, ok := optionIDMap["estate_type:agri_specialty:pl"]; ok {
		parentID_396_ptr := parentID_396
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "orchard", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_396,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_396_ptr,
			"locale":      "pl",
			"short_code":  "orchard",
			"key":         "orchard",
			"label":       "Sad",
			"description": "",
			"value":       "Orchard",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_397 := uuid.New().String()
	optionIDMap["estate_subtype:vineyard:en"] = optID_397
	if parentID_397, ok := optionIDMap["estate_type:agri_specialty:en"]; ok {
		parentID_397_ptr := parentID_397
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "vineyard", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_397,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_397_ptr,
			"locale":      "en",
			"short_code":  "vineyard",
			"key":         "vineyard",
			"label":       "Vineyard",
			"description": "",
			"value":       "Vineyard",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_398 := uuid.New().String()
	optionIDMap["estate_subtype:vineyard:es"] = optID_398
	if parentID_398, ok := optionIDMap["estate_type:agri_specialty:es"]; ok {
		parentID_398_ptr := parentID_398
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "vineyard", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_398,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_398_ptr,
			"locale":      "es",
			"short_code":  "vineyard",
			"key":         "vineyard",
			"label":       "Viñedo",
			"description": "",
			"value":       "Vineyard",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_399 := uuid.New().String()
	optionIDMap["estate_subtype:vineyard:pl"] = optID_399
	if parentID_399, ok := optionIDMap["estate_type:agri_specialty:pl"]; ok {
		parentID_399_ptr := parentID_399
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "vineyard", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_399,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_399_ptr,
			"locale":      "pl",
			"short_code":  "vineyard",
			"key":         "vineyard",
			"label":       "Winnica",
			"description": "",
			"value":       "Vineyard",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_400 := uuid.New().String()
	optionIDMap["estate_subtype:greenhouse:en"] = optID_400
	if parentID_400, ok := optionIDMap["estate_type:agri_specialty:en"]; ok {
		parentID_400_ptr := parentID_400
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "greenhouse", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_400,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_400_ptr,
			"locale":      "en",
			"short_code":  "greenhou",
			"key":         "greenhouse",
			"label":       "Greenhouse",
			"description": "",
			"value":       "Greenhouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_401 := uuid.New().String()
	optionIDMap["estate_subtype:greenhouse:es"] = optID_401
	if parentID_401, ok := optionIDMap["estate_type:agri_specialty:es"]; ok {
		parentID_401_ptr := parentID_401
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "greenhouse", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_401,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_401_ptr,
			"locale":      "es",
			"short_code":  "greenhou",
			"key":         "greenhouse",
			"label":       "Invernadero",
			"description": "",
			"value":       "Greenhouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_402 := uuid.New().String()
	optionIDMap["estate_subtype:greenhouse:pl"] = optID_402
	if parentID_402, ok := optionIDMap["estate_type:agri_specialty:pl"]; ok {
		parentID_402_ptr := parentID_402
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "greenhouse", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_402,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_402_ptr,
			"locale":      "pl",
			"short_code":  "greenhou",
			"key":         "greenhouse",
			"label":       "Szklarnia",
			"description": "",
			"value":       "Greenhouse",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_403 := uuid.New().String()
	optionIDMap["estate_subtype:fishery:en"] = optID_403
	if parentID_403, ok := optionIDMap["estate_type:agri_specialty:en"]; ok {
		parentID_403_ptr := parentID_403
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "fishery", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_403,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_403_ptr,
			"locale":      "en",
			"short_code":  "fishery",
			"key":         "fishery",
			"label":       "Fishery",
			"description": "",
			"value":       "Fishery",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_404 := uuid.New().String()
	optionIDMap["estate_subtype:fishery:es"] = optID_404
	if parentID_404, ok := optionIDMap["estate_type:agri_specialty:es"]; ok {
		parentID_404_ptr := parentID_404
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "fishery", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_404,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_404_ptr,
			"locale":      "es",
			"short_code":  "fishery",
			"key":         "fishery",
			"label":       "Pesquera",
			"description": "",
			"value":       "Fishery",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_405 := uuid.New().String()
	optionIDMap["estate_subtype:fishery:pl"] = optID_405
	if parentID_405, ok := optionIDMap["estate_type:agri_specialty:pl"]; ok {
		parentID_405_ptr := parentID_405
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "fishery", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_405,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_405_ptr,
			"locale":      "pl",
			"short_code":  "fishery",
			"key":         "fishery",
			"label":       "Gospodarstwo rybne",
			"description": "",
			"value":       "Fishery",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_406 := uuid.New().String()
	optionIDMap["estate_subtype:equestrian_estate:en"] = optID_406
	if parentID_406, ok := optionIDMap["estate_type:agri_specialty:en"]; ok {
		parentID_406_ptr := parentID_406
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "equestrian_estate", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_406,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_406_ptr,
			"locale":      "en",
			"short_code":  "equestri",
			"key":         "equestrian_estate",
			"label":       "Equestrian estate",
			"description": "",
			"value":       "Equestrian estate",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_407 := uuid.New().String()
	optionIDMap["estate_subtype:equestrian_estate:es"] = optID_407
	if parentID_407, ok := optionIDMap["estate_type:agri_specialty:es"]; ok {
		parentID_407_ptr := parentID_407
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "equestrian_estate", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_407,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_407_ptr,
			"locale":      "es",
			"short_code":  "equestri",
			"key":         "equestrian_estate",
			"label":       "Hípico / Ecuestre",
			"description": "",
			"value":       "Equestrian estate",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_408 := uuid.New().String()
	optionIDMap["estate_subtype:equestrian_estate:pl"] = optID_408
	if parentID_408, ok := optionIDMap["estate_type:agri_specialty:pl"]; ok {
		parentID_408_ptr := parentID_408
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "equestrian_estate", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_408,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_408_ptr,
			"locale":      "pl",
			"short_code":  "equestri",
			"key":         "equestrian_estate",
			"label":       "Posiadłość jeździecka",
			"description": "",
			"value":       "Equestrian estate",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_409 := uuid.New().String()
	optionIDMap["estate_subtype:marina:en"] = optID_409
	if parentID_409, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_409_ptr := parentID_409
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "marina", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_409,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_409_ptr,
			"locale":      "en",
			"short_code":  "marina",
			"key":         "marina",
			"label":       "Marina",
			"description": "",
			"value":       "Marina",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_410 := uuid.New().String()
	optionIDMap["estate_subtype:marina:es"] = optID_410
	if parentID_410, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_410_ptr := parentID_410
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "marina", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_410,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_410_ptr,
			"locale":      "es",
			"short_code":  "marina",
			"key":         "marina",
			"label":       "Marina",
			"description": "",
			"value":       "Marina",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_411 := uuid.New().String()
	optionIDMap["estate_subtype:marina:pl"] = optID_411
	if parentID_411, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_411_ptr := parentID_411
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "marina", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_411,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_411_ptr,
			"locale":      "pl",
			"short_code":  "marina",
			"key":         "marina",
			"label":       "Marina",
			"description": "",
			"value":       "Marina",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_412 := uuid.New().String()
	optionIDMap["estate_subtype:boat_slip:en"] = optID_412
	if parentID_412, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_412_ptr := parentID_412
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "boat_slip", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_412,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_412_ptr,
			"locale":      "en",
			"short_code":  "boat_sli",
			"key":         "boat_slip",
			"label":       "Boat slip",
			"description": "",
			"value":       "Boat slip",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_413 := uuid.New().String()
	optionIDMap["estate_subtype:boat_slip:es"] = optID_413
	if parentID_413, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_413_ptr := parentID_413
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "boat_slip", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_413,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_413_ptr,
			"locale":      "es",
			"short_code":  "boat_sli",
			"key":         "boat_slip",
			"label":       "Amarra",
			"description": "",
			"value":       "Boat slip",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_414 := uuid.New().String()
	optionIDMap["estate_subtype:boat_slip:pl"] = optID_414
	if parentID_414, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_414_ptr := parentID_414
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "boat_slip", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_414,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_414_ptr,
			"locale":      "pl",
			"short_code":  "boat_sli",
			"key":         "boat_slip",
			"label":       "Miejsce postojowe (łódź)",
			"description": "",
			"value":       "Boat slip",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_415 := uuid.New().String()
	optionIDMap["estate_subtype:dock:en"] = optID_415
	if parentID_415, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_415_ptr := parentID_415
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "dock", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_415,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_415_ptr,
			"locale":      "en",
			"short_code":  "dock",
			"key":         "dock",
			"label":       "Dock",
			"description": "",
			"value":       "Dock",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_416 := uuid.New().String()
	optionIDMap["estate_subtype:dock:es"] = optID_416
	if parentID_416, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_416_ptr := parentID_416
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "dock", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_416,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_416_ptr,
			"locale":      "es",
			"short_code":  "dock",
			"key":         "dock",
			"label":       "Muelle",
			"description": "",
			"value":       "Dock",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_417 := uuid.New().String()
	optionIDMap["estate_subtype:dock:pl"] = optID_417
	if parentID_417, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_417_ptr := parentID_417
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "dock", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_417,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_417_ptr,
			"locale":      "pl",
			"short_code":  "dock",
			"key":         "dock",
			"label":       "Pomost",
			"description": "",
			"value":       "Dock",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_418 := uuid.New().String()
	optionIDMap["estate_subtype:hangar:en"] = optID_418
	if parentID_418, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_418_ptr := parentID_418
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "hangar", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_418,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_418_ptr,
			"locale":      "en",
			"short_code":  "hangar",
			"key":         "hangar",
			"label":       "Hangar (airplane)",
			"description": "",
			"value":       "Hangar (airplane)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_419 := uuid.New().String()
	optionIDMap["estate_subtype:hangar:es"] = optID_419
	if parentID_419, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_419_ptr := parentID_419
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "hangar", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_419,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_419_ptr,
			"locale":      "es",
			"short_code":  "hangar",
			"key":         "hangar",
			"label":       "Hangar (aviones)",
			"description": "",
			"value":       "Hangar (airplane)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_420 := uuid.New().String()
	optionIDMap["estate_subtype:hangar:pl"] = optID_420
	if parentID_420, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_420_ptr := parentID_420
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "hangar", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_420,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_420_ptr,
			"locale":      "pl",
			"short_code":  "hangar",
			"key":         "hangar",
			"label":       "Hangar",
			"description": "",
			"value":       "Hangar (airplane)",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_421 := uuid.New().String()
	optionIDMap["estate_subtype:railway:en"] = optID_421
	if parentID_421, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_421_ptr := parentID_421
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "railway", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_421,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_421_ptr,
			"locale":      "en",
			"short_code":  "railway",
			"key":         "railway",
			"label":       "Railway property",
			"description": "",
			"value":       "Railway property",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_422 := uuid.New().String()
	optionIDMap["estate_subtype:railway:es"] = optID_422
	if parentID_422, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_422_ptr := parentID_422
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "railway", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_422,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_422_ptr,
			"locale":      "es",
			"short_code":  "railway",
			"key":         "railway",
			"label":       "Propiedad ferroviaria",
			"description": "",
			"value":       "Railway property",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_423 := uuid.New().String()
	optionIDMap["estate_subtype:railway:pl"] = optID_423
	if parentID_423, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_423_ptr := parentID_423
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "railway", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_423,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_423_ptr,
			"locale":      "pl",
			"short_code":  "railway",
			"key":         "railway",
			"label":       "Nieruchomość kolejowa",
			"description": "",
			"value":       "Railway property",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_424 := uuid.New().String()
	optionIDMap["estate_subtype:parking_lot:en"] = optID_424
	if parentID_424, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_424_ptr := parentID_424
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "parking_lot", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_424,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_424_ptr,
			"locale":      "en",
			"short_code":  "parking_",
			"key":         "parking_lot",
			"label":       "Parking lot",
			"description": "",
			"value":       "Parking lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_425 := uuid.New().String()
	optionIDMap["estate_subtype:parking_lot:es"] = optID_425
	if parentID_425, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_425_ptr := parentID_425
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "parking_lot", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_425,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_425_ptr,
			"locale":      "es",
			"short_code":  "parking_",
			"key":         "parking_lot",
			"label":       "Playa de estacionamiento",
			"description": "",
			"value":       "Parking lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_426 := uuid.New().String()
	optionIDMap["estate_subtype:parking_lot:pl"] = optID_426
	if parentID_426, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_426_ptr := parentID_426
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "parking_lot", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_426,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_426_ptr,
			"locale":      "pl",
			"short_code":  "parking_",
			"key":         "parking_lot",
			"label":       "Parking naziemny",
			"description": "",
			"value":       "Parking lot",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_427 := uuid.New().String()
	optionIDMap["estate_subtype:parking_garage:en"] = optID_427
	if parentID_427, ok := optionIDMap["estate_type:transportation:en"]; ok {
		parentID_427_ptr := parentID_427
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "parking_garage", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_427,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_427_ptr,
			"locale":      "en",
			"short_code":  "parking_",
			"key":         "parking_garage",
			"label":       "Parking garage",
			"description": "",
			"value":       "Parking garage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_428 := uuid.New().String()
	optionIDMap["estate_subtype:parking_garage:es"] = optID_428
	if parentID_428, ok := optionIDMap["estate_type:transportation:es"]; ok {
		parentID_428_ptr := parentID_428
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "parking_garage", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_428,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_428_ptr,
			"locale":      "es",
			"short_code":  "parking_",
			"key":         "parking_garage",
			"label":       "Estacionamiento cubierto",
			"description": "",
			"value":       "Parking garage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_429 := uuid.New().String()
	optionIDMap["estate_subtype:parking_garage:pl"] = optID_429
	if parentID_429, ok := optionIDMap["estate_type:transportation:pl"]; ok {
		parentID_429_ptr := parentID_429
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "parking_garage", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_429,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_429_ptr,
			"locale":      "pl",
			"short_code":  "parking_",
			"key":         "parking_garage",
			"label":       "Parking podziemny/garaz",
			"description": "",
			"value":       "Parking garage",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_430 := uuid.New().String()
	optionIDMap["estate_subtype:power_station:en"] = optID_430
	if parentID_430, ok := optionIDMap["estate_type:utilities:en"]; ok {
		parentID_430_ptr := parentID_430
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "power_station", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_430,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_430_ptr,
			"locale":      "en",
			"short_code":  "power_st",
			"key":         "power_station",
			"label":       "Power station",
			"description": "",
			"value":       "Power station",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_431 := uuid.New().String()
	optionIDMap["estate_subtype:power_station:es"] = optID_431
	if parentID_431, ok := optionIDMap["estate_type:utilities:es"]; ok {
		parentID_431_ptr := parentID_431
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "power_station", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_431,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_431_ptr,
			"locale":      "es",
			"short_code":  "power_st",
			"key":         "power_station",
			"label":       "Central eléctrica",
			"description": "",
			"value":       "Power station",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_432 := uuid.New().String()
	optionIDMap["estate_subtype:power_station:pl"] = optID_432
	if parentID_432, ok := optionIDMap["estate_type:utilities:pl"]; ok {
		parentID_432_ptr := parentID_432
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "power_station", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_432,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_432_ptr,
			"locale":      "pl",
			"short_code":  "power_st",
			"key":         "power_station",
			"label":       "Elektrownia",
			"description": "",
			"value":       "Power station",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_433 := uuid.New().String()
	optionIDMap["estate_subtype:water_tower:en"] = optID_433
	if parentID_433, ok := optionIDMap["estate_type:utilities:en"]; ok {
		parentID_433_ptr := parentID_433
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "water_tower", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_433,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_433_ptr,
			"locale":      "en",
			"short_code":  "water_to",
			"key":         "water_tower",
			"label":       "Water tower",
			"description": "",
			"value":       "Water tower",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_434 := uuid.New().String()
	optionIDMap["estate_subtype:water_tower:es"] = optID_434
	if parentID_434, ok := optionIDMap["estate_type:utilities:es"]; ok {
		parentID_434_ptr := parentID_434
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "water_tower", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_434,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_434_ptr,
			"locale":      "es",
			"short_code":  "water_to",
			"key":         "water_tower",
			"label":       "Torre de agua",
			"description": "",
			"value":       "Water tower",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_435 := uuid.New().String()
	optionIDMap["estate_subtype:water_tower:pl"] = optID_435
	if parentID_435, ok := optionIDMap["estate_type:utilities:pl"]; ok {
		parentID_435_ptr := parentID_435
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "water_tower", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_435,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_435_ptr,
			"locale":      "pl",
			"short_code":  "water_to",
			"key":         "water_tower",
			"label":       "Wieża ciśnień",
			"description": "",
			"value":       "Water tower",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_436 := uuid.New().String()
	optionIDMap["estate_subtype:wind_farm:en"] = optID_436
	if parentID_436, ok := optionIDMap["estate_type:utilities:en"]; ok {
		parentID_436_ptr := parentID_436
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "wind_farm", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_436,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_436_ptr,
			"locale":      "en",
			"short_code":  "wind_far",
			"key":         "wind_farm",
			"label":       "Wind farm",
			"description": "",
			"value":       "Wind farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_437 := uuid.New().String()
	optionIDMap["estate_subtype:wind_farm:es"] = optID_437
	if parentID_437, ok := optionIDMap["estate_type:utilities:es"]; ok {
		parentID_437_ptr := parentID_437
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "wind_farm", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_437,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_437_ptr,
			"locale":      "es",
			"short_code":  "wind_far",
			"key":         "wind_farm",
			"label":       "Parque eólico",
			"description": "",
			"value":       "Wind farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_438 := uuid.New().String()
	optionIDMap["estate_subtype:wind_farm:pl"] = optID_438
	if parentID_438, ok := optionIDMap["estate_type:utilities:pl"]; ok {
		parentID_438_ptr := parentID_438
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "wind_farm", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_438,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_438_ptr,
			"locale":      "pl",
			"short_code":  "wind_far",
			"key":         "wind_farm",
			"label":       "Farma wiatrowa",
			"description": "",
			"value":       "Wind farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_439 := uuid.New().String()
	optionIDMap["estate_subtype:solar_farm:en"] = optID_439
	if parentID_439, ok := optionIDMap["estate_type:utilities:en"]; ok {
		parentID_439_ptr := parentID_439
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "solar_farm", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_439,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_439_ptr,
			"locale":      "en",
			"short_code":  "solar_fa",
			"key":         "solar_farm",
			"label":       "Solar farm",
			"description": "",
			"value":       "Solar farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_440 := uuid.New().String()
	optionIDMap["estate_subtype:solar_farm:es"] = optID_440
	if parentID_440, ok := optionIDMap["estate_type:utilities:es"]; ok {
		parentID_440_ptr := parentID_440
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "solar_farm", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_440,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_440_ptr,
			"locale":      "es",
			"short_code":  "solar_fa",
			"key":         "solar_farm",
			"label":       "Parque solar",
			"description": "",
			"value":       "Solar farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_441 := uuid.New().String()
	optionIDMap["estate_subtype:solar_farm:pl"] = optID_441
	if parentID_441, ok := optionIDMap["estate_type:utilities:pl"]; ok {
		parentID_441_ptr := parentID_441
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "solar_farm", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_441,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_441_ptr,
			"locale":      "pl",
			"short_code":  "solar_fa",
			"key":         "solar_farm",
			"label":       "Farma fotowoltaiczna",
			"description": "",
			"value":       "Solar farm",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_442 := uuid.New().String()
	optionIDMap["estate_subtype:telecom_tower:en"] = optID_442
	if parentID_442, ok := optionIDMap["estate_type:utilities:en"]; ok {
		parentID_442_ptr := parentID_442
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "telecom_tower", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_442,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_442_ptr,
			"locale":      "en",
			"short_code":  "telecom_",
			"key":         "telecom_tower",
			"label":       "Telecom tower",
			"description": "",
			"value":       "Telecom tower",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_443 := uuid.New().String()
	optionIDMap["estate_subtype:telecom_tower:es"] = optID_443
	if parentID_443, ok := optionIDMap["estate_type:utilities:es"]; ok {
		parentID_443_ptr := parentID_443
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "telecom_tower", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_443,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_443_ptr,
			"locale":      "es",
			"short_code":  "telecom_",
			"key":         "telecom_tower",
			"label":       "Torre de telecomunicaciones",
			"description": "",
			"value":       "Telecom tower",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_444 := uuid.New().String()
	optionIDMap["estate_subtype:telecom_tower:pl"] = optID_444
	if parentID_444, ok := optionIDMap["estate_type:utilities:pl"]; ok {
		parentID_444_ptr := parentID_444
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "telecom_tower", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_444,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_444_ptr,
			"locale":      "pl",
			"short_code":  "telecom_",
			"key":         "telecom_tower",
			"label":       "Maszt telekomunikacyjny",
			"description": "",
			"value":       "Telecom tower",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_445 := uuid.New().String()
	optionIDMap["estate_subtype:military:en"] = optID_445
	if parentID_445, ok := optionIDMap["estate_type:institutional:en"]; ok {
		parentID_445_ptr := parentID_445
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "military", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_445,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_445_ptr,
			"locale":      "en",
			"short_code":  "military",
			"key":         "military",
			"label":       "Military facility",
			"description": "",
			"value":       "Military facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_446 := uuid.New().String()
	optionIDMap["estate_subtype:military:es"] = optID_446
	if parentID_446, ok := optionIDMap["estate_type:institutional:es"]; ok {
		parentID_446_ptr := parentID_446
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "military", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_446,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_446_ptr,
			"locale":      "es",
			"short_code":  "military",
			"key":         "military",
			"label":       "Instalación militar",
			"description": "",
			"value":       "Military facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_447 := uuid.New().String()
	optionIDMap["estate_subtype:military:pl"] = optID_447
	if parentID_447, ok := optionIDMap["estate_type:institutional:pl"]; ok {
		parentID_447_ptr := parentID_447
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "military", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_447,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_447_ptr,
			"locale":      "pl",
			"short_code":  "military",
			"key":         "military",
			"label":       "Obiekt wojskowy",
			"description": "",
			"value":       "Military facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_448 := uuid.New().String()
	optionIDMap["estate_subtype:correctional:en"] = optID_448
	if parentID_448, ok := optionIDMap["estate_type:institutional:en"]; ok {
		parentID_448_ptr := parentID_448
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "correctional", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_448,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_448_ptr,
			"locale":      "en",
			"short_code":  "correcti",
			"key":         "correctional",
			"label":       "Correctional facility",
			"description": "",
			"value":       "Correctional facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_449 := uuid.New().String()
	optionIDMap["estate_subtype:correctional:es"] = optID_449
	if parentID_449, ok := optionIDMap["estate_type:institutional:es"]; ok {
		parentID_449_ptr := parentID_449
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "correctional", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_449,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_449_ptr,
			"locale":      "es",
			"short_code":  "correcti",
			"key":         "correctional",
			"label":       "Cárcel / Penal",
			"description": "",
			"value":       "Correctional facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_450 := uuid.New().String()
	optionIDMap["estate_subtype:correctional:pl"] = optID_450
	if parentID_450, ok := optionIDMap["estate_type:institutional:pl"]; ok {
		parentID_450_ptr := parentID_450
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "correctional", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_450,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_450_ptr,
			"locale":      "pl",
			"short_code":  "correcti",
			"key":         "correctional",
			"label":       "Zakład karny",
			"description": "",
			"value":       "Correctional facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_451 := uuid.New().String()
	optionIDMap["estate_subtype:embassy:en"] = optID_451
	if parentID_451, ok := optionIDMap["estate_type:institutional:en"]; ok {
		parentID_451_ptr := parentID_451
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "embassy", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_451,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_451_ptr,
			"locale":      "en",
			"short_code":  "embassy",
			"key":         "embassy",
			"label":       "Embassy / Consulate",
			"description": "",
			"value":       "Embassy / Consulate",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_452 := uuid.New().String()
	optionIDMap["estate_subtype:embassy:es"] = optID_452
	if parentID_452, ok := optionIDMap["estate_type:institutional:es"]; ok {
		parentID_452_ptr := parentID_452
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "embassy", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_452,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_452_ptr,
			"locale":      "es",
			"short_code":  "embassy",
			"key":         "embassy",
			"label":       "Embajada / Consulado",
			"description": "",
			"value":       "Embassy / Consulate",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_453 := uuid.New().String()
	optionIDMap["estate_subtype:embassy:pl"] = optID_453
	if parentID_453, ok := optionIDMap["estate_type:institutional:pl"]; ok {
		parentID_453_ptr := parentID_453
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "embassy", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_453,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_453_ptr,
			"locale":      "pl",
			"short_code":  "embassy",
			"key":         "embassy",
			"label":       "Ambasada / Konsulat",
			"description": "",
			"value":       "Embassy / Consulate",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_454 := uuid.New().String()
	optionIDMap["estate_subtype:library:en"] = optID_454
	if parentID_454, ok := optionIDMap["estate_type:institutional:en"]; ok {
		parentID_454_ptr := parentID_454
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "library", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_454,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_454_ptr,
			"locale":      "en",
			"short_code":  "library",
			"key":         "library",
			"label":       "Library",
			"description": "",
			"value":       "Library",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_455 := uuid.New().String()
	optionIDMap["estate_subtype:library:es"] = optID_455
	if parentID_455, ok := optionIDMap["estate_type:institutional:es"]; ok {
		parentID_455_ptr := parentID_455
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "library", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_455,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_455_ptr,
			"locale":      "es",
			"short_code":  "library",
			"key":         "library",
			"label":       "Biblioteca",
			"description": "",
			"value":       "Library",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_456 := uuid.New().String()
	optionIDMap["estate_subtype:library:pl"] = optID_456
	if parentID_456, ok := optionIDMap["estate_type:institutional:pl"]; ok {
		parentID_456_ptr := parentID_456
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "library", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_456,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_456_ptr,
			"locale":      "pl",
			"short_code":  "library",
			"key":         "library",
			"label":       "Biblioteka",
			"description": "",
			"value":       "Library",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_457 := uuid.New().String()
	optionIDMap["estate_subtype:post_office:en"] = optID_457
	if parentID_457, ok := optionIDMap["estate_type:institutional:en"]; ok {
		parentID_457_ptr := parentID_457
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "post_office", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_457,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_457_ptr,
			"locale":      "en",
			"short_code":  "post_off",
			"key":         "post_office",
			"label":       "Post office",
			"description": "",
			"value":       "Post office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_458 := uuid.New().String()
	optionIDMap["estate_subtype:post_office:es"] = optID_458
	if parentID_458, ok := optionIDMap["estate_type:institutional:es"]; ok {
		parentID_458_ptr := parentID_458
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "post_office", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_458,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_458_ptr,
			"locale":      "es",
			"short_code":  "post_off",
			"key":         "post_office",
			"label":       "Correo",
			"description": "",
			"value":       "Post office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_459 := uuid.New().String()
	optionIDMap["estate_subtype:post_office:pl"] = optID_459
	if parentID_459, ok := optionIDMap["estate_type:institutional:pl"]; ok {
		parentID_459_ptr := parentID_459
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "post_office", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_459,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_459_ptr,
			"locale":      "pl",
			"short_code":  "post_off",
			"key":         "post_office",
			"label":       "Poczta",
			"description": "",
			"value":       "Post office",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_460 := uuid.New().String()
	optionIDMap["estate_subtype:event_venue:en"] = optID_460
	if parentID_460, ok := optionIDMap["estate_type:recreational:en"]; ok {
		parentID_460_ptr := parentID_460
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "event_venue", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_460,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_460_ptr,
			"locale":      "en",
			"short_code":  "event_ve",
			"key":         "event_venue",
			"label":       "Event venue",
			"description": "",
			"value":       "Event venue",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_461 := uuid.New().String()
	optionIDMap["estate_subtype:event_venue:es"] = optID_461
	if parentID_461, ok := optionIDMap["estate_type:recreational:es"]; ok {
		parentID_461_ptr := parentID_461
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "event_venue", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_461,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_461_ptr,
			"locale":      "es",
			"short_code":  "event_ve",
			"key":         "event_venue",
			"label":       "Salón de eventos",
			"description": "",
			"value":       "Event venue",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_462 := uuid.New().String()
	optionIDMap["estate_subtype:event_venue:pl"] = optID_462
	if parentID_462, ok := optionIDMap["estate_type:recreational:pl"]; ok {
		parentID_462_ptr := parentID_462
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "event_venue", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_462,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_462_ptr,
			"locale":      "pl",
			"short_code":  "event_ve",
			"key":         "event_venue",
			"label":       "Sala eventowa",
			"description": "",
			"value":       "Event venue",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_463 := uuid.New().String()
	optionIDMap["estate_subtype:conference_center:en"] = optID_463
	if parentID_463, ok := optionIDMap["estate_type:recreational:en"]; ok {
		parentID_463_ptr := parentID_463
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "conference_center", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_463,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_463_ptr,
			"locale":      "en",
			"short_code":  "conferen",
			"key":         "conference_center",
			"label":       "Conference center",
			"description": "",
			"value":       "Conference center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_464 := uuid.New().String()
	optionIDMap["estate_subtype:conference_center:es"] = optID_464
	if parentID_464, ok := optionIDMap["estate_type:recreational:es"]; ok {
		parentID_464_ptr := parentID_464
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "conference_center", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_464,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_464_ptr,
			"locale":      "es",
			"short_code":  "conferen",
			"key":         "conference_center",
			"label":       "Centro de convenciones",
			"description": "",
			"value":       "Conference center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_465 := uuid.New().String()
	optionIDMap["estate_subtype:conference_center:pl"] = optID_465
	if parentID_465, ok := optionIDMap["estate_type:recreational:pl"]; ok {
		parentID_465_ptr := parentID_465
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "conference_center", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_465,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_465_ptr,
			"locale":      "pl",
			"short_code":  "conferen",
			"key":         "conference_center",
			"label":       "Centrum konferencyjne",
			"description": "",
			"value":       "Conference center",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_466 := uuid.New().String()
	optionIDMap["estate_subtype:amusement:en"] = optID_466
	if parentID_466, ok := optionIDMap["estate_type:recreational:en"]; ok {
		parentID_466_ptr := parentID_466
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:en"], "key": "amusement", "locale": "en"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_466,
			"set_id":      setIDMap["estate_subtype:en"],
			"parent_id":   &parentID_466_ptr,
			"locale":      "en",
			"short_code":  "amusemen",
			"key":         "amusement",
			"label":       "Amusement facility",
			"description": "",
			"value":       "Amusement facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_467 := uuid.New().String()
	optionIDMap["estate_subtype:amusement:es"] = optID_467
	if parentID_467, ok := optionIDMap["estate_type:recreational:es"]; ok {
		parentID_467_ptr := parentID_467
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:es"], "key": "amusement", "locale": "es"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_467,
			"set_id":      setIDMap["estate_subtype:es"],
			"parent_id":   &parentID_467_ptr,
			"locale":      "es",
			"short_code":  "amusemen",
			"key":         "amusement",
			"label":       "Parque de diversiones",
			"description": "",
			"value":       "Amusement facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	optID_468 := uuid.New().String()
	optionIDMap["estate_subtype:amusement:pl"] = optID_468
	if parentID_468, ok := optionIDMap["estate_type:recreational:pl"]; ok {
		parentID_468_ptr := parentID_468
		_, _ = optionsCollection.UpdateOne(ctx, bson.M{"set_id": setIDMap["estate_subtype:pl"], "key": "amusement", "locale": "pl"}, bson.M{"$setOnInsert": bson.M{
			"_id":         optID_468,
			"set_id":      setIDMap["estate_subtype:pl"],
			"parent_id":   &parentID_468_ptr,
			"locale":      "pl",
			"short_code":  "amusemen",
			"key":         "amusement",
			"label":       "Park rozrywki",
			"description": "",
			"value":       "Amusement facility",
			"order":       0,
			"active":      true,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
			"created_by":  "system",
			"updated_by":  "system",
		}}, options.Update().SetUpsert(true))
	}

	if err := seedMediaDictionary(ctx, setsCollection, optionsCollection); err != nil {
		return err
	}

	return nil
}

type mediaSeedOption struct {
	Key    string
	Labels map[string]string
}

type mediaSeedSet struct {
	Name    string
	Labels  map[string]string
	Options []mediaSeedOption
}

var mediaSeedSets = []mediaSeedSet{
	{Name: "media_location", Labels: map[string]string{"en": "Media Location", "pl": "Lokalizacja mediów", "es": "Ubicación de medios"}, Options: []mediaSeedOption{
		{Key: "living_room", Labels: map[string]string{"en": "Living Room", "pl": "Salon", "es": "Sala de estar"}},
		{Key: "kitchen", Labels: map[string]string{"en": "Kitchen", "pl": "Kuchnia", "es": "Cocina"}},
		{Key: "bathroom", Labels: map[string]string{"en": "Bathroom", "pl": "Łazienka", "es": "Baño"}},
		{Key: "bedroom", Labels: map[string]string{"en": "Bedroom", "pl": "Sypialnia", "es": "Dormitorio"}},
		{Key: "dining_room", Labels: map[string]string{"en": "Dining Room", "pl": "Jadalnia", "es": "Comedor"}},
		{Key: "balcony", Labels: map[string]string{"en": "Balcony", "pl": "Balkon", "es": "Balcón"}},
		{Key: "terrace", Labels: map[string]string{"en": "Terrace", "pl": "Taras", "es": "Terraza"}},
		{Key: "garden", Labels: map[string]string{"en": "Garden", "pl": "Ogród", "es": "Jardín"}},
		{Key: "garage", Labels: map[string]string{"en": "Garage", "pl": "Garaż", "es": "Garaje"}},
		{Key: "basement", Labels: map[string]string{"en": "Basement", "pl": "Piwnica", "es": "Sótano"}},
		{Key: "attic", Labels: map[string]string{"en": "Attic", "pl": "Strych", "es": "Ático"}},
		{Key: "exterior", Labels: map[string]string{"en": "Exterior", "pl": "Zewnętrze", "es": "Exterior"}},
		{Key: "entrance", Labels: map[string]string{"en": "Entrance", "pl": "Wejście", "es": "Entrada"}},
		{Key: "aerial_view", Labels: map[string]string{"en": "Aerial View", "pl": "Widok z powietrza", "es": "Vista aérea"}},
		{Key: "pool", Labels: map[string]string{"en": "Pool", "pl": "Basen", "es": "Piscina"}},
		{Key: "sauna", Labels: map[string]string{"en": "Sauna", "pl": "Sauna", "es": "Sauna"}},
		{Key: "gym", Labels: map[string]string{"en": "Fitness Room", "pl": "Siłownia", "es": "Gimnasio"}},
		{Key: "lobby", Labels: map[string]string{"en": "Lobby", "pl": "Hol", "es": "Lobby"}},
		{Key: "rooftop", Labels: map[string]string{"en": "Rooftop", "pl": "Dach", "es": "Azotea"}},
		{Key: "courtyard", Labels: map[string]string{"en": "Courtyard", "pl": "Dziedziniec", "es": "Patio"}},
		{Key: "kids_room", Labels: map[string]string{"en": "Kids Room", "pl": "Pokój dziecięcy", "es": "Cuarto infantil"}},
		{Key: "home_office", Labels: map[string]string{"en": "Home Office", "pl": "Gabinet", "es": "Oficina"}},
		{Key: "laundry_room", Labels: map[string]string{"en": "Laundry Room", "pl": "Pralnia", "es": "Lavadero"}},
	}},
	{Name: "media_kind", Labels: map[string]string{"en": "Media Kind", "pl": "Rodzaj mediów", "es": "Tipo de medio"}, Options: []mediaSeedOption{
		{Key: "real", Labels: map[string]string{"en": "Real Photo", "pl": "Zdjęcie rzeczywiste", "es": "Foto real"}},
		{Key: "drawing", Labels: map[string]string{"en": "Technical Drawing", "pl": "Rysunek techniczny", "es": "Dibujo técnico"}},
		{Key: "sketch", Labels: map[string]string{"en": "Sketch", "pl": "Szkic", "es": "Boceto"}},
		{Key: "ai_generated", Labels: map[string]string{"en": "AI Generated", "pl": "Wygenerowane przez AI", "es": "Generado por IA"}},
		{Key: "render", Labels: map[string]string{"en": "3D Render", "pl": "Render 3D", "es": "Render 3D"}},
		{Key: "floor_plan", Labels: map[string]string{"en": "Floor Plan", "pl": "Plan mieszkania", "es": "Plano"}},
	}},
	{Name: "media_style", Labels: map[string]string{"en": "Style Tags", "pl": "Tagi stylu", "es": "Etiquetas de estilo"}, Options: []mediaSeedOption{
		{Key: "modern", Labels: map[string]string{"en": "Modern", "pl": "Nowoczesny", "es": "Moderno"}},
		{Key: "rustic", Labels: map[string]string{"en": "Rustic", "pl": "Rustykalny", "es": "Rústico"}},
		{Key: "classic", Labels: map[string]string{"en": "Classic", "pl": "Klasyczny", "es": "Clásico"}},
		{Key: "industrial", Labels: map[string]string{"en": "Industrial", "pl": "Przemysłowy", "es": "Industrial"}},
		{Key: "minimalist", Labels: map[string]string{"en": "Minimalist", "pl": "Minimalistyczny", "es": "Minimalista"}},
	}},
	{Name: "media_mood", Labels: map[string]string{"en": "Mood Tags", "pl": "Tagi nastroju", "es": "Etiquetas de ambiente"}, Options: []mediaSeedOption{
		{Key: "sunny", Labels: map[string]string{"en": "Sunny", "pl": "Słoneczne", "es": "Soleado"}},
		{Key: "cozy", Labels: map[string]string{"en": "Cozy", "pl": "Przytulny", "es": "Acogedor"}},
		{Key: "night", Labels: map[string]string{"en": "Night Shot", "pl": "Zdjęcie nocne", "es": "Foto nocturna"}},
		{Key: "staged", Labels: map[string]string{"en": "Staged", "pl": "Zainscenizowane", "es": "Ambientada"}},
		{Key: "virtual_staging", Labels: map[string]string{"en": "Virtual Staging", "pl": "Wirtualna aranżacja", "es": "Ambientación virtual"}},
	}},
	{Name: "media_capture_conditions", Labels: map[string]string{"en": "Capture Conditions", "es": "Condiciones de captura", "pl": "Warunki zdjęcia"}, Options: []mediaSeedOption{
		{Key: "daylight", Labels: map[string]string{"en": "Daylight", "es": "Luz diurna", "pl": "Światło dzienne"}},
		{Key: "twilight", Labels: map[string]string{"en": "Twilight", "es": "Crepúsculo", "pl": "Zmierzch"}},
		{Key: "night", Labels: map[string]string{"en": "Night", "es": "Noche", "pl": "Noc"}},
		{Key: "virtual_tour_still", Labels: map[string]string{"en": "Virtual Tour Still", "es": "Fotograma de tour virtual", "pl": "Kadr wirtualnego spaceru"}},
		{Key: "hdr", Labels: map[string]string{"en": "HDR", "es": "HDR", "pl": "HDR"}},
		{Key: "drone", Labels: map[string]string{"en": "Drone Shot", "es": "Toma con dron", "pl": "Zdjęcie z drona"}},
	}},
	{Name: "media_usage", Labels: map[string]string{"en": "Media Usage", "es": "Uso del medio", "pl": "Zastosowanie"}, Options: []mediaSeedOption{
		{Key: "hero_image", Labels: map[string]string{"en": "Hero Image", "es": "Imagen principal", "pl": "Zdjęcie główne"}},
		{Key: "gallery", Labels: map[string]string{"en": "Gallery", "es": "Galería", "pl": "Galeria"}},
		{Key: "thumbnail", Labels: map[string]string{"en": "Thumbnail", "es": "Miniatura", "pl": "Miniatura"}},
		{Key: "floor_plan", Labels: map[string]string{"en": "Floor Plan", "es": "Plano", "pl": "Plan"}},
		{Key: "virtual_staging", Labels: map[string]string{"en": "Virtual Staging", "es": "Ambientación virtual", "pl": "Wirtualna aranżacja"}},
		{Key: "before_after", Labels: map[string]string{"en": "Before & After", "es": "Antes y después", "pl": "Przed i po"}},
	}},
	{Name: "media_feature_tags", Labels: map[string]string{"en": "Feature Tags", "es": "Etiquetas de características", "pl": "Tagi funkcji"}, Options: []mediaSeedOption{
		{Key: "fireplace", Labels: map[string]string{"en": "Fireplace", "es": "Chimenea", "pl": "Kominek"}},
		{Key: "walk_in_closet", Labels: map[string]string{"en": "Walk-in Closet", "es": "Vestidor", "pl": "Garderoba"}},
		{Key: "panoramic_view", Labels: map[string]string{"en": "Panoramic View", "es": "Vista panorámica", "pl": "Widok panoramiczny"}},
		{Key: "smart_home", Labels: map[string]string{"en": "Smart Home", "es": "Smart home", "pl": "Smart home"}},
		{Key: "eco_friendly", Labels: map[string]string{"en": "Eco Friendly", "es": "Ecológico", "pl": "Ekologiczny"}},
		{Key: "new_build", Labels: map[string]string{"en": "New Construction", "es": "Obra nueva", "pl": "Nowa inwestycja"}},
	}},
}

func seedMediaDictionary(ctx context.Context, setsCollection, optionsCollection *mongo.Collection) error {
	for _, set := range mediaSeedSets {
		setIDs := make(map[string]string)
		for locale, label := range set.Labels {
			setID, err := ensureMediaSet(ctx, setsCollection, set.Name, locale, label)
			if err != nil {
				return err
			}
			setIDs[locale] = setID
		}

		for order, opt := range set.Options {
			for locale := range set.Labels {
				label := opt.Labels[locale]
				if label == "" {
					label = opt.Labels["en"]
				}
				if err := ensureMediaOption(ctx, optionsCollection, setIDs[locale], locale, opt.Key, label, order); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func ensureMediaSet(ctx context.Context, setsCollection *mongo.Collection, name, locale, label string) (string, error) {
	filter := bson.M{"name": name, "locale": locale}
	var existing struct {
		ID string `bson:"_id"`
	}
	err := setsCollection.FindOne(ctx, filter).Decode(&existing)
	if err == nil && existing.ID != "" {
		return existing.ID, nil
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return "", err
	}

	id := uuid.New().String()
	doc := bson.M{
		"_id":         id,
		"name":        name,
		"locale":      locale,
		"label":       label,
		"description": "",
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}
	if _, err := setsCollection.UpdateOne(ctx, filter, bson.M{"$setOnInsert": doc}, options.Update().SetUpsert(true)); err != nil {
		return "", err
	}
	return id, nil
}

func ensureMediaOption(ctx context.Context, optionsCollection *mongo.Collection, setID, locale, key, label string, order int) error {
	filter := bson.M{"set_id": setID, "key": key, "locale": locale}
	if err := optionsCollection.FindOne(ctx, filter).Err(); err == nil {
		return nil
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	id := uuid.New().String()
	doc := bson.M{
		"_id":         id,
		"set_id":      setID,
		"parent_id":   nil,
		"locale":      locale,
		"short_code":  shortCode(key),
		"key":         key,
		"label":       label,
		"description": "",
		"value":       label,
		"order":       order,
		"active":      true,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  "system",
		"updated_by":  "system",
	}
	_, err := optionsCollection.UpdateOne(ctx, filter, bson.M{"$setOnInsert": doc}, options.Update().SetUpsert(true))
	return err
}

func shortCode(key string) string {
	if len(key) <= 10 {
		return key
	}
	return key[:10]
}
