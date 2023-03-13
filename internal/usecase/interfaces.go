// busyness logic interfaces
package usecase

import (
	"context"

	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
)

type (
	// Item
	Item interface {
		Get(context.Context) ([]models.Item, error)
		Save(context.Context, *models.Item) (*models.Item, error)
		Delete(context.Context, int, int) (bool, error)
		Update(context.Context, *models.Item) (*models.Item, error)
		SetCache(ctx context.Context, key string, value []byte) error
		GetCache(ctx context.Context, key string) ([]byte, error)
		InvalidateCache(ctx context.Context, key string) error
	}

	// ItemRepo
	ItemsRepo interface {
		GetItems(context.Context) ([]models.Item, error)
		SaveItem(context.Context, *models.Item) (*models.Item, error)
		DeleteItem(context.Context, int, int) (bool, error)
		UpdateItem(context.Context, *models.Item) (*models.Item, error)
	}

	// Cache
	Cacher interface {
		Set(ctx context.Context, key string, value []byte) error
		Get(ctx context.Context, key string) ([]byte, error)
		Invalidate(ctx context.Context, key string) error
	}
)
