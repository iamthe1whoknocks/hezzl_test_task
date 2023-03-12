package repo

import (
	"context"
	sq "database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/pkg/postgres"
)

type ItemsRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) ItemsRepo {
	return ItemsRepo{pg}
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

	items := make([]models.Item, 0, 0)

	for rows.Next() {
		i := models.Item{}

		err = rows.Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("ItemsRepo - GetItems - rows.Scan: %w", err)
		}

		items = append(items, i)
	}
	return items, nil
}

func (r *ItemsRepo) SaveItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	sql, _, err := r.Builder.Insert("items").
		Columns("campaign_id,name,description,removed,created_at").
		Values(item.CampainID, item.Name, fmt.Sprintf("description of %s", item.Name), false, time.Now()).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Builder insert: %w", err)
	}
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Pool.Begin: %w", err)
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - tx.Exec insert: %w", err)
	}

	var i models.Item

	sql, _, err = r.Builder.
		Select("id,campaign_id,name,description,priority,removed,created_at").
		From("items").
		Where(squirrel.Eq{"id": item.ID, "campaign_id": item.CampainID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Builder select: %w", err)
	}

	err = tx.QueryRow(ctx, sql).Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - tx.QueryRow: %w", err)
	}
	tx.Commit(ctx)
	return &i, nil
}

func (r *ItemsRepo) DeleteItem(ctx context.Context, id, campaignID int) (bool, error) {
	sql, _, err := r.Builder.
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

	defer tx.Rollback(ctx)

	var i models.Item

	err = tx.QueryRow(ctx, sql).
		Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)

	if err == sq.ErrNoRows {
		return false, err
	} else if err != nil {
		return false, fmt.Errorf("ItemsRepo - DeleteItem - tx.QueryRow: %w", err)
	}

	sql, _, err = r.Builder.Update("items").
		Set("removed", false).
		Where(squirrel.Eq{"id": id, "campaign_id": campaignID}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("ItemsRepo - SaveItem - r.Builder update: %w", err)
	}

	_, err = tx.Exec(ctx, sql)
	if err != nil {
		return false, fmt.Errorf("ItemsRepo - SaveItem - tx.Exec: %w", err)
	}
	tx.Commit(ctx)

	return true, nil

}

func (r *ItemsRepo) UpdateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	sql, _, err := r.Builder.
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

	defer tx.Rollback(ctx)

	var i models.Item

	err = tx.QueryRow(ctx, sql).
		Scan(&i.ID, &i.CampainID, &i.Name, &i.Description, &i.Priority, &i.Removed, &i.CreatedAt)

	if err == sq.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("ItemsRepo - DeleteItem - tx.QueryRow: %w", err)
	}

	sql, _, err = r.Builder.Update("items").
		Set("name", item.Name).
		Set("description", item.Description).
		Where(squirrel.Eq{"id": item.ID, "campaign_id": item.CampainID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - r.Builder update: %w", err)
	}

	_, err = tx.Exec(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("ItemsRepo - SaveItem - tx.Exec: %w", err)
	}

	tx.Commit(ctx)

	i.Name = item.Name
	i.Description = item.Description

	return &i, nil

}
