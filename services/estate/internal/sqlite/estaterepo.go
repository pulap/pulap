package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"github.com/pulap/pulap/pkg/lib/core"
	"github.com/pulap/pulap/services/estate/internal/config"
	"github.com/pulap/pulap/services/estate/internal/estate"
)

// EstateSQLiteRepo implements the EstateRepo interface using SQLite.
// SQLite requires more complex logic to handle aggregates across multiple related tables.
type EstateSQLiteRepo struct {
	db      *sql.DB
	xparams config.XParams
}

// NewEstateSQLiteRepo creates a new SQLite repository for Estate aggregates.
func NewEstateSQLiteRepo(xparams config.XParams) *EstateSQLiteRepo {
	return &EstateSQLiteRepo{
		xparams: xparams,
	}
}

// Start opens the database connection and pings it.
func (r *EstateSQLiteRepo) Start(ctx context.Context) error {
	appCfg := r.xparams.Cfg()

	dbPath := appCfg.Database.Path

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbPath))
	if err != nil {
		return fmt.Errorf("cannot open database: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("cannot connect to database: %w", err)
	}
	r.db = db
	// TODO: Run migrations here
	return nil
}

// Stop closes the database connection.
func (r *EstateSQLiteRepo) Stop(ctx context.Context) error {
	if r.db != nil {
		if err := r.db.Close(); err != nil {
			return fmt.Errorf("cannot close database: %w", err)
		}
	}
	return nil
}

