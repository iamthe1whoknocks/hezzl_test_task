package usecase

import (
	"context"
	"fmt"

	"github.com/hezzl_task5/internal/models"
	"github.com/hezzl_task5/internal/usecase/repo"
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
		return nil, fmt.Errorf("ItemsUseCase - Get - iu.repo.GetItems")
	}
	return items, nil
}

func (iu *ItemUseCase) Save(ctx context.Context, name string) (models.Item, error) {
	item, err := iu.repo.SaveItem(ctx, name)
	if err != nil {
		return models.Item{}, fmt.Errorf("ItemsUseCase - Save - iu.repo.SaveItem")
	}
	return item, nil
}
