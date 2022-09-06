package models

import (
	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title       string
	Description string
	VideoURL    string
	Topics      string
	Duration    string
	Category    string
}
