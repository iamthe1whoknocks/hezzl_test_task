// busyness usecase logic
package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase/repo"
)

type ItemUseCase struct {
	repo   repo.ItemsRepo
	cache  Cacher
	broker Broker
}

func New(r repo.ItemsRepo, c Cacher, b Broker) *ItemUseCase {
	return &ItemUseCase{
		repo:   r,
		cache:  c,
		broker: b,
	}
}

func (iu *ItemUseCase) Get(ctx context.Context) ([]models.Item, error) {
	items, err := iu.repo.GetItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Get - iu.repo.GetItems : %w", err)
	}
	return items, nil
}

func (iu *ItemUseCase) Save(ctx context.Context, item *models.Item) (*models.Item, error) {
	item, err := iu.repo.SaveItem(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - iu.repo.SaveItem : %w", err)
	}

	data, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - Marshal : %w", err)
	}

	err = iu.broker.Publish(ctx, iu.broker.GetSubject(), data)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - Publish : %w", err)
	}

	return item, nil
}

func (iu *ItemUseCase) Delete(ctx context.Context, id, campaignID int) (bool, error) {
	isDeleted, err := iu.repo.DeleteItem(ctx, id, campaignID)
	if err != nil {
		return false, fmt.Errorf("ItemsUseCase - Delete - iu.repo.DeleteItem : %w", err)
	}

	data, err := json.Marshal(models.Item{
		ID:        id,
		CampainID: campaignID,
		Removed:   isDeleted,
		CreatedAt: time.Now(), //done for simplicity
	})

	if err != nil {
		return false, fmt.Errorf("ItemsUseCase - Delete - Marshal : %w", err)
	}

	err = iu.broker.Publish(ctx, iu.broker.GetSubject(), data)
	if err != nil {
		return false, fmt.Errorf("ItemsUseCase - Delete - Publish : %w", err)
	}
	return isDeleted, nil
}

func (iu *ItemUseCase) Update(ctx context.Context, item *models.Item) (*models.Item, error) {
	item, err := iu.repo.UpdateItem(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - iu.repo.SaveItem : %w", err)
	}

	data, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - Marshal : %w", err)
	}

	err = iu.broker.Publish(ctx, iu.broker.GetSubject(), data)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - Publish : %w", err)
	}

	return item, nil
}

func (iu *ItemUseCase) SetCache(ctx context.Context, key string, value []byte) error {
	err := iu.cache.Set(ctx, key, value)
	if err != nil {
		return fmt.Errorf("ItemsUseCase - SetCache - iu.cache.Set: %w", err)
	}
	return nil
}

func (iu *ItemUseCase) GetCache(ctx context.Context, key string) ([]byte, error) {
	b, err := iu.cache.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - GetCache - iu.cache.Get: %w", err)
	}
	return b, nil
}

func (iu *ItemUseCase) InvalidateCache(ctx context.Context, key string) error {
	err := iu.cache.Invalidate(ctx, key)
	if err != nil {
		return fmt.Errorf("ItemsUseCase - InvalidateCache - iu.cache.Invalidate: %w", err)
	}
	return nil
}
