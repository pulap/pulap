# Dictionary Service Architecture

**Status:** Draft | **Updated:** 2025-11-01 | **Version:** 0.1

## Overview

This document describes the **Dictionary Service** architecture for Pulap.
The service manages **Sets** and **Options** that define application-wide taxonomies and reference data.
It provides a flexible way to represent domain-specific enumerations (such as property types, amenities, payment methods, or locations) in a consistent, centrally managed structure.

## 1. Core Concepts

### 1.1 Set

A **Set** represents a logical group or category containing related options.

**Examples:**

* Property Types
* Amenities
* Payment Methods
* Locations

**Storage:** `sets` collection in MongoDB

**Purpose:** Sets provide a high-level grouping for options, allowing the system to maintain structured taxonomies across modules.

### 1.2 Option

An **Option** represents an individual value or selectable item within a specific Set.

**Examples:**

* In the "Property Types" set → House, Apartment, Condo
* In the "Amenities" set → Pool, Garage, Garden

**Storage:** `options` collection in MongoDB

**Purpose:** Options define the concrete selectable values that belong to a Set. They can be hierarchical, meaning each option can optionally reference a parent option.

## 2. Relationship Model

| Relationship   | Description                                                  |
| -------------- | ------------------------------------------------------------ |
| **Type**       | One-to-Many (One Set → Many Options)                         |
| **Dependency** | Options depend on Sets (a Set must exist before its Options) |
| **Storage**    | Separate MongoDB collections (`sets` and `options`)          |
| **Uniqueness** | Option keys are unique within their Set (not globally)       |

## 3. Data Model and Storage

### 3.1 Set Document (MongoDB)

```json
{
  "_id": "abc-123-def",
  "name": "property_types",
  "label": "Property Types",
  "description": "Main property classification",
  "active": true,
  "created_at": "2025-11-01T12:00:00Z",
  "updated_at": "2025-11-01T12:00:00Z"
}
```

### 3.2 Option Document (MongoDB)

```json
{
  "_id": "opt-001",
  "set_id": "abc-123-def",
  "parent_id": null,
  "short_code": "HSE",
  "key": "house",
  "label": "House",
  "value": "house",
  "order": 1,
  "active": true,
  "created_at": "2025-11-01T12:10:00Z",
  "updated_at": "2025-11-01T12:10:00Z"
}
```

### 3.3 Hierarchical Options Example

Options can form parent-child relationships within the same Set:

```
Set: "Locations"
├── USA (parent_id: null)
│   ├── California (parent_id: USA_ID)
│   └── Texas (parent_id: USA_ID)
└── Canada (parent_id: null)
    └── Ontario (parent_id: Canada_ID)
```

## 4. API Workflow

### 4.1 Create a Set

```bash
POST /sets
Content-Type: application/json

{
  "name": "property_types",
  "label": "Property Types",
  "description": "Main property classification",
  "active": true
}
```

**Response:**

```json
{ "id": "abc-123-def" }
```

### 4.2 Create Options for a Set

```bash
POST /options
Content-Type: application/json

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

## 5. Validation and Rules

1. **Set must exist before creating Options**
   Options reference an existing Set through `set_id`.

2. **Each Option must reference a Set**
   `set_id` is required for all Option records.

3. **Option keys must be unique within their Set**
   Ensures consistency without requiring global uniqueness.

4. **Options can have a parent_id**
   Enables hierarchical relationships within the same Set.

5. **Sets and Options are stored separately**
   They are independent collections and not embedded structures.

## Conclusion

The Dictionary Service provides a centralized mechanism to manage static or semi-static reference data.
By separating **Sets** (categories) from **Options** (values), it ensures data consistency, reusability, and extensibility across the Pulap platform.
