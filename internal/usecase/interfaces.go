package usecase

import (
	"context"

	"github.com/hezzl_task5/internal/models"
)

type (
	// Item
	Item interface {
		Get(context.Context) ([]models.Item, error)
	}

	// ItemRepo
	TranslationRepo interface {
		GetItems(context.Context) ([]models.Item, error)
	}
)
