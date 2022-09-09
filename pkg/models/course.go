package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Course struct {
	// Primary key ID
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt int64
	UpdatedAt int64

	// Course Details
	Title       string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	VideoURL    string         `gorm:"not null"`
	Topics      pq.StringArray `gorm:"type:text[]"`
	Duration    float64        `gorm:"type:decimal"`
	Category    string         `gorm:"not null"`
	Points      uint           `gorm:"not null"`

	// Super Admin Approval
	IsApproved bool `gorm:"default:false"`

	// Foreign key
	UserID *uuid.UUID
}

// Join Table for viewed courses (pending/completed) by user (employee/admin)
type ViewedCourses struct {
	UserID            uuid.UUID `gorm:"primaryKey"`
	CourseID          uuid.UUID `gorm:"primaryKey"`
	LastViewed        int64
	CompletedDuration float64 `gorm:"type:decimal"`
	IsCompleted       bool
}

// Hooks; get triggered on every creation or updation on course table row
func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now().Unix()
	return nil
}

func (c *Course) BeforeSave(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().Unix()
	return nil
}

func (c *Course) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now().Unix()
	return nil
}
