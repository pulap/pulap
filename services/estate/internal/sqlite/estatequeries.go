package sqlite

const (
	// Queries for Estate aggregate root operations

	// QueryCreateEstateRoot creates a new Estate aggregate root record.
	QueryCreateEstateRoot = `INSERT INTO estates (id, description, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	// QueryGetEstateRoot retrieves a Estate aggregate root record by ID.
	QueryGetEstateRoot = `SELECT id, description, name, created_at, updated_at FROM estates WHERE id = ?`

	// QueryUpdateEstateRoot updates an existing Estate aggregate root record.
	QueryUpdateEstateRoot = `UPDATE estates SET description = ?, name = ?, updated_at = ? WHERE id = ?`

	// QueryDeleteEstateRoot deletes a Estate aggregate root record by ID.
	QueryDeleteEstateRoot = `DELETE FROM estates WHERE id = ?`

	// QueryEstateEstateRoot estates all Estate aggregate root records.
	QueryEstateEstateRoot = `SELECT id FROM estates ORDER BY created_at DESC`

	// Queries for Estate's Items child entities

	// QueryCreateEstateItems creates new Item records for a Estate aggregate.
	QueryCreateEstateItems = `INSERT INTO items (text, done) VALUES `

	// QueryGetEstateItems retrieves all Item records for a specific Estate aggregate.
	QueryGetEstateItems = `SELECT text, done FROM items WHERE Estate_id = ? ORDER BY created_at`

	// QueryUpdateEstateItem updates an existing Item record within a Estate aggregate.
	QueryUpdateEstateItem = `UPDATE items SET text = ?, done = ? WHERE id = ?`

	// QueryDeleteEstateItems deletes all Item records for a specific Estate aggregate.
	QueryDeleteEstateItems = `DELETE FROM items WHERE Estate_id = ?`

	// QueryDeleteEstateItemsByIDs deletes specific Item records by their IDs.
	QueryDeleteEstateItemsByIDs = `DELETE FROM items WHERE id IN `

	// Helper query parts for batch operations
	ItemValuePlaceholder = `(?, ?)`

	// Queries for Estate's Tags child entities

	// QueryCreateEstateTags creates new Tag records for a Estate aggregate.
	QueryCreateEstateTags = `INSERT INTO tags (name, color) VALUES `

	// QueryGetEstateTags retrieves all Tag records for a specific Estate aggregate.
	QueryGetEstateTags = `SELECT name, color FROM tags WHERE Estate_id = ? ORDER BY created_at`

	// QueryUpdateEstateTag updates an existing Tag record within a Estate aggregate.
	QueryUpdateEstateTag = `UPDATE tags SET name = ?, color = ? WHERE id = ?`

	// QueryDeleteEstateTags deletes all Tag records for a specific Estate aggregate.
	QueryDeleteEstateTags = `DELETE FROM tags WHERE Estate_id = ?`

	// QueryDeleteEstateTagsByIDs deletes specific Tag records by their IDs.
	QueryDeleteEstateTagsByIDs = `DELETE FROM tags WHERE id IN `

	// Helper query parts for batch operations
	TagValuePlaceholder = `(?, ?)`
)
