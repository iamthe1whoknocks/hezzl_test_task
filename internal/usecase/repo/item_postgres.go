package repo

import (
	"context"
	"fmt"

	"github.com/hezzl_task5/internal/models"
	"github.com/hezzl_task5/pkg/postgres"
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

func (r *ItemsRepo) SaveItem(ctx context.Context, name string) (models.Item, error) {
	sql, _, err := r.Builder.Insert("items").Columns("")
}
