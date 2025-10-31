package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// SeedData represents the JSON structure from seed file.
type SeedData struct {
	Format  string       `json:"_format"`
	Sets    []SeedSet    `json:"sets"`
	Options []SeedOption `json:"options"`
}

// SeedSet represents a set definition in the seed file.
type SeedSet struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

// SeedOption represents an option definition in the seed file.
type SeedOption struct {
	Set       string            `json:"set"`
	Key       string            `json:"key"`
	ShortCode string            `json:"short_code"`
	Value     string            `json:"value"`
	Labels    map[string]string `json:"labels"`
	ParentKey *string           `json:"parent_key"`
	Locale    string            `json:"locale"`
	Active    bool              `json:"active"`
	Order     int               `json:"order"`
}

// Geographic sets to exclude from seeding.
var geoSetsToExclude = map[string]bool{
	"country":                 true,
	"pl_voivodeship":          true,
	"ar_province":             true,
	"es_autonomous_community": true,
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input.json> <output.go>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read and parse JSON
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var seedData SeedData
	if err := json.Unmarshal(data, &seedData); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Filter geographic sets
	filteredSets := make([]SeedSet, 0)
	for _, set := range seedData.Sets {
		if !geoSetsToExclude[set.Name] {
			filteredSets = append(filteredSets, set)
		}
	}

	filteredOptions := make([]SeedOption, 0)
	for _, opt := range seedData.Options {
		if !geoSetsToExclude[opt.Set] {
			filteredOptions = append(filteredOptions, opt)
		}
	}

	// Get unique locales
	localeMap := make(map[string]bool)
	for _, opt := range filteredOptions {
		localeMap[opt.Locale] = true
	}
	locales := make([]string, 0, len(localeMap))
	for locale := range localeMap {
		locales = append(locales, locale)
	}
	sort.Strings(locales)

	// Generate Go code
	code := generateGoCode(filteredSets, filteredOptions, locales)

	// Write to output file
	if err := os.WriteFile(outputFile, []byte(code), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s with %d sets and %d options (excluded geographic data)\n",
		outputFile, len(filteredSets), len(filteredOptions))
}

func generateGoCode(sets []SeedSet, options []SeedOption, locales []string) string {
	var b strings.Builder

	// Reset counter for each generation
	optionCounter = 0

	b.WriteString(`package dictionary

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
			ID:          "2025-10-30_real_estate_dictionary",
			Description: "Load real estate dictionary (excluding geographic data)",
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
`)

	// Generate sets code
	b.WriteString("\t// Create sets\n")
	for _, set := range sets {
		for _, locale := range locales {
			b.WriteString(fmt.Sprintf("\tsetID_%s_%s := uuid.New().String()\n",
				sanitizeName(set.Name), locale))
			b.WriteString(fmt.Sprintf("\tsetIDMap[\"%s:%s\"] = setID_%s_%s\n",
				set.Name, locale, sanitizeName(set.Name), locale))
			b.WriteString(fmt.Sprintf("\t_, _ = setsCollection.UpdateOne(ctx, bson.M{\"name\": %q, \"locale\": %q}, bson.M{\"$setOnInsert\": bson.M{\n",
				set.Name, locale))
			b.WriteString(fmt.Sprintf("\t\t\"_id\": setID_%s_%s,\n", sanitizeName(set.Name), locale))
			b.WriteString(fmt.Sprintf("\t\t\"name\": %q,\n", set.Name))
			b.WriteString(fmt.Sprintf("\t\t\"locale\": %q,\n", locale))
			b.WriteString(fmt.Sprintf("\t\t\"label\": %q,\n", set.Label))
			b.WriteString("\t\t\"description\": \"\",\n")
			b.WriteString(fmt.Sprintf("\t\t\"active\": %v,\n", set.Active))
			b.WriteString("\t\t\"created_at\": time.Now(),\n")
			b.WriteString("\t\t\"updated_at\": time.Now(),\n")
			b.WriteString("\t\t\"created_by\": \"system\",\n")
			b.WriteString("\t\t\"updated_by\": \"system\",\n")
			b.WriteString("\t}}, options.Update().SetUpsert(true))\n\n")
		}
	}

	// Generate options code
	b.WriteString("\t// Track option IDs for parent relationships\n")
	b.WriteString("\toptionIDMap := make(map[string]string)\n\n")

	// First pass: options without parents
	b.WriteString("\t// Create options without parents\n")
	for _, opt := range options {
		if opt.ParentKey != nil {
			continue
		}
		b.WriteString(generateOptionCode(opt))
	}

	// Second pass: options with parents
	b.WriteString("\t// Create options with parents\n")
	for _, opt := range options {
		if opt.ParentKey == nil {
			continue
		}
		b.WriteString(generateOptionCode(opt))
	}

	b.WriteString("\n\treturn nil\n}\n")

	return b.String()
}

