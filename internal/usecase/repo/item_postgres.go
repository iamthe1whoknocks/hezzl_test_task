// methods for postgres.
package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/Masterminds/squirrel"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/postgres"
	"go.uber.org/zap"
)

// items repo struct.
type ItemsRepo struct {
	*postgres.Postgres
	Logger *zap.Logger
}

func New(pg *postgres.Postgres, logger *zap.Logger) ItemsRepo {
	return ItemsRepo{pg, logger}
}

func (r *ItemsRepo) GetItems(ctx context.Context) ([]models.Item, error) {
	sql, _, err := r.Builder.
		Select("id,campaign_id,name,description,priority,removed,created_at").
		From("items").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - GetItems - r.Builder : %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - GetItems - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	items := make([]models.Item, 0, 0) //nolint:gosimple // avoid redundant alloc

	for rows.Next() {
		i := models.Item{} //nolint:exhaustruct // struct to scan

		err = rows.Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("ItemsRepo - GetItems - rows.Scan: %w", err)
		}

		items = append(items, i)
	}
	return items, nil
}

func (r *ItemsRepo) SaveItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	creationTime := time.Now()

	sql, args, err := r.Builder.Insert("items").
		Columns("campaign_id,name,description,removed,created_at").
		Values(item.CampainID, item.Name, fmt.Sprintf("description of %s", item.Name), false, creationTime).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Builder insert: %w", err)
	}
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Pool.Begin: %w", err)
	}

	defer func() {
		err = tx.Rollback(ctx)
	}()
	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - tx.Exec insert: %w", err)
	}

	var i models.Item

	sql, args, err = r.Builder.
		Select("id,campaign_id,name,description,priority,removed,created_at").
		From("items").
		Where(squirrel.Eq{"created_at": creationTime, "campaign_id": item.CampainID, "name": item.Name}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Builder select: %w", err)
	}

	// r.Logger.Debug("ItemsRepo - SaveItem - tx.QueryRow - sql", zap.String("sql", sql), zap.Any("args", args))

	err = tx.QueryRow(ctx, sql, args...).Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt) //nolint:lll /// scan
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - tx.QueryRow - Scan: %w", err)
	}
	err = tx.Commit(ctx)
	return &i, err
}

func (r *ItemsRepo) DeleteItem(ctx context.Context, id, campaignID int) (bool, error) {
	sql, args, err := r.Builder.
		Select("id,campaign_id,name,description,priority,removed,created_at").
		From("items").
		Where(squirrel.Eq{"id": id, "campaign_id": campaignID}).Suffix("FOR UPDATE").
		ToSql()
	if err != nil {
		return false, fmt.Errorf("ItemsRepo - DeleteItem - r.Builder select: %w", err)
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return false, fmt.Errorf("ItemsRepo - DeleteItem - r.Pool.Begin: %w", err)
	}

	defer func() {
		err = tx.Rollback(ctx)
	}()
	var i models.Item

	err = tx.QueryRow(ctx, sql, args...).
		Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, err
	} else if err != nil {
		return false, fmt.Errorf("ItemsRepo - DeleteItem - tx.QueryRow: %w", err)
	}

	sql, args, err = r.Builder.Update("items").
		Set("removed", true).
		Where(squirrel.Eq{"id": id, "campaign_id": campaignID}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("ItemsRepo - SaveItem - r.Builder update: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("ItemsRepo - SaveItem - tx.Exec: %w", err)
	}
	err = tx.Commit(ctx)

	return true, err
}

func (r *ItemsRepo) UpdateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	sql, args, err := r.Builder.
		Select("id,campaign_id,name,description,priority,removed,created_at").
		From("items").
		Where(squirrel.Eq{"id": item.ID, "campaign_id": item.CampainID}).Suffix("FOR UPDATE").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - UpdateItem - r.Builder select: %w", err)
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - UpdateItem - r.Pool.Begin: %w", err)
	}

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			err = fmt.Errorf("ItemsRepo - UpdateItem - tx.Rollback: %w", err)
		}
	}()

	var i models.Item

	err = tx.QueryRow(ctx, sql, args...).
		Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("ItemsRepo - UpdateItem - tx.QueryRow: %w", err)
	}

	sql, args, err = r.Builder.Update("items").
		Set("name", item.Name).
		Set("description", item.Description).
		Where(squirrel.Eq{"id": item.ID, "campaign_id": item.CampainID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - UpdateItem - r.Builder update: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - UpdateItem - tx.Exec: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - UpdateItem - tx.Commit: %w", err)
	}

	i.Name = item.Name
	i.Description = item.Description

	return &i, nil
}
