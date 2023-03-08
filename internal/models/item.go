package models

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `json:"id" valid:",required"`
	CampainID   uuid.UUID `json:"campain_id" valid:",required"`
	Name        string    `json:"name" valid:",required"`
	Description string    `json:"description" valid:",required"`
	Priority    int       `json:"priority" valid:",required"`
	Removed     bool      `json:"is_removed" valid:",required"`
	CreatedAt   time.Time `json:"created_at" valid:",required"`
}
