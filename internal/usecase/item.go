package usecase

import (
	"context"
	"fmt"

	"github.com/iamthe1whoknocks/hezzl_test_task/internal/models"
	"github.com/iamthe1whoknocks/hezzl_test_task/internal/usecase/repo"
)

type ItemUseCase struct {
	repo repo.ItemsRepo
}

func New(r repo.ItemsRepo) *ItemUseCase {
	return &ItemUseCase{
		repo: r,
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
	return item, nil
}

func (iu *ItemUseCase) Delete(ctx context.Context, id, campaignID int) (bool, error) {
	isDeleted, err := iu.repo.DeleteItem(ctx, id, campaignID)
	if err != nil {
		return false, fmt.Errorf("ItemsUseCase - Delete - iu.repo.DeleteItem : %w", err)
	}
	return isDeleted, nil
}

func (iu *ItemUseCase) Update(ctx context.Context, item *models.Item) (*models.Item, error) {
	item, err := iu.repo.UpdateItem(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("ItemsUseCase - Save - iu.repo.SaveItem : %w", err)
	}
	return item, nil
}
