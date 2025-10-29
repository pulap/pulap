# Dictionary Service

## Overview

The Dictionary service manages Sets and Options for application-wide taxonomies and reference data.

## Core Concepts

### Set
A category or group that contains related options.

**Examples:**
- Property Types
- Amenities
- Payment Methods
- Locations

**Storage:** `sets` collection in MongoDB

### Option
Individual values/choices within a Set.

**Examples:**
- In "Property Types" set → House, Apartment, Condo
- In "Amenities" set → Pool, Garage, Garden

**Storage:** `options` collection in MongoDB

## Relationship

- **Type:** One-to-Many (One Set → Many Options)
- **Dependency:** Options depend on Sets (Set must exist first)
- **Storage:** Separate MongoDB collections
- **Uniqueness:** Option keys are unique within their Set (not globally)

## Workflow

### 1. Create a Set
```http
POST /sets
{
  "name": "property_types",
  "label": "Property Types",
  "description": "Main property classification",
  "active": true
}
```

Returns Set ID: `abc-123-def`

### 2. Create Options for that Set
```http
POST /options
{
  "set_id": "abc-123-def",
  "parent_id": null,
  "short_code": "HSE",
  "key": "house",
  "label": "House",
  "value": "house",
  "order": 1,
  "active": true
}
```

## Hierarchical Options

Options can have parent-child relationships within the same Set:

```
Set: "Locations"
├── USA (parent_id: null)
│   ├── California (parent_id: USA_ID)
│   └── Texas (parent_id: USA_ID)
└── Canada (parent_id: null)
    └── Ontario (parent_id: Canada_ID)
```

## Key Rules

1. **Set must exist before creating Options**
2. **Each Option must reference a Set** (set_id is required)
3. **Option keys must be unique within their Set**
4. **Options can optionally have a parent_id** (for hierarchies)
5. **Sets and Options are stored separately** (not embedded)
