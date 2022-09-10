package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Employee   Role = "employee"
	Admin      Role = "admin"
	Superadmin Role = "superadmin"
)

type User struct {
	// Primary key ID
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt int64
	UpdatedAt int64

	// User Details
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Role         Role   `gorm:"type:role;default:employee"`
	RewardPoints uint

	// Courses designed by Admin, hence will only non-empty if user is Admin
	Courses []Course

	// Courses viewed by Employee (pending/completed)
	ViewedCourses []Course `gorm:"many2many:user_courses;"`
}

// Hooks; get triggered on every creation or updation on user table row
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now().Unix()
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().Unix()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().Unix()
	return nil
}