// Create creates a new Estate aggregate in SQLite.
// This involves inserting the root and all child entities in a single transaction.
func (r *EstateSQLiteRepo) Create(ctx context.Context, aggregate *estate.Estate) error {
	if aggregate == nil {
		return fmt.Errorf("aggregate cannot be nil")
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	aggregate.EnsureID()
	aggregate.BeforeCreate()

	if err := r.insertRoot(ctx, tx, aggregate); err != nil {
		return fmt.Errorf("could not insert aggregate root: %w", err)
	}

	if err := r.insertItems(ctx, tx, aggregate.GetID(), aggregate.Items); err != nil {
		return fmt.Errorf("could not insert Items: %w", err)
	}

	if err := r.insertTags(ctx, tx, aggregate.GetID(), aggregate.Tags); err != nil {
		return fmt.Errorf("could not insert Tags: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

// Get retrieves a complete Estate aggregate by ID from SQLite.
// This involves loading the root and all child entities from multiple tables.
func (r *EstateSQLiteRepo) Get(ctx context.Context, id uuid.UUID) (*estate.Estate, error) {
	aggregate, err := r.getRoot(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get aggregate root: %w", err)
	}

	Items, err := r.getItems(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get Items: %w", err)
	}
	aggregate.Items = Items

	Tags, err := r.getTags(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get Tags: %w", err)
	}
	aggregate.Tags = Tags

	return aggregate, nil
}

// Save performs a unit-of-work save operation on the Estate aggregate.
// This computes diffs and updates/inserts/deletes child entities as needed.
func (r *EstateSQLiteRepo) Save(ctx context.Context, aggregate *estate.Estate) error {
	if aggregate == nil {
		return fmt.Errorf("aggregate cannot be nil")
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	aggregate.BeforeUpdate()

	if err := r.updateRoot(ctx, tx, aggregate); err != nil {
		return fmt.Errorf("could not update aggregate root: %w", err)
	}

	if err := r.saveItems(ctx, tx, aggregate.GetID(), aggregate.Items); err != nil {
		return fmt.Errorf("could not save Items: %w", err)
	}

	if err := r.saveTags(ctx, tx, aggregate.GetID(), aggregate.Tags); err != nil {
		return fmt.Errorf("could not save Tags: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

// Delete removes the entire Estate aggregate from SQLite.
// This cascades to all child entities.
func (r *EstateSQLiteRepo) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	if err := r.deleteItems(ctx, tx, id); err != nil {
		return fmt.Errorf("could not delete Items: %w", err)
	}

	if err := r.deleteTags(ctx, tx, id); err != nil {
		return fmt.Errorf("could not delete Tags: %w", err)
	}

	result, err := tx.ExecContext(ctx, QueryDeleteEstateRoot, id.String())
	if err != nil {
		return fmt.Errorf("could not delete aggregate root: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Estate aggregate with ID %s not found for deletion", id.String())
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

// Estate retrieves all Estate aggregates from SQLite.
// This loads each aggregate with all its child entities.
func (r *EstateSQLiteRepo) Estate(ctx context.Context) ([]*estate.Estate, error) {
	rows, err := r.db.QueryContext(ctx, QueryEstateEstateRoot)
	if err != nil {
		return nil, fmt.Errorf("could not query aggregate IDs: %w", err)
	}
	defer rows.Close()

	var aggregates []*estate.Estate

	for rows.Next() {
		var idStr string
		if err := rows.Scan(&idStr); err != nil {
			return nil, fmt.Errorf("could not scan aggregate ID: %w", err)
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("could not parse UUID %s: %w", idStr, err)
		}

		aggregate, err := r.Get(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("could not get aggregate %s: %w", idStr, err)
		}

		aggregates = append(aggregates, aggregate)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return aggregates, nil
}

// Helper methods for aggregate root operations

func (r *EstateSQLiteRepo) insertRoot(ctx context.Context, tx *sql.Tx, aggregate *estate.Estate) error {
	_, err := tx.ExecContext(ctx, QueryCreateEstateRoot, aggregate.GetID().String(), aggregate.Description, aggregate.Name, aggregate.CreatedAt, aggregate.UpdatedAt)
	return err
}

func (r *EstateSQLiteRepo) getRoot(ctx context.Context, id uuid.UUID) (*estate.Estate, error) {
	var aggregate estate.Estate
	var idStr string

	err := r.db.QueryRowContext(ctx, QueryGetEstateRoot, id.String()).Scan(
		&idStr, &aggregate.Description, &aggregate.Name, &aggregate.CreatedAt, &aggregate.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Estate aggregate with ID %s not found", id.String())
		}
		return nil, fmt.Errorf("could not scan aggregate root: %w", err)
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse UUID %s: %w", idStr, err)
	}
	aggregate.SetID(parsedID)

	return &aggregate, nil
}

func (r *EstateSQLiteRepo) updateRoot(ctx context.Context, tx *sql.Tx, aggregate *estate.Estate) error {
	result, err := tx.ExecContext(ctx, QueryUpdateEstateRoot, aggregate.Description, aggregate.Name, aggregate.UpdatedAt, aggregate.GetID().String())
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Estate aggregate with ID %s not found for update", aggregate.GetID().String())
	}

	return nil
}

// Helper methods for Items child entities

func (r *EstateSQLiteRepo) insertItems(ctx context.Context, tx *sql.Tx, rootID uuid.UUID, items []estate.Item) error {
	if len(items) == 0 {
		return nil
	}

	query := `INSERT INTO items (id, Estate_id, text, done, created_at, updated_at) VALUES {{.Placeholders}}`

	var args []interface{}
	var placeholders []string

	for _, item := range items {
		item.EnsureID()
		item.BeforeCreate()

		placeholders = append(placeholders, "(?, ?, ?, ?, ?, ?)")
		args = append(args, item.GetID().String(), rootID.String(), item.Text, item.Done, item.CreatedAt, item.UpdatedAt)
	}

	finalQuery := strings.Replace(query, "{{.Placeholders}}", strings.Join(placeholders, ", "), 1)
	_, err := tx.ExecContext(ctx, finalQuery, args...)
	return err
}

func (r *EstateSQLiteRepo) getItems(ctx context.Context, rootID uuid.UUID) ([]estate.Item, error) {
	return r.getItemsWithTx(ctx, nil, rootID)
}

func (r *EstateSQLiteRepo) getItemsWithTx(ctx context.Context, tx *sql.Tx, rootID uuid.UUID) ([]estate.Item, error) {
	query := `SELECT id, text, done, created_at, updated_at FROM items WHERE Estate_id = ? ORDER BY created_at`

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, rootID.String())
	} else {
		rows, err = r.db.QueryContext(ctx, query, rootID.String())
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []estate.Item

	for rows.Next() {
		var item estate.Item
		var idStr string

		err := rows.Scan(&idStr, &item.Text, &item.Done, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("error parse UUID %s: %w", idStr, err)
		}
		item.SetID(id)

		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *EstateSQLiteRepo) saveItems(ctx context.Context, tx *sql.Tx, rootID uuid.UUID, newItems []estate.Item) error {
	// Get current items from database using the transaction
	currentItems, err := r.getItemsWithTx(ctx, tx, rootID)
	if err != nil {
		return fmt.Errorf("could not get current items: %w", err)
	}

	// Compute diff
	toInsert, toUpdate, toDelete := r.computeItemDiff(currentItems, newItems)

	// Apply changes
	if len(toDelete) > 0 {
		if err := r.deleteItemsByIDs(ctx, tx, toDelete); err != nil {
			return fmt.Errorf("error delete items: %w", err)
		}
	}

	if len(toInsert) > 0 {
		if err := r.insertItems(ctx, tx, rootID, toInsert); err != nil {
			return fmt.Errorf("error insert items: %w", err)
		}
	}

	if len(toUpdate) > 0 {
		if err := r.updateItems(ctx, tx, toUpdate); err != nil {
			return fmt.Errorf("error update items: %w", err)
		}
	}

	return nil
}

func (r *EstateSQLiteRepo) deleteItems(ctx context.Context, tx *sql.Tx, rootID uuid.UUID) error {
	_, err := tx.ExecContext(ctx, `DELETE FROM items WHERE Estate_id = ?`, rootID.String())
	return err
}

func (r *EstateSQLiteRepo) deleteItemsByIDs(ctx context.Context, tx *sql.Tx, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id.String()
	}

	query := fmt.Sprintf(`DELETE FROM items WHERE id IN (%s)`, strings.Join(placeholders, ", "))
	_, err := tx.ExecContext(ctx, query, args...)
	return err
}

func (r *EstateSQLiteRepo) updateItems(ctx context.Context, tx *sql.Tx, items []estate.Item) error {
	for _, item := range items {
		item.BeforeUpdate()

		query := `UPDATE items SET text = ?, done = ?, updated_at = ? WHERE id = ?`
		_, err := tx.ExecContext(ctx, query, item.Text, item.Done, item.UpdatedAt, item.GetID().String())
		if err != nil {
			return fmt.Errorf("error update item %s: %w", item.GetID().String(), err)
		}
	}
	return nil
}

// computeItemDiff computes the difference between current and new items
func (r *EstateSQLiteRepo) computeItemDiff(current, new []estate.Item) (toInsert, toUpdate []estate.Item, toDelete []uuid.UUID) {
	// Create maps for efficient lookup
	currentMap := make(map[string]estate.Item)
	newMap := make(map[string]estate.Item)

	for _, item := range current {
		currentMap[item.GetID().String()] = item
	}

	for _, item := range new {
		if item.GetID() == uuid.Nil {
			// New item without ID - needs insert
			toInsert = append(toInsert, item)
		} else {
			newMap[item.GetID().String()] = item
			if _, exists := currentMap[item.GetID().String()]; exists {
				// Item exists - needs update
				toUpdate = append(toUpdate, item)
			} else {
				// Item with ID but not in current - needs insert
				toInsert = append(toInsert, item)
			}
		}
	}

	// Find items to delete (in current but not in new)
	for id := range currentMap {
		if _, exists := newMap[id]; !exists {
			uid, _ := uuid.Parse(id)
			toDelete = append(toDelete, uid)
		}
	}

	return toInsert, toUpdate, toDelete
}

// Helper methods for Tags child entities

func (r *EstateSQLiteRepo) insertTags(ctx context.Context, tx *sql.Tx, rootID uuid.UUID, items []estate.Tag) error {
	if len(items) == 0 {
		return nil
	}

	query := `INSERT INTO tags (id, Estate_id, name, color, created_at, updated_at) VALUES {{.Placeholders}}`

	var args []interface{}
	var placeholders []string

	for _, item := range items {
		item.EnsureID()
		item.BeforeCreate()

		placeholders = append(placeholders, "(?, ?, ?, ?, ?, ?)")
		args = append(args, item.GetID().String(), rootID.String(), item.Name, item.Color, item.CreatedAt, item.UpdatedAt)
	}

	finalQuery := strings.Replace(query, "{{.Placeholders}}", strings.Join(placeholders, ", "), 1)
	_, err := tx.ExecContext(ctx, finalQuery, args...)
	return err
}

func (r *EstateSQLiteRepo) getTags(ctx context.Context, rootID uuid.UUID) ([]estate.Tag, error) {
	return r.getTagsWithTx(ctx, nil, rootID)
}

func (r *EstateSQLiteRepo) getTagsWithTx(ctx context.Context, tx *sql.Tx, rootID uuid.UUID) ([]estate.Tag, error) {
	query := `SELECT id, name, color, created_at, updated_at FROM tags WHERE Estate_id = ? ORDER BY created_at`

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, rootID.String())
	} else {
		rows, err = r.db.QueryContext(ctx, query, rootID.String())
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []estate.Tag

	for rows.Next() {
		var item estate.Tag
		var idStr string

		err := rows.Scan(&idStr, &item.Name, &item.Color, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("error parse UUID %s: %w", idStr, err)
		}
		item.SetID(id)

		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *EstateSQLiteRepo) saveTags(ctx context.Context, tx *sql.Tx, rootID uuid.UUID, newItems []estate.Tag) error {
	// Get current items from database using the transaction
	currentItems, err := r.getTagsWithTx(ctx, tx, rootID)
	if err != nil {
		return fmt.Errorf("could not get current items: %w", err)
	}

	// Compute diff
	toInsert, toUpdate, toDelete := r.computeTagDiff(currentItems, newItems)

	// Apply changes
	if len(toDelete) > 0 {
		if err := r.deleteTagsByIDs(ctx, tx, toDelete); err != nil {
			return fmt.Errorf("error delete items: %w", err)
		}
	}

	if len(toInsert) > 0 {
		if err := r.insertTags(ctx, tx, rootID, toInsert); err != nil {
			return fmt.Errorf("error insert items: %w", err)
		}
	}

	if len(toUpdate) > 0 {
		if err := r.updateTags(ctx, tx, toUpdate); err != nil {
			return fmt.Errorf("error update items: %w", err)
		}
	}

	return nil
}

func (r *EstateSQLiteRepo) deleteTags(ctx context.Context, tx *sql.Tx, rootID uuid.UUID) error {
	_, err := tx.ExecContext(ctx, `DELETE FROM tags WHERE Estate_id = ?`, rootID.String())
	return err
}

func (r *EstateSQLiteRepo) deleteTagsByIDs(ctx context.Context, tx *sql.Tx, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id.String()
	}

	query := fmt.Sprintf(`DELETE FROM tags WHERE id IN (%s)`, strings.Join(placeholders, ", "))
	_, err := tx.ExecContext(ctx, query, args...)
	return err
}

func (r *EstateSQLiteRepo) updateTags(ctx context.Context, tx *sql.Tx, items []estate.Tag) error {
	for _, item := range items {
		item.BeforeUpdate()

		query := `UPDATE tags SET name = ?, color = ?, updated_at = ? WHERE id = ?`
		_, err := tx.ExecContext(ctx, query, item.Name, item.Color, item.UpdatedAt, item.GetID().String())
		if err != nil {
			return fmt.Errorf("error update item %s: %w", item.GetID().String(), err)
		}
	}
	return nil
}

// computeTagDiff computes the difference between current and new items
func (r *EstateSQLiteRepo) computeTagDiff(current, new []estate.Tag) (toInsert, toUpdate []estate.Tag, toDelete []uuid.UUID) {
	// Create maps for efficient lookup
	currentMap := make(map[string]estate.Tag)
	newMap := make(map[string]estate.Tag)

	for _, item := range current {
		currentMap[item.GetID().String()] = item
	}

	for _, item := range new {
		if item.GetID() == uuid.Nil {
			// New item without ID - needs insert
			toInsert = append(toInsert, item)
		} else {
			newMap[item.GetID().String()] = item
			if _, exists := currentMap[item.GetID().String()]; exists {
				// Item exists - needs update
				toUpdate = append(toUpdate, item)
			} else {
				// Item with ID but not in current - needs insert
				toInsert = append(toInsert, item)
			}
		}
	}

	// Find items to delete (in current but not in new)
	for id := range currentMap {
		if _, exists := newMap[id]; !exists {
			uid, _ := uuid.Parse(id)
			toDelete = append(toDelete, uid)
		}
	}

	return toInsert, toUpdate, toDelete
}

func (r *EstateSQLiteRepo) log() core.Logger {
	return r.xparams.Log()
}

func (r *EstateSQLiteRepo) cfg() *config.Config {
	return r.xparams.Cfg()
}

func (r *EstateSQLiteRepo) trace() core.Tracer {
	return r.xparams.Tracer()
}