var optionCounter int

func generateOptionCode(opt SeedOption) string {
	var b strings.Builder

	optionCounter++

	optKey := fmt.Sprintf("%s:%s:%s", opt.Set, opt.Key, opt.Locale)
	optIDVar := fmt.Sprintf("optID_%d", optionCounter)

	b.WriteString(fmt.Sprintf("\t%s := uuid.New().String()\n", optIDVar))
	b.WriteString(fmt.Sprintf("\toptionIDMap[%q] = %s\n", optKey, optIDVar))

	// Get parent ID if needed
	parentIDExpr := "nil"
	if opt.ParentKey != nil {
		parentSet := getParentSetName(opt.Set)
		parentKey := fmt.Sprintf("%s:%s:%s", parentSet, *opt.ParentKey, opt.Locale)
		parentIDVar := fmt.Sprintf("parentID_%d", optionCounter)
		b.WriteString(fmt.Sprintf("\tif %s, ok := optionIDMap[%q]; ok {\n", parentIDVar, parentKey))
		b.WriteString(fmt.Sprintf("\t\tparentID_%d_ptr := %s\n", optionCounter, parentIDVar))
		parentIDExpr = fmt.Sprintf("&parentID_%d_ptr", optionCounter)
		// Close the if later after using the variable
	}

	// Get label for locale
	label := opt.Labels[opt.Locale]
	if label == "" {
		label = opt.Value
	}

	b.WriteString(fmt.Sprintf("\t_, _ = optionsCollection.UpdateOne(ctx, bson.M{\"set_id\": setIDMap[%q], \"key\": %q, \"locale\": %q}, bson.M{\"$setOnInsert\": bson.M{\n",
		fmt.Sprintf("%s:%s", opt.Set, opt.Locale), opt.Key, opt.Locale))
	b.WriteString(fmt.Sprintf("\t\t\"_id\": %s,\n", optIDVar))
	b.WriteString(fmt.Sprintf("\t\t\"set_id\": setIDMap[%q],\n", fmt.Sprintf("%s:%s", opt.Set, opt.Locale)))
	b.WriteString(fmt.Sprintf("\t\t\"parent_id\": %s,\n", parentIDExpr))
	b.WriteString(fmt.Sprintf("\t\t\"locale\": %q,\n", opt.Locale))
	b.WriteString(fmt.Sprintf("\t\t\"short_code\": %q,\n", escapeString(opt.ShortCode)))
	b.WriteString(fmt.Sprintf("\t\t\"key\": %q,\n", escapeString(opt.Key)))
	b.WriteString(fmt.Sprintf("\t\t\"label\": %q,\n", escapeString(label)))
	b.WriteString("\t\t\"description\": \"\",\n")
	b.WriteString(fmt.Sprintf("\t\t\"value\": %q,\n", escapeString(opt.Value)))
	b.WriteString(fmt.Sprintf("\t\t\"order\": %d,\n", opt.Order))
	b.WriteString(fmt.Sprintf("\t\t\"active\": %v,\n", opt.Active))
	b.WriteString("\t\t\"created_at\": time.Now(),\n")
	b.WriteString("\t\t\"updated_at\": time.Now(),\n")
	b.WriteString("\t\t\"created_by\": \"system\",\n")
	b.WriteString("\t\t\"updated_by\": \"system\",\n")
	b.WriteString("\t}}, options.Update().SetUpsert(true))\n")

	if opt.ParentKey != nil {
		b.WriteString("\t}\n")
	}
	b.WriteString("\n")

	return b.String()
}

func getParentSetName(setName string) string {
	switch setName {
	case "estate_type":
		return "estate_category"
	case "estate_subtype":
		return "estate_type"
	default:
		return setName
	}
}

func sanitizeName(name string) string {
	s := strings.ReplaceAll(name, "_", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func escapeString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\\", "\\\\"), "\"", "\\\"")
}
